package server

import (
	"github.com/martinkaburu/warehouse-manager/pkg/utils"
	"github.com/martinkaburu/warehouse-manager/warehousemanager-server/internal/handlers"
	"gorm.io/gorm"
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) InitializeRoutes(db *gorm.DB) {
	r.HandleFunc("/health", utils.HealthCheckHandler()).
		Methods(http.MethodGet).
		Name("healthcheck")

	r.HandleFunc("/cargo", handlers.CargoManifestHandler(db)).
		Methods(http.MethodGet).
		Name("cargo")

	r.Use(handlers.Pagination)
}
