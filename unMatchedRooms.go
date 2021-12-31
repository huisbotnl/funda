package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type UnMatchedRoom struct {
	gorm.Model
	Price         string `json:"price" gorm:"type:varchar(10);index"`
	PricePeriod   string `json:"price_period" gorm:"type:varchar(10);index"`
	PriceCurrency string `json:"price_currency" gorm:"type:varchar(5);index"`
	BedroomId     int    `json:"bedroom_id" gorm:"type:integer;index"`
	Bedroom       string `json:"bedroom" gorm:"type:varchar(50);index"`
	TypeId        int    `json:"type_id" gorm:"type:integer;index"`
	Type          string `json:"type" gorm:"type:varchar(50);index"`
	FurnitureId   int    `json:"furniture_id" gorm:"type:integer;index"`
	Furniture     string `json:"furniture" gorm:"type:varchar(50);index"`
	CityId        int    `json:"city_id" gorm:"type:integer;index"`
	City          string `json:"city" gorm:"type:varchar(50);index"`
	DistrictsId   int    `json:"districts_id" gorm:"type:integer;index"`
	Districts     string `json:"districts" gorm:"type:varchar(50);index"`
	FacilitiesIds string `json:"facilities_ids" gorm:"type:text;index"`
	Facilities    string `json:"facilities" gorm:"type:text;index"`
	AmenitiesIds  string `json:"amenities_ids" gorm:"type:text;index"`
	Amenities     string `json:"amenities" gorm:"type:text;index"`
	Url           string `json:"url" gorm:"type:varchar(100);index"`
	ScraperName   string `json:"scraper_name" gorm:"type:varchar(20);index"`
	Image         string `json:"image" gorm:"type:varchar(200);index"`
}

func (unMatchedRoom UnMatchedRoom) setFacilities(facilities []string) {
	data, _ := json.Marshal(&facilities)
	unMatchedRoom.Facilities = string(data)
}

func (unMatchedRoom UnMatchedRoom) setFacilitiesIds(facilities []int) {
	data, _ := json.Marshal(&facilities)
	unMatchedRoom.FacilitiesIds = string(data)
}

func (unMatchedRoom UnMatchedRoom) setAmenities(amenities []string) {
	data, _ := json.Marshal(&amenities)
	unMatchedRoom.Amenities = string(data)
}

func (unMatchedRoom UnMatchedRoom) setAmenitiesIds(amenities []int) {
	data, _ := json.Marshal(&amenities)
	unMatchedRoom.Amenities = string(data)
}
