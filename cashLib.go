package cacshLib

var myMap = make(map[string]interface{})

func Set(key string, value interface{}) {
	myMap[key] = value
}

func Get(key string) interface{} {
	return myMap[key]
}

func Delete(key string) {
	delete(myMap, key)
}
