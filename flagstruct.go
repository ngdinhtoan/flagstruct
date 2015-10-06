package flagstruct

import (
	"errors"
	"flag"
	"os"
	"reflect"
	"unsafe"
)

const tagName = "flag"

// fieldData contains StructField and Value of a property,
// that extracted from input struct
type fieldData struct {
	field reflect.StructField
	value reflect.Value
	tag   tagData
}

// Parse properties of struct to flag,
// use default flag set, which is flag.CommandLine
func Parse(i interface{}) error {
	return parseByFlagSet(i, flag.CommandLine, os.Args[1:])
}

func parseByFlagSet(i interface{}, fs *flag.FlagSet, args []string) error {
	if !isStruct(i) {
		return errors.New("type of the first argument must be a pointer that point to a struct")
	}

	registerFlagStruct(i, fs)
	return fs.Parse(args)
}

// registerFlagStruct parse struct field, and register with flag set
func registerFlagStruct(i interface{}, fs *flag.FlagSet) {
	sf := structFields(i)
	for _, fd := range sf {
		field := fd.field

		flagName := fd.tag.name
		flagUsage := fd.tag.usage
		fieldPtr := unsafe.Pointer(fd.value.UnsafeAddr())

		switch field.Type.Kind() {
		case reflect.Int:
			fs.IntVar((*int)(fieldPtr), flagName, fd.tag.intValue(), flagUsage)
		case reflect.Int64:
			fs.Int64Var((*int64)(fieldPtr), flagName, fd.tag.int64Value(), flagUsage)
		case reflect.Uint:
			fs.UintVar((*uint)(fieldPtr), flagName, fd.tag.uintValue(), flagUsage)
		case reflect.Uint64:
			fs.Uint64Var((*uint64)(fieldPtr), flagName, fd.tag.uint64Value(), flagUsage)
		case reflect.String:
			fs.StringVar((*string)(fieldPtr), flagName, fd.tag.stringValue(), flagUsage)
		case reflect.Bool:
			fs.BoolVar((*bool)(fieldPtr), flagName, fd.tag.boolValue(), flagUsage)
		case reflect.Float64:
			fs.Float64Var((*float64)(fieldPtr), flagName, fd.tag.float64Value(), flagUsage)
		default:
			panic("only support field types: int, int64, uint, uint64, float64, string and bool.")
		}
	}
}

func strctVal(i interface{}) reflect.Value {
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

func structFields(i interface{}) []fieldData {
	sv := strctVal(i)
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

// isStruct returns true if the given interface is a pointer to struct.
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
