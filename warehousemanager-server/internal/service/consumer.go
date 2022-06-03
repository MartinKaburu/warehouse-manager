package service

import (
	"fmt"
	"github.com/martinkaburu/warehouse-manager/pkg/models"
	"github.com/martinkaburu/warehouse-manager/pkg/utils"
	"github.com/martinkaburu/warehouse-manager/warehousemanager-server/internal/handlers"
	consumer "github.com/martinkaburu/warehouse-manager/warehouseproto"
	"log"
	"strconv"
)

type OrderConsumerServer struct {
	consumer.UnimplementedOrderConsumerServer
}

func (o *OrderConsumerServer) ReceiveOrders(stream consumer.OrderConsumer_ReceiveOrdersServer) error {
	log.Println("Streaming Orders From Client")
	dbConfig := DBConfig{Host: utils.GetEnv("HOST", "localhost"), Port: utils.GetEnv("DB_PORT", "5432"), User: utils.GetEnv("DB_USER", "postgres"), Name: utils.GetEnv("DB_NAME", "postgres")}

	db, err := ConnectToPostgres(fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Name, dbConfig.Password))
	if err != nil {
		log.Fatalf("Unable to connect to Postgres: %v", err)
	}

	orders := make([]models.Order, 0)

	for {
		entry, _ := stream.Recv()
		if entry == nil {
			break
		}

		externalID, _ := strconv.ParseInt(entry.Order.Id, 10, 64)

		order := models.Order{
			ExternalId: externalID,
			Email:      entry.Order.Email,
			Phone:      entry.Order.Phone,
			Weight:     entry.Order.Weight,
		}
		orders = append(orders, order)
	}
	log.Println("Inserting Orders to DB")
	handlers.InsertOrders(db, orders)

	log.Println("Done Inserting Orders to DB")

	log.Println("Calculating Cargo")

	calculateCargo(db)

	return nil
}
