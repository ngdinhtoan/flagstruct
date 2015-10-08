package flagstruct

import "reflect"

const tagName = "flag"

// fieldData contains StructField and Value of a property,
// that extracted from input struct
type fieldData struct {
	field reflect.StructField
	value reflect.Value
	tag   tagData
}

// structVal returns reflect.Value of struct
func structVal(i interface{}) reflect.Value {
	v := reflect.ValueOf(i)

	// if pointer get the underlying element≤
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("not struct")
	}

	return v
}

// structFields returns list of field and data which have tag flag
func structFields(i interface{}) []fieldData {
	sv := structVal(i)
	t := sv.Type()

	var f []fieldData
	for i := 0; i < t.NumField(); i = i + 1 {
		field := t.Field(i)

		// we can't access the value of unexported or anonymous fields
		if field.PkgPath != "" || field.Anonymous {
			continue
		}

		// don't check if it's omitted
		var tag string
		if tag = field.Tag.Get(tagName); tag == "-" || tag == "" {
			continue
		}

		fd := fieldData{
			field: field,
			value: sv.FieldByName(field.Name),
			tag:   parseTag(tag),
		}

		f = append(f, fd)
	}

	return f
}

// isStruct returns true if the given interface is a pointer to a struct.
func isStruct(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	} else {
		return false
	}

	// uninitialized zero value of a struct
	if v.Kind() == reflect.Invalid {
		return false
	}

	return v.Kind() == reflect.Struct
}
