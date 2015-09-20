package util

import (
	"fmt"
	"reflect"
	"runtime"
	// "runtime/debug"
)

func Inspect(data interface{}) (typ string, value string) {
	typ = ""
	value = ""
	isPointer := false

	if data == nil {
		typ = "Nil"
		return
	}

	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		isPointer = true
	}

	t := v.Type()

	if v.IsValid() {
		isIndex := false
		isMap := false
		isField := false

		switch v.Kind() {
		case reflect.Int:
			typ = "Int"
			value = fmt.Sprintf("%d(0x%x)", v.Int(), v.Int())
		case reflect.Int8:
			typ = "Int8"
			value = fmt.Sprintf("%d(0x%x)", v.Int(), v.Int())
		case reflect.Int16:
			typ = "Int16"
			value = fmt.Sprintf("%d(0x%x)", v.Int(), v.Int())
		case reflect.Int32:
			typ = "Int32"
			value = fmt.Sprintf("%d(0x%x)", v.Int(), v.Int())
		case reflect.Int64:
			typ = "Int64"
			value = fmt.Sprintf("%d(0x%x)", v.Int(), v.Int())
		case reflect.Uint:
			typ = "UInt"
			value = fmt.Sprintf("%d(0x%x)", v.Uint(), v.Uint())
		case reflect.Uint8:
			typ = "UInt8"
			value = fmt.Sprintf("%d(0x%x)", v.Uint(), v.Uint())
		case reflect.Uint16:
			typ = "UInt16"
			value = fmt.Sprintf("%d(0x%x)", v.Uint(), v.Uint())
		case reflect.Uint32:
			typ = "UInt32"
			value = fmt.Sprintf("%d(0x%x)", v.Uint(), v.Uint())
		case reflect.Uint64:
			typ = "UInt64"
			value = fmt.Sprintf("%d(0x%x)", v.Uint(), v.Uint())
		case reflect.String:
			typ = "String"
			value = fmt.Sprintf("\"%s\"", v.String())
		case reflect.Bool:
			typ = "Bool"
			if v.Bool() {
				value = "true"
			} else {
				value = "false"
			}
		case reflect.Array:
			isIndex = true
			typ = fmt.Sprintf("Array(%s, %d)", t.Elem(), v.Len())
		case reflect.Slice:
			isIndex = true
			typ = fmt.Sprintf("Slice(%s, %d)", t.Elem(), v.Len())
		case reflect.Map:
			isMap = true
			typ = fmt.Sprintf("Map[%s]%s", t.Key(), t.Elem())
		case reflect.Func:
			typ = t.String()
		case reflect.Struct:
			isField = true
			typ = "struct"
		// case reflect.Ptr:

		default:
			fmt.Printf("unknown type: %s\n", v.Kind().String())
			return
		}
		if isIndex {
			value += "["
			isOmit := false
			ln := v.Len()
			if ln > 10 {
				ln = 10
				isOmit = true
			}
			for i := 0; i < ln; i++ {
				_, b := Inspect(v.Index(i).Interface())
				if i == 0 {
					value += fmt.Sprintf("%s", b)
				} else {
					value += fmt.Sprintf(", %s", b)
				}
			}
			if isOmit {
				value += "..."
			}
			value += "]"
		}
		if isMap {
			value += "{"
			isOmit := false
			ln := v.Len()
			if ln > 10 {
				ln = 10
				isOmit = true
			}
			keys := v.MapKeys()[0:ln]
			for i, key := range keys {
				_, keyString := Inspect(key.Interface())
				_, valString := Inspect(v.MapIndex(key).Interface())
				if i != 0 {
					value += ", "
				}
				value += fmt.Sprintf("%s -> %s", keyString, valString)
			}
			if isOmit {
				value += "..."
			}
			value += "}"
		}
		if isField {
			ln := v.NumField()
			for i := 0; i < ln; i++ {
				if v.Field(i).CanInterface() {
					a, b := Inspect(v.Field(i).Interface())
					value += fmt.Sprintf("\n  %s `%s` %s %s", t.Field(i).Name, t.Field(i).Tag, a, b)
				} else {
					value += fmt.Sprintf("\n  %s `%s` unexported", t.Field(i).Name, t.Field(i).Tag)
				}
			}
		}

		for i := 0; i < t.NumMethod(); i++ {
			if value[len(value)-1] != '\n' {
				value += "\n"
			}

			value += fmt.Sprintf("  func: %s\n", t.Method(i).Name)
		}

		if isPointer {
			typ = typ + "Ptr"
		}

		return
	}
	return "", ""
}

func Dump(data interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		fmt.Printf("%s:%d: ", file, line)
	}

	a, b := Inspect(data)
	fmt.Printf("%s:%s\n", a, b)
}
