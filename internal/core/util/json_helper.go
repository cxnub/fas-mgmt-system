package util

import (
	"github.com/go-viper/mapstructure/v2"
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

// StructToMap converts any struct to a map[string]interface{}
func StructToMap(input any) (map[string]any, error) {
	output := map[string]any{}
	if err := mapstructure.Decode(input, &output); err != nil {
		return nil, err
	}

	for key, val := range output {
		valType := reflect.TypeOf(val)
		switch valType.Kind() {
		case reflect.Pointer:
			actualVal := reflect.ValueOf(val).Elem().Interface()
			curValMap, err := StructToMap(actualVal)
			if err != nil {
				return nil, err
			}
			output[key] = curValMap
		case reflect.Struct:
			curValMap, err := StructToMap(val)
			if err != nil {
				return nil, err
			}
			output[key] = curValMap
		default:
			continue
		}
	}
	return output, nil
}
