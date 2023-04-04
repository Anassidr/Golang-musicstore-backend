package models

import (
	"github.com/anassidr/go-musicstore/pkg/config"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
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

func (i *Instrument) CreateInstrument() *Instrument {
	db.NewRecord(i)
	db.Create(&i)
	return i
}

func GetAllInstruments() []Instrument {
	var Instruments []Instrument
	db.Find(&Instruments)
	return Instruments
}

func GetInstrumentById(Id int64) (*Instrument, *gorm.DB) {
	var getInstrument Instrument
	db := db.Where("ID=?", Id).Find(&getInstrument)
	return &getInstrument, db
}

func DeleteInstrument(Id int64) Instrument {
	var instrument Instrument
	db.Where("ID=?", Id).Delete(instrument)
	return instrument
}

type User struct {
	gorm.Model
	Name               string `json:"Name" valid:"required"`
	Address            string `json:"address" valid:"required"`
	Username           string `json:"username" valid:"required"`
	Email              string `json:"email" valid:"required,email"`
	Email_verification string `json:"email_verification" valid:"required"`
	PasswordHash       string `json:"-"` // not exposed in JSON responses
	Role               string `json:"role" valid:"required"`
	Gender             string `json:"qty_stock" valid:"required,numeric,min=0"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
