package flagstruct

import (
	"errors"
	"flag"
	"os"
	"reflect"
	"unsafe"
)

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
