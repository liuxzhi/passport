package util

import "reflect"

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var hashMap = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		hashMap[t.Field(i).Name] = v.Field(i).Interface()
	}

	return hashMap
}
