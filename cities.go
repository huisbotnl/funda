package main

import "github.com/jinzhu/gorm"

var citiesMap = map[string]map[string]interface{}{
	"Alkmaar": {
		"key": "alkmaar",
	},
	"Almere-Stad": {
		"key": "almere_stad",
	},
	"Alphen aan den Rijn": {
		"key": "alphen_aan_den_rijn",
	},
	"Amersfoort": {
		"key": "amersfoort",
	},
	"Amstelveen": {
		"key": "amstelveen",
	},
	"Amsterdam": {
		"key": "amsterdam",
	},
	"Bleiswijk": {
		"key": "bleiswijk",
	},
	"Capelle aan den IJssel": {
		"key": "capelle_aan_den_i_jssel",
	},
	"Delft": {
		"key": "delft",
	},
	"Diemen": {
		"key": "diemen",
	},
	"Driebergen-rijsenburg": {
		"key": "driebergen_rijsenburg",
	},
	"Egmond aan den Hoef": {
		"key": "egmond_aan_den_hoef",
	},
	"Enschede": {
		"key": "enschede",
	},
	"Groningen": {
		"key": "groningen",
	},
	"Haarlem": {
		"key": "haarlem",
	},
	"Haren": {
		"key": "haren",
	},
	"Hazerswoude-Rijndijk": {
		"key": "hazerswoude_rijndijk",
	},
	"Hengelo": {
		"key": "hengelo",
	},
	"Hillegom": {
		"key": "hillegom",
	},
	"Hilversum": {
		"key": "hilversum",
	},
	"Hoofddorp": {
		"key": "hoofddorp",
	},
	"Leiden": {
		"key": "leiden",
	},
	"Maarssen": {
		"key": "maarssen",
	},
	"Nieuw-Vennep": {
		"key": "nieuw_vennep",
	},
	"Nieuwveen": {
		"key": "nieuwveen",
	},
	"Ouderkerk aan de Amstel": {
		"key": "ouderkerk_aan_de_amstel",
	},
	"Putten": {
		"key": "putten",
	},
	"Rijswijk": {
		"key": "rijswijk",
	},
	"Rotterdam": {
		"key": "rotterdam",
	},
	"Scheveningen": {
		"key": "scheveningen",
	},
	"Schiedam": {
		"key": "schiedam",
	},
	"The Hague": {
		"key": "the_hague",
	},
	"Nieuwegein": {
		"key": "nieuwegein",
	},
	"Breda": {
		"key": "breda",
	},
	"Utrecht": {
		"key": "utrecht",
	},
	"Veenendaal": {
		"key": "veenendaal",
	},
	"Voorburg": {
		"key": "voorburg",
	},
	"Voorschoten": {
		"key": "voorschoten",
	},
	"Warmond": {
		"key": "warmond",
	},
	"Zaandam": {
		"key": "zaandam",
	},
	"Zeist": {
		"key": "zeist",
	},

	"Zoetermeer": {
		"key": "zoetermeer",
	},
	"Almere": {
		"key": "almere",
	},
	"Maastricht": {
		"key": "maastricht",
	},
	"Vlaardingen": {
		"key": "vlaardingen",
	},
	"koop": {
		"key": "koop",
	},
}

func setupCitiesMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range citiesMap {
		DB.Raw("SELECT * FROM `cities`  WHERE (`key` = ?) ORDER BY `cities`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		citiesMap[name] = data
	}
}

type City struct {
	gorm.Model
	Key            string `json:"key" gorm:"type:varchar(50);index;unique"`
	Name           string `json:"name" gorm:"type:varchar(50);index;unique"`
	Neighbourhoods []Neighbourhood
}
