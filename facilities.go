package main

var facilitiesMap = map[string]map[string]interface{}{
	"Pets": {
		"key": "pets_allowed",
	},
	"Parking": {
		"key": "parking",
	},
	"Balcony": {
		"key": "balcony",
	},
	"Private bathroom": {
		"key": "private_bathroom",
	},
}

func setupFacilitiesMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range facilitiesMap {
		DB.Raw("SELECT * FROM `facilities`  WHERE (`key` = ?) ORDER BY `facilities`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		facilitiesMap[name] = data
	}
}
