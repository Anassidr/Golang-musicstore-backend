package models

import (
	"github.com/anassidr/go-musicstore/pkg/config"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

type Instrument struct {
	gorm.Model
	Name        string  `json:"name" valid:"required"`
	Brand       string  `json:"brand" valid:"required"`
	Type        string  `json:"type" valid:"required"`
	Price       float64 `json:"price" valid:"required,numeric,min=0"`
	Color       string
	Description string `json:"description" valid:"required"`
	Qty_stock   int    `json:"qty_stock" valid:"required,numeric,min=0"`
}

// initialize the database connection and create any necessary database
// tables for the Instrument model before the application starts running

func (i *Instrument) Validate() error {
	_, err := govalidator.ValidateStruct(i)
	return err
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Instrument{}) //Automigrate expects a reference to a struct type
}
