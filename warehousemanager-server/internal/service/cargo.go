package service

import (
	"github.com/martinkaburu/warehouse-manager/pkg/models"
	"gorm.io/gorm"
	"log"
)

type CargoList struct {
	Orders  []models.Order
	Weight  float64
	Country string
}

const (
	MaxWeight float64 = 500.00

	Mozambique string = "MOZAMBIQUE"
	Uganda     string = "UGANDA"
	Cameroon   string = "CAMEROON"
	Morocco    string = "MOROCCO"
	Ethiopia   string = "ETHIOPIA"
)

func (cl *CargoList) AddOrder(db *gorm.DB, order models.Order) {
	if full := cl.isFull(order); full == true {
		cl.commitOrder(db)
		*cl = CargoList{Orders: []models.Order{}, Weight: 0.00, Country: cl.Country}
	}

	cl.Orders = append(cl.Orders, order)
}

func (cl *CargoList) isFull(next models.Order) bool {
	newCargoWeight := cl.Weight + next.Weight
	if newCargoWeight >= MaxWeight {
		return true
	}
	cl.Weight = newCargoWeight
	return false
}

func (cl *CargoList) commitOrder(db *gorm.DB) {
	db.Create(&models.Cargo{
		Orders:  cl.Orders,
		Weight:  cl.Weight,
		Country: cl.Country,
	})
}

func calculateCargo(db *gorm.DB) {
	go func() {
		log.Printf("Calculating %v Cargo!", Cameroon)
		var caOrders []models.Order
		db.Raw("select *, REGEXP_MATCHES(phone, '(237) ?[2368]\\d{7,8}$') from orders order by weight asc;").Scan(&caOrders)
		var cl CargoList = struct {
			Orders  []models.Order
			Weight  float64
			Country string
		}{Orders: []models.Order{}, Weight: 0.00, Country: Cameroon}

		for _, order := range caOrders {
			cl.AddOrder(db, order)
		}
		log.Printf("Completed %v Cargo!", Cameroon)
	}()
	go func() {
		log.Printf("Calculating %v Cargo!", Uganda)
		var ugOrders []models.Order
		db.Raw("select *, REGEXP_MATCHES(phone, '(256) ?\\d{9}$') from orders order by weight asc;").Scan(&ugOrders)
		var cl CargoList = struct {
			Orders  []models.Order
			Weight  float64
			Country string
		}{Orders: []models.Order{}, Weight: 0.00, Country: Uganda}

		for _, order := range ugOrders {
			cl.AddOrder(db, order)
		}
		log.Printf("Completed %v Cargo!", Uganda)
	}()
	go func() {
		log.Printf("Calculating %v Cargo!", Ethiopia)
		var etOrders []models.Order
		db.Raw("select *, REGEXP_MATCHES(phone, '(251) ?[1-59]\\d{8}$') from orders order by weight asc;").Scan(&etOrders)
		var cl CargoList = struct {
			Orders  []models.Order
			Weight  float64
			Country string
		}{Orders: []models.Order{}, Weight: 0.00, Country: Ethiopia}

		for _, order := range etOrders {
			cl.AddOrder(db, order)
		}
		log.Printf("Completed %v Cargo!", Ethiopia)
	}()
	go func() {
		log.Printf("Calculating %v Cargo!", Morocco)
		var moOrders []models.Order
		db.Raw("select *, REGEXP_MATCHES(phone, '(212) ?[5-9]\\d{8}$') from orders order by weight asc;").Scan(&moOrders)
		var cl CargoList = struct {
			Orders  []models.Order
			Weight  float64
			Country string
		}{Orders: []models.Order{}, Weight: 0.00, Country: Morocco}

		for _, order := range moOrders {
			cl.AddOrder(db, order)
		}
		log.Printf("Completed %v Cargo!", Morocco)
	}()
	go func() {
		log.Printf("Calculating %v Cargo!", Mozambique)
		var mzOrders []models.Order
		db.Raw("select *, REGEXP_MATCHES(phone, '(258) ?[28]\\d{7,8}$') from orders order by weight asc;").Scan(&mzOrders)
		var cl CargoList = struct {
			Orders  []models.Order
			Weight  float64
			Country string
		}{Orders: []models.Order{}, Weight: 0.00, Country: Mozambique}

		for _, order := range mzOrders {
			cl.AddOrder(db, order)
		}
		log.Printf("Completed %v Cargo!", Mozambique)
	}()
}
