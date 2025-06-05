package activity

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// Quick & simple CSV serialization assuming flat structs with string/int fields
func toCSVHeader(entry interface{}) ([]string, error) {
	val := reflect.ValueOf(entry)
	if val.Kind() != reflect.Struct {
		return nil, errors.New("CSV format only supported for struct types")
	}

	typ := val.Type()
	var header []string
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		name := f.Tag.Get("json")
		if name == "" || name == "-" {
			name = f.Name
		}
		header = append(header, name)
	}
	return header, nil
}

func toCSVRecord(entry interface{}) ([]string, error) {
	val := reflect.ValueOf(entry)
	if val.Kind() != reflect.Struct {
		return nil, errors.New("CSV format only supported for struct types")
	}

	var record []string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			record = append(record, field.String())
		case reflect.Int:
			record = append(record, strconv.Itoa(int(field.Int())))
		default:
			record = append(record, fmt.Sprintf("%v", field.Interface()))
		}
	}
	return record, nil
}
