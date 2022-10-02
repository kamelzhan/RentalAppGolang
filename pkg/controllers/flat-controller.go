package controllers

import (
	"encoding/json"
	"example/rental_app/pkg/models"
	"example/rental_app/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewFlat models.Flat

func GetFlat(w http.ResponseWriter, r *http.Request) {
	newFlats := models.GetAllFlats()
	res, _ := json.Marshal(newFlats)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFlatById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flatId := vars["flatid"]
	ID, err := strconv.ParseInt(flatId, 0, 0)
	if err != nil {
		fmt.Println("err")
	}
	flatDetails, _ := models.GetFlatById(ID)
	res, _ := json.Marshal(flatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateFlat(w http.ResponseWriter, r *http.Request) {
	CreateFlat := &models.Flat{}
	utils.ParseBody(r, CreateFlat)
	f := CreateFlat.CreateFlat()
	res, _ := json.Marshal(f)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFlat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	flatId := vars["flatId"]
	ID, err := strconv.ParseInt(flatId, 0, 0)
	if err != nil {
		fmt.Println("err")
	}
	flat := models.DeleteFlat(ID)
	res, _ := json.Marshal(flat)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateFlat(w http.ResponseWriter, r *http.Request) {
	var updateFlat = &models.Flat{}
	utils.ParseBody(r, updateFlat)
	vars := mux.Vars(r)
	flatId := vars["flatId"]
	ID, err := strconv.ParseInt(flatId, 0, 0)
	if err != nil {
		fmt.Println("err")
	}
	flatDetails, db := models.GetFlatById(ID)
	if updateFlat.City != "" {
		flatDetails.City = updateFlat.City
	}
	flatDetails.Price = updateFlat.Price
	if updateFlat.Address != "" {
		flatDetails.Address = updateFlat.Address
	}
	flatDetails.IsActive = updateFlat.IsActive
	db.Save(&flatDetails)
	res, _ := json.Marshal(flatDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
