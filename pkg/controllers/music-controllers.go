package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/anassidr/go-musicstore/pkg/models"
	"github.com/anassidr/go-musicstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewInstrument models.Instrument

func GetInstrument(w http.ResponseWriter, r *http.Request) {
	newInstruments := models.GetAll(models.Instrument{})
	res, _ := json.Marshal(newInstruments)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetInstrumentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	instrumentId := vars["instrumentId"]
	ID, err := strconv.ParseInt(instrumentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	instrumentDetails, _ := models.GetInstrumentById(ID)
	res, _ := json.Marshal(instrumentDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateInstrument(w http.ResponseWriter, r *http.Request) {
	CreateInstrument := &models.Instrument{}
	utils.ParseBody(r, CreateInstrument)
	i := CreateInstrument.CreateInstrument()
	res, _ := json.Marshal(i)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteInstrument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	instrumentId := vars["instrumentId"]
	ID, err := strconv.ParseInt(instrumentId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	instrument := models.DeleteInstrument(ID)
	res, _ := json.Marshal(instrument)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateInstrument(w http.ResponseWriter, r *http.Request) {
	var updateInstrument = &models.Instrument{}
	utils.ParseBody(r, updateInstrument)
	vars := mux.Vars(r)
	instrumentId := vars["instrumentId"]
	ID, err := strconv.ParseInt(instrumentId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	instrumentDetails, db := models.GetInstrumentById(ID)

	// reflection
	elem := reflect.ValueOf(instrumentDetails).Elem()
	update := reflect.ValueOf(updateInstrument).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		if field.CanSet() {
			updateField := update.Field(i)
			if !updateField.IsZero() {
				field.Set(updateField)
			}
		}
	}

	db.Save(&instrumentDetails)
	res, _ := json.Marshal(instrumentDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
