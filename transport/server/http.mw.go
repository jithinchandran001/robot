package server

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"robot/constant"
	"robot/pkg/codecs"
	"robot/pkg/logger"

	"net/http"

	debugPkg "robot/pkg/debug"

	"runtime/debug"
)

func NewHttpMiddleware(r *mux.Router) []mux.MiddlewareFunc {
	mw := make([]mux.MiddlewareFunc, 0)
	mw = append(mw, corsMiddleware)
	mw = append(mw, RecoveryHttpMiddleware)
	return mw
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		var validDomain = request.Header.Get("Origin")
		crs := cors.New(cors.Options{
			AllowedOrigins:         []string{validDomain},
			AllowOriginFunc:        nil,
			AllowOriginRequestFunc: nil,
			AllowedMethods: []string{http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete},
			AllowedHeaders:     []string{"*"},
			ExposedHeaders:     []string{c.CacheHeader},
			AllowCredentials:   true,
			OptionsPassthrough: false,
		})
		crs.ServeHTTP(writer, request, func(writer http.ResponseWriter, request *http.Request) {
			next.ServeHTTP(writer, request)
		})
	})
}

func RecoveryHttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer httpRecover(writer)
		next.ServeHTTP(writer, request)
	})
}

func httpRecover(writer http.ResponseWriter) {
	if a := recover(); a != nil {
		var debugMessage = "[http mw] Unknown error encounters, seems panic, has recovered"
		if v, ok := a.(error); ok {
			debugMessage = v.Error()
		}

		logger.Get().ErrorWithoutSTT("system error", "recovery", true, "error", a, "stack trace", string(debug.Stack()))
		codecs.ResponseInternalServerError(debugPkg.DebugMessage(debugMessage, "system error"), writer)
	}
}
