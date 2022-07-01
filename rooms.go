package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Room struct {
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

func (room *Room) setFacilities(facilities []string, ids []int) {
	data, _ := json.Marshal(&facilities)
	dataIds, _ := json.Marshal(&ids)
	room.Facilities = string(data)
	room.FacilitiesIds = string(dataIds)
}

func (room *Room) setAmenities(amenities []string, ids []int) {
	data, _ := json.Marshal(&amenities)
	room.Amenities = string(data)
	dataIds, _ := json.Marshal(&ids)
	room.AmenitiesIds = string(dataIds)
}

func (room Room) UnMatchedRoom() *UnMatchedRoom {
	return &UnMatchedRoom{
		Price:         room.Price,
		PricePeriod:   room.PricePeriod,
		PriceCurrency: room.PriceCurrency,
		BedroomId:     room.BedroomId,
		Bedroom:       room.Bedroom,
		TypeId:        room.TypeId,
		Type:          room.Type,
		FurnitureId:   room.FurnitureId,
		Furniture:     room.Furniture,
		CityId:        room.CityId,
		City:          room.City,
		DistrictsId:   room.DistrictsId,
		Districts:     room.Districts,
		FacilitiesIds: room.FacilitiesIds,
		Facilities:    room.Facilities,
		AmenitiesIds:  room.AmenitiesIds,
		Amenities:     room.Amenities,
		Url:           room.Url,
		ScraperName:   room.ScraperName,
		Image:         room.Image,
	}
}

func (room *Room) SetCity(value string) {
	var city City
	key := RemoveSpecialCharacterAndToLower(value)
	DB.Where("(`key` = ?)", key).First(&city)
	if city.ID == 0 {
		city.Name = value
		city.Key = key
		DB.Save(&city)
	}
	room.CityId = int(city.ID)
	room.City = city.Name
}

func (room *Room) SetDistrict(value string, CityId int) {
	var neighbourhood Neighbourhood
	key := RemoveSpecialCharacterAndToLower(value)
	DB.Where("(`key` = ?)", key).First(&neighbourhood)
	if neighbourhood.ID == 0 {
		neighbourhood.Name = value
		neighbourhood.Key = key
		neighbourhood.CityId = CityId
		DB.Save(&neighbourhood)
	}
	room.DistrictsId = int(neighbourhood.ID)
	room.Districts = neighbourhood.Name
}

func (room *Room) SetType(value string) {
	if data, ok := typeMap[value]; ok {
		room.TypeId = data["id"].(int)
		room.Type = data["name"].(string)
	}else {
		room.Type = value
	}
}

func (room *Room) SetBedroom(value string) {
	if data, ok := bedroomsMap[value]; ok {
		room.BedroomId = data["id"].(int)
		room.Bedroom = data["name"].(string)
	}else {
		room.Bedroom = value
	}
}

func (room *Room) SetFurniture(value string) {
	if data, ok := furnituresMap[value]; ok {
		room.FurnitureId = data["id"].(int)
		room.Furniture = data["name"].(string)
	}else {
		room.Furniture = value
	}
}

func (room *Room) AddFacility(value string) {
	if data, ok := facilitiesMap[value]; ok {
		var facilities []string
		json.Unmarshal([]byte(room.Facilities), &facilities)
		var facilitiesIds []int
		json.Unmarshal([]byte(room.FacilitiesIds), &facilitiesIds)
		facilities = append(facilities, data["name"].(string))
		facilitiesIds = append(facilitiesIds, data["id"].(int))
		room.setFacilities(facilities, facilitiesIds)
	}
}

func (room Room) Create() {
	oldRoom := Room{}
	DB.Where("url = ?", room.Url).First(&oldRoom)
	if oldRoom.ID == 0 {
		if
		room.TypeId != 0 &&
			room.DistrictsId != 0 &&
			room.CityId != 0 &&
			room.BedroomId != 0 &&
			room.FurnitureId != 0 {
			DB.Create(&room)
		} else {
			oldUnMatchedRoom := UnMatchedRoom{}
			DB.Where("url = ?", room.Url).First(&oldUnMatchedRoom)
			if oldUnMatchedRoom.ID == 0 {
				DB.Create(room.UnMatchedRoom())
			}
		}
	}
}
