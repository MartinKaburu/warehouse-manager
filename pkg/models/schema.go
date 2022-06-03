package models

import (
	"github.com/jinzhu/gorm"
)

type OrderEntry struct {
	Id     string `csv:"id"`
	Email  string `csv:"email"`
	Phone  string `csv:"phone_number"`
	Weight string `csv:"parcel_weight"`
}

type Cargo struct {
	gorm.Model `json:"metadata"`
	Orders     []Order `json:"orders"`
	Weight     float64 `json:"weight"`
	Country    string  `json:"country"`
}

type CargoList struct {
	Cargo      []Cargo `json:"cargos"`
	NextPageID int     `json:"next_page_id"`
}

type Order struct {
	gorm.Model `json:"-"`
	ExternalId int64 `gorm:"primaryKey"`
	Email      string
	Phone      string
	Weight     float64
	CargoID    *int
	Cargo      Cargo `json:"-" gorm:"foreignKey:CargoID;references:ID"`
}
