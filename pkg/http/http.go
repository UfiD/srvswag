package pkgHttp

import (
	"net/http"

	"github.com/go-chi/chi"
)

func CreateAndRunServer(addr string, r chi.Router) error {
	httpServer := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return httpServer.ListenAndServe()
}
