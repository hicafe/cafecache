package handler

import (
	"encoding/json"
	"fmt"
	"sync"
)

var buckets = make(map[string]map[string]string)
var bucketsLock = make(map[string]sync.RWMutex)
var lock = new(sync.RWMutex)

func HandleData(data []byte) {
	var _map = make(map[string]string)
	json.Unmarshal(data, &_map)

	op := _map["o"]
	bucket := _map["b"]
	key := _map["k"]
	value := _map["v"]

	switch op {
	case "CREATE_BUCKET":
		createBucket(bucket)
	case "DEL_BUCKET":
		deleteBucket(bucket)
	case "PUT":
		put(bucket, key, value)
	case "DEL":
		del(bucket, key)
	default:
		fmt.Println(op)
	}
}
