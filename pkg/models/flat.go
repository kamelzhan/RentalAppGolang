package models

import (
	"example/rental_app/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Flat struct {
	gorm.Model
	Id       int64   `gorm:""json:"name"`
	Price    float64 `json:"price"`
	Size     int     `json:"size"`
	Address  string  `json:"address"`
	City     string  `json:"city"`
	IsActive bool    `json:"isActive"`
	owner    int64   `json:"owner"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Flat{})
}
func (f *Flat) CreateFlat() *Flat {
	db.NewRecord(f)
	db.Create(&f)
	return f
}
func GetAllFlats() []Flat {
	var Flats []Flat
	db.Find(&Flats)
	return Flats
}
func GetFlatById(Id int64) (*Flat, *gorm.DB) {
	var getFlat Flat
	db := db.Where("ID=?", Id).Find(&getFlat)
	return &getFlat, db
}
func DeleteFlat(ID int64) Flat {
	var flat Flat
	db.Where("ID=?", ID).Delete(flat)
	return flat

}
