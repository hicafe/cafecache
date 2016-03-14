package handler

import (
	"codes"
)

func del(bucket string, key string) int {
	lock.RLock()
	l := bucketsLock[bucket]
	defer lock.Unlock()

	l.Lock()
	if _, ok := buckets[bucket]; ok {
		delete(buckets, key)
		return codes.OK
	} else {
		return codes.BUCKET_NOT_EXIST
	}
	defer l.Unlock()
	return codes.OK
}
