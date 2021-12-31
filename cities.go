package main

import "github.com/jinzhu/gorm"

var citiesMap = map[string]map[string]interface{}{
	"alkmaar": {
		"key": "alkmaar",
	},
	"almere-stad": {
		"key": "almere_stad",
	},
	"alphen-aan-den-rijn": {
		"key": "alphen_aan_den_rijn",
	},
	"amersfoort": {
		"key": "amersfoort",
	},
	"amstelveen": {
		"key": "amstelveen",
	},
	"amsterdam": {
		"key": "amsterdam",
	},
	"bleiswijk": {
		"key": "bleiswijk",
	},
	"capelle-aan=den--jssel": {
		"key": "capelle_aan_den_i_jssel",
	},
	"delft": {
		"key": "delft",
	},
	"diemen": {
		"key": "diemen",
	},
	"driebergen-rijsenburg": {
		"key": "driebergen_rijsenburg",
	},
	"egmond-aan-den-hoef": {
		"key": "egmond_aan_den_hoef",
	},
	"enschede": {
		"key": "enschede",
	},
	"groningen": {
		"key": "groningen",
	},
	"haarlem": {
		"key": "haarlem",
	},
	"haren": {
		"key": "haren",
	},
	"hazerswoude-Rijndijk": {
		"key": "hazerswoude_rijndijk",
	},
	"hengelo": {
		"key": "hengelo",
	},
	"hillegom": {
		"key": "hillegom",
	},
	"hilversum": {
		"key": "hilversum",
	},
	"hoofddorp": {
		"key": "hoofddorp",
	},
	"leiden": {
		"key": "leiden",
	},
	"maarssen": {
		"key": "maarssen",
	},
	"nieuw-vennep": {
		"key": "nieuw_vennep",
	},
	"nieuwveen": {
		"key": "nieuwveen",
	},
	"ouderkerk aan de Amstel": {
		"key": "ouderkerk_aan_de_amstel",
	},
	"putten": {
		"key": "putten",
	},
	"rijswijk": {
		"key": "rijswijk",
	},
	"rotterdam": {
		"key": "rotterdam",
	},
	"scheveningen": {
		"key": "scheveningen",
	},
	"schiedam": {
		"key": "schiedam",
	},
	"the-hague": {
		"key": "the_hague",
	},
	"nieuwegein": {
		"key": "nieuwegein",
	},
	"breda": {
		"key": "breda",
	},
	"utrecht": {
		"key": "utrecht",
	},
	"veenendaal": {
		"key": "veenendaal",
	},
	"voorburg": {
		"key": "voorburg",
	},
	"voorschoten": {
		"key": "voorschoten",
	},
	"warmond": {
		"key": "warmond",
	},
	"zaandam": {
		"key": "zaandam",
	},
	"zeist": {
		"key": "zeist",
	},
	"zoetermeer": {
		"key": "zoetermeer",
	},
	"almere": {
		"key": "almere",
	},
	"maastricht": {
		"key": "maastricht",
	},
	"vlaardingen": {
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
