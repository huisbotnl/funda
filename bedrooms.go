package main

var bedroomsMap = map[string]map[string]interface{}{
	"1 room": {
		"key": "1",
	},
	"1 rooms": {
		"key": "1",
	},
	"2 rooms": {
		"key": "2",
	},
	"3 rooms": {
		"key": "3",
	},
	"4 rooms": {
		"key": "4",
	},
	"5 rooms": {
		"key": "+5",
	},
	"6 rooms": {
		"key": "+5",
	},
	"7 rooms": {
		"key": "+5",
	},
	"8 rooms": {
		"key": "+5",
	},
	"9 rooms": {
		"key": "+5",
	},
}

func setupBedroomsMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range bedroomsMap {
		DB.Raw("SELECT * FROM `bedrooms`  WHERE (`key` = ?) ORDER BY `bedrooms`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		bedroomsMap[name] = data
	}
}
