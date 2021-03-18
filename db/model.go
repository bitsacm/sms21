package db

import (
	"reflect"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Serialize will return an interface whose "true" type would be
// a model instance struct.
func Serialize(dataType reflect.Type, n neo4j.Node) interface{} {
	result := reflect.New(dataType)
	data := n.Props

	for i := 0; i < result.Elem().NumField(); i++ {
		fieldInfo := dataType.Field(i)

		field := result.Elem().FieldByName(fieldInfo.Name)
		neoFieldName, ok := fieldInfo.Tag.Lookup("neoKey")
		if !ok {
			neoFieldName = fieldInfo.Name
		}

		kind := field.Type().Kind()
		switch kind {

		case reflect.Int64:
			field.SetInt(data[neoFieldName].(int64))
		case reflect.Int32:
			field.SetInt(int64(data[neoFieldName].(int32)))

		case reflect.Float64:
			field.SetFloat(data[neoFieldName].(float64))
		case reflect.Float32:
			field.SetFloat(float64(data[neoFieldName].(float32)))

		case reflect.Bool:
			field.SetBool(data[neoFieldName].(bool))
		case reflect.String:
			field.SetString(data[neoFieldName].(string))
		}
	}

	return result.Elem().Interface()
}

func SerializeEdge(dataType reflect.Type, e neo4j.Relationship) interface{} {
	result := reflect.New(dataType)

	data := e.Props

	for i := 0; i < result.Elem().NumField(); i++ {
		fieldInfo := dataType.Field(i)

		field := result.Elem().FieldByName(fieldInfo.Name)
		neoFieldName, ok := fieldInfo.Tag.Lookup("neoKey")
		if !ok {
			neoFieldName = fieldInfo.Name
		}

		kind := field.Type().Kind()
		switch kind {

		case reflect.Int64:
			field.SetInt(data[neoFieldName].(int64))
		case reflect.Int32:
			field.SetInt(int64(data[neoFieldName].(int32)))

		case reflect.Float64:
			field.SetFloat(data[neoFieldName].(float64))
		case reflect.Float32:
			field.SetFloat(float64(data[neoFieldName].(float32)))

		case reflect.Bool:
			field.SetBool(data[neoFieldName].(bool))
		case reflect.String:
			field.SetString(data[neoFieldName].(string))
		}
	}

	return result.Elem().Interface()
}
