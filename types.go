package main

var typeMap = map[string]map[string]interface{}{
	"bouwgrond": {
		"key": "land",
	},
	"huis": {
		"key": "house",
	},
	"appartement": {
		"key": "apartment",
	},
}

func setupTypesMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range typeMap {
		DB.Raw("SELECT * FROM `property_types`  WHERE (`key` = ?) ORDER BY `property_types`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		typeMap[name] = data
	}
}
