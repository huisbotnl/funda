package main

var furnituresMap = map[string]map[string]interface{}{
	"No": {
		"key": "unfurnished",
	},
}

func setupFurnituresMap() {
	typeData := struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	}{}
	for name, data := range furnituresMap {
		DB.Raw("SELECT * FROM `furnitures`  WHERE (`key` = ?) ORDER BY `furnitures`.`id` ASC LIMIT 1", data["key"]).Scan(&typeData)
		data["name"] = typeData.Name
		data["id"] = typeData.ID
		furnituresMap[name] = data
	}
}
