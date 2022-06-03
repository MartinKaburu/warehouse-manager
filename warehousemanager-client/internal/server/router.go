package server

import (
	"github.com/martinkaburu/warehouse-manager/pkg/utils"
	"github.com/martinkaburu/warehouse-manager/warehousemanager-client/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) InitializeRoutes(port *int) {
	r.HandleFunc("/health", utils.HealthCheckHandler()).
		Methods(http.MethodGet).
		Name("healthcheck")

	r.HandleFunc("/upload", handlers.UploadHandler(port)).
		Methods(http.MethodPost).
		Name("upload")
}
