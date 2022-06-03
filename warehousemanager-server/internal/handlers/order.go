package handlers

import (
	"github.com/martinkaburu/warehouse-manager/pkg/models"
	"gorm.io/gorm"
)

func InsertOrders(db *gorm.DB, orders []models.Order) []models.Order {
	db.CreateInBatches(orders, 1000)
	return orders
}
