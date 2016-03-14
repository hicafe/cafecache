package handler

import (
	"codes"
	"sync"
)

func createBucket(bucketName string) int {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := buckets[bucketName]; ok {
		return codes.OK
	}

	_map := make(map[string]string)
	buckets[bucketName] = _map
	lock := new(sync.RWMutex)
	bucketsLock[bucketName] = *lock
	return codes.OK
}
