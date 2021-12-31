package main

var amenitiesMap = map[string]map[string]interface{}{
	"Dishwasher": {
		"key": "dishwasher",
	},
	"Washing machine": {
		"key": "washing_machine",
	},
	"Dryer": {
		"key": "dryer",
	},
	"Air conditioning": {
		"key": "air_conditioning",
	},
	"Electrical heating": {
		"key": "heating",
	},
}

func setupAmenitiesMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range amenitiesMap {
		DB.Raw("SELECT * FROM `amenities`  WHERE (`key` = ?) ORDER BY `amenities`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		amenitiesMap[name] = data
	}
}
