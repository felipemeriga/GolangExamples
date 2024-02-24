package example

import "net/http"

func HealthCheckHandler(w http.ResponseWriter, req *http.Request) {
}

func Middleware() func(handler http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(
			func(writer http.ResponseWriter, request *http.Request) {

			},
		)
	}
}

func CreateHttpServer() http.Server {
	middleware := Middleware()

	middleware(http.HandlerFunc(HealthCheckHandler))

	return http.Server{
		Addr:              "",
		Handler:           middleware(http.HandlerFunc(HealthCheckHandler)),
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

}



