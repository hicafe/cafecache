package handler

import (
	"codes"
)

func deleteBucket(bucketName string) int {
	lock.Lock()
	defer lock.Unlock()
	delete(buckets, bucketName)
	delete(bucketsLock, bucketName)
	return codes.OK
}
