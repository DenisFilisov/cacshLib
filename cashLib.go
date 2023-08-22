package cacshLib

import (
	"errors"
	"time"
)

var myMap = make(map[string]interface{})

func Set(key string, ttl time.Duration, value interface{}) {
	myMap[key] = value
	go worker(ttl, key)
}

func Get(key string) interface{} {
	value := myMap[key]
	if value == nil {
		return errors.New("Key didn't find")
	}
	return value
}

func Delete(key string) {
	delete(myMap, key)
}

func worker(ttl time.Duration, key string) {
	time.Sleep(ttl)
	if Get(key) != nil {
		Delete(key)
	}
}
