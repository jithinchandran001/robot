package bin

import (
	"context"
	"fmt"
	"github.com/oklog/run"
	"net/http"
	"os"
	"os/signal"
	"robot/config"
	"robot/pkg/logger"
	repopg "robot/repository/pg"
	"robot/service"
	"robot/transport/server"
	"time"

	"github.com/go-pg/pg/v10"
)

func Run(mode string, envFile string) (http.Handler, *pg.DB) {
	defer func() {
		if a := recover(); a != nil {
			var debugMessage = "[top-level] Unknown error encounters, seems panic, has recovered"
			if v, ok := a.(error); ok {
				debugMessage = v.Error()
			}
			logger.Get().Error("system error", "recovery", true, "error", debugMessage)
		}
	}()

	err := config.InitConfig(envFile)

	if err != nil {
		logger.Get().Error("loading configs failed", "error", err)
	}
	//connect to db
	db := pg.Connect(&pg.Options{
		Addr:        fmt.Sprintf("%s:%s", config.Get().PGHost, config.Get().PGPort),
		Database:    config.Get().PGDb,
		User:        config.Get().PGUser,
		Password:    config.Get().PGPassword,
		DialTimeout: time.Duration(config.Get().PgConnTimeout) * time.Second,
	})
	//defer db.Close()

	err = db.Ping(context.Background())
	if err != nil {
		logger.Get().ErrorWithoutSTT("DB connection failed. Application exiting", "error", err)
		return nil, nil
	}

	r := repopg.New(db)

	svc := service.New(r)

	// Add handlers
	httpHandler := server.NewHttpHandler(svc)


	// Create new Server instance
	srv := server.NewHttpServer(config.Get().HttpListenAddr, httpHandler)

	// Create new group to run actors in async
	var g = run.Group{}

	// Init the http server
	g = initHttpServer(g, srv)

	// Run the group if the mod is not in Test mode

	g = initCancelInterrupt(g, srv)
	runtimeRes := g.Run()
	if runtimeRes != nil {
		logger.Get().Error("runtime failed", "error", runtimeRes)
	}

	return nil, nil
}

func initHttpServer(g run.Group, srv http.Server) run.Group {
	g.Add(func() error {
		logger.Get().Info("transport http server starting", srv.Addr)
		return srv.ListenAndServe()
	}, func(err error) {
		// Got an error shutdown the server
		logger.Get().Info("transport server shutting down", "error", err)
		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	})

	return g
}

func initCancelInterrupt(g run.Group, srv http.Server) run.Group {
	g.Add(func() error {
		// Create sig channel
		c := make(chan os.Signal, 1)

		// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
		// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
		signal.Notify(c, os.Interrupt)

		// Block until we receive our signal.
		return fmt.Errorf("received signal %s", <-c)
	}, func(err error) {
		logger.Get().Info("transport server shutting down", "error", err)

		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.
		srv.Shutdown(ctx)

		logger.Get().Info("Shutting down")
		os.Exit(0)
	})

	return g
}
