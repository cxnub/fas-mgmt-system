package util

import (
	"reflect"
)

func GetJSONTag(obj interface{}, fieldName string) string {
	t := reflect.TypeOf(obj) // Get the type of the object
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Name == fieldName {
			jsonTag := field.Tag.Get("json")
			if jsonTag != "" {
				return jsonTag
			}
		}
	}
	return fieldName // Fallback to the struct field name if no JSON tag
}
