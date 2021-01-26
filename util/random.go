package util

import (
	"strconv"
	"time"
)

// GenerateID takes a key and generates an ID
// thats almost guaranteed to be unique.
func GenerateID(key string) string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	return timestamp + "_" + key
}
