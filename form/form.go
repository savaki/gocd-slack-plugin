package form

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
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

	formStruct := reflect.ValueOf(v).Elem()

	for i := 0; i < r.NumField(); i++ {
		typeField := r.Field(i)
		if key := typeField.Tag.Get("form"); key != "" {
			structField := formStruct.Field(i)
			value := Stringify(typeField, structField)
			if value != "" {
				params.Add(key, value)
			}
		}
	}

	return params
}

func Stringify(typeField reflect.StructField, structField reflect.Value) string {
	switch structField.Type().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v := int(structField.Int())
		return strconv.Itoa(v)
	case reflect.String:
		return structField.String()
	case reflect.Bool:
		return strconv.FormatBool(structField.Bool())
	case reflect.Float32:
		return strconv.FormatFloat(structField.Float(), 'f', precision(typeField), 32)
	case reflect.Float64:
		return strconv.FormatFloat(structField.Float(), 'f', precision(typeField), 64)
	case reflect.Struct, reflect.Array, reflect.Slice:
		v, err := json.Marshal(structField.Interface())
		if err != nil {
			return "ERROR - unable to marshal field"
		}
		return string(v)
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
