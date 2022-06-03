package server

import (
	"fmt"
	"github.com/martinkaburu/warehouse-manager/pkg/models"
	"github.com/martinkaburu/warehouse-manager/pkg/utils"
	"github.com/martinkaburu/warehouse-manager/warehousemanager-server/internal/service"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type NewServer struct {
	Router *Router
}

//RunServer ...
func RunServer(httpPort int) (err error) {

	log.Printf("Starting HTTP server at localhost:%d", httpPort)

	dbConfig := service.DBConfig{Host: utils.GetEnv("HOST", "localhost"), Port: utils.GetEnv("DB_PORT", "5432"), User: utils.GetEnv("DB_USER", "postgres"), Name: utils.GetEnv("DB_NAME", "postgres")}

	db, err := service.ConnectToPostgres(fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Password))
	if err != nil {
		log.Printf("Unable to connect to Postgres: %v", err)
	}

	autoMigrateTables(db)

	server := NewServer{NewRouter()}
	server.Router.InitializeRoutes(db)

	if err := http.ListenAndServe(fmt.Sprintf("%v:%d", "localhost", httpPort), *server.Router); err != nil {
		return err
	}

	return nil
}

func autoMigrateTables(db *gorm.DB) {
	db.AutoMigrate(&models.Cargo{}, &models.Order{})
}
