package handler

import (
	"encoding/json"
	"fmt"
	"sync"
)

var buckets = make(map[string]map[string]string)
var bucketsLock = make(map[string]sync.RWMutex)
var lock = new(sync.RWMutex)

func HandleData(data []byte) []byte {
	var _map = make(map[string]string)
	json.Unmarshal(data, &_map)

	op := _map["o"]
	bucket := _map["b"]
	key := _map["k"]
	value := _map["v"]
	var code int
	var result string

	switch op {
	case "CREATE_BUCKET":
		code = createBucket(bucket)
	case "DEL_BUCKET":
		code = deleteBucket(bucket)
	case "PUT":
		code = put(bucket, key, value)
	case "DEL":
		code = del(bucket, key)
	case "GET":
		code, result = get(bucket, key)
	default:
		fmt.Println(op)
	}

	res := &resultStruct{}
	res.code = code
	res.value = result
	j, _ := json.Marshal(res)
	return j

}

type resultStruct struct {
	code  int    `json:"code"`
	value string `json:"value"`
}
