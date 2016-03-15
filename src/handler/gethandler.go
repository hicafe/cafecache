package handler

import (
	"codes"
)

func get(bucket string, key string) (int, string) {
	lock.RLock()
	l := bucketsLock[bucket]
	defer lock.Unlock()
	l.Lock()
	if bucketMap, ok := buckets[bucket]; ok {
		value := bucketMap[key]
		return codes.OK, value
	} else {
		return codes.BUCKET_NOT_EXIST, ""
	}
	defer l.Unlock()
	return codes.OK, ""
}
