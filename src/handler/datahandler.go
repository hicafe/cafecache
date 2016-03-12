package handler

import (
	"encoding/json"
	"fmt"
)

func HandleData(data []byte) {
	var _map = make(map[string]interface{})
	json.Unmarshal(data, &_map)

	op := _map["o"]
	bucket := _map["b"]
	key := _map["k"]
	value := _map["v"]

}
