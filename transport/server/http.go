package server

import (
	"context"
	"net/http"
	"robot/config"
	c "robot/constant"
	"robot/pkg/codecs"
	"robot/pkg/debug"
	"robot/pkg/logger"

	"robot/endpoint"
	"robot/service"

	"time"

	kitHttp "github.com/go-kit/kit/transport/http"
	gorillaMux "github.com/gorilla/mux"
)

// this method must be used to run http server
func NewHttpServer(addr string, handler http.Handler) http.Server {
	httpSer := http.Server{
		Addr:              addr,
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       time.Second * 50,
		ReadHeaderTimeout: 0,
		WriteTimeout:      time.Second * 50,
	}
	return httpSer
}

func NewHttpHandler(srv service.RobotService) http.Handler {
	r := gorillaMux.NewRouter()
	ep := endpoint.NewServiceEndpoints(srv)

	opt := kitHttp.ServerErrorEncoder(HttpErrorEncoder)

	r.Path("*").Methods(http.MethodOptions)

	// Greeting Endpoint
	r.Path(config.Get().HttpBaseRequestUrl+c.RouteGreeting).Methods(http.MethodGet, http.MethodOptions).Handler(kitHttp.NewServer(
		ep.Greeting,
		endpoint.DecodeGreetingIncomingRequest,
		endpoint.EncodeGreetingOutgoingResponse,
		opt,
	))


	r.Path(config.Get().HttpBaseRequestUrl+c.RouteSurvivor).Methods(http.MethodPost, http.MethodOptions).Handler(kitHttp.NewServer(
		ep.AddSurvivor,
		endpoint.DecodeAddSurvivorRequest,
		endpoint.EncodeAddSurvivorResponse,
		opt,
	))

	r.Path(config.Get().HttpBaseRequestUrl+c.RouteSurvivor).Methods(http.MethodPatch, http.MethodOptions).Handler(kitHttp.NewServer(
		ep.UpdateLocation,
		endpoint.DecodeLocationRequest,
		endpoint.EncodeLocationResponse,
		opt,
	))

	r.Path(config.Get().HttpBaseRequestUrl+c.RouteSurvivor+c.RouteInfected).Methods(http.MethodPost, http.MethodOptions).Handler(kitHttp.NewServer(
		ep.UpdateInfected,
		endpoint.DecodeInfectedRequest,
		endpoint.EncodeInfectedResponse,
		opt,
	))

	r.Path(config.Get().HttpBaseRequestUrl+c.RouteRobot).Methods(http.MethodGet, http.MethodOptions).Handler(kitHttp.NewServer(
		ep.GetRobot,
		endpoint.DecodeRobotGetRequest,
		endpoint.EncodeRobotGetResponse,
		opt,
	))

	r.Path(config.Get().HttpBaseRequestUrl+c.RouteSurvivor).Methods(http.MethodGet, http.MethodOptions).Handler(kitHttp.NewServer(
		ep.GetSurvivors,
		endpoint.DecodeGetSurvivorsRequest,
		endpoint.EncodeGetSurvivorsResponse,
		opt,
	))

	r.Path(config.Get().HttpBaseRequestUrl+c.RouteSurvivor+c.RouteIDParam).Methods(http.MethodGet, http.MethodOptions).Handler(kitHttp.NewServer(
		ep.GetSurvivor,
		endpoint.DecodeSurvivorGetRequest,
		endpoint.EncodeSurvivorGetResponse,
		opt,
	))

	r.Path(config.Get().HttpBaseRequestUrl+c.RouteReport).Methods(http.MethodGet, http.MethodOptions).Handler(kitHttp.NewServer(
		ep.GetReport,
		endpoint.DecodeReportRequest,
		endpoint.EncodeReportResponse,
		opt,
	))

	r.NotFoundHandler = NotFoundHandler()

	// Add core to endpoints
	mw := NewHttpMiddleware(r)
	if mw != nil && len(mw) > 0 {
		for i := len(mw) - 1; i >= 0; i-- {
			r.Use(mw[i])
		}
	}

	return r
}

type HttpErrorHandler struct {
}

func HttpErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	if v, ok := err.(*codecs.EndpointError); ok {
		if v.Code != 200 && v.Code != 201 {
			codecs.ResponseStatusJson(v.Message, v.Code, w)
			return
		} else {
			logger.Get().ErrorWithoutSTT("EndpointError cannot have code 200 or 201, it must be an error")
			codecs.ResponseInternalServerError(debug.DebugMessage("EndpointError cannot have code 200 or 201, it must be an error", ""), w)
			return
		}
	} else {
		codecs.ResponseInternalServerError(debug.DebugMessage(err.Error(), ""), w)
		return
	}
}

func NotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 not found"))
	})
}
