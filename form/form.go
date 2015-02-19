package form

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

// the Bool type allows for a three state bool; true, false, and not set
var (
	t     = true
	f     = false
	True  = &t
	False = &f
)

func AsValues(v interface{}) url.Values {
	params := url.Values{}
	r := reflect.TypeOf(v)
	formStruct := reflect.ValueOf(v)

	if r.Kind() == reflect.Ptr {
		return AsValues(formStruct.Elem().Interface())
	}

	for i := 0; i < r.NumField(); i++ {
		typeField := r.Field(i)
		if formTag := typeField.Tag.Get("form"); formTag != "" {
			// extract key and omitEmpty from formTag
			parts := strings.Split(formTag, ",")
			key := parts[0]
			omitEmpty := (len(parts) > 1 && parts[1] == "omitempty")

			// structField contains the value for this field
			structField := formStruct.Field(i)

			// use reflection to stringify the value
			value := Stringify(typeField, structField, omitEmpty)
			if value != "" {
				params.Add(key, value)
			}
		}
	}

	return params
}

func Stringify(typeField reflect.StructField, structField reflect.Value, omitEmpty bool) string {
	switch structField.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := int(structField.Int())
		if omitEmpty && v == 0 {
			return ""
		} else {
			return strconv.Itoa(v)
		}
	case reflect.String:
		return structField.String()
	case reflect.Bool:
		return strconv.FormatBool(structField.Bool())
	case reflect.Float32, reflect.Float64:
		v := structField.Float()
		if omitEmpty && v == 0.0 {
			return ""
		} else {
			return strconv.FormatFloat(structField.Float(), 'f', precision(typeField), 32)
		}
	case reflect.Struct, reflect.Array, reflect.Slice:
		v := structField.Interface()
		if omitEmpty && v == nil {
			return ""
		}
		data, err := json.Marshal(v)
		if err != nil {
			return "ERROR - unable to marshal field"
		}
		return string(data)
	case reflect.Ptr:
		switch v := structField.Interface().(type) {
		case *bool:
			switch v {
			case True:
				return "true"
			case False:
				return "false"
			default:
				return ""
			}
		default:
			return ""
		}
	default:
		return ""
	}
}

func precision(typeField reflect.StructField) int {
	prec := 2
	if p := typeField.Tag.Get("precision"); p != "" {
		if v, err := strconv.Atoi(p); err == nil {
			prec = v
		}
	}

	return prec
}
