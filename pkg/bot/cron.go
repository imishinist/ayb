package bot

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	appEngineCronHeader = "X-Appengine-Cron"
)

func fromCron(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if header := r.Header.Get(appEngineCronHeader); header != "true" {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("header check failed"))
			return
		}
		host := echo.ExtractIPDirect()(r)
		if host != "10.0.0.1" || host != "0.1.0.1" {
			w.WriteHeader(http.StatusNotAcceptable)
			w.Write([]byte("ip check failed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
