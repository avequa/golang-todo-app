package core_http_middleware

import (
	"net/http"
	"github.com/google/uuid"
)

const (
	requestIDHeader = "X-Request-ID"
)

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}
			r.Header.Set(requestIDHeader, requestID)
			wr.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(wr, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
			requestID := r.Header.Get(requestIDHeader)
			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)
			ctx := context.WithValue(r.Context(), "log", l )
			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w httpResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)

			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"panic http handler",
					)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFucn(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			before := time.Now()

			log.Debug(
				"income http request",
				zap.Time("time", before.UTC()),
			)

			next.ServeHTTP(w, r)

			log.Debug(
				"done http request", 
				zap.Int("status_code", rw.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Now().Sub(before)),
		)
		})
	}
}


