package flagstruct

import (
	"errors"
	"flag"
	"os"
	"reflect"
	"unsafe"
)

var (
	// ErrNotPointer returns when pass an argument to Parse function that not a pointer which point to a struct
	ErrNotPointer = errors.New("type of the first argument must be a pointer that point to a struct")
	// ErrFlagParsed returns when try to parse while flag has been parsed
	ErrFlagParsed = errors.New("flag set has been parsed, could not register more flag")
	// ErrUnsupportType returns when given struct have field that is not supported by flag
	ErrUnsupportType = errors.New("only support field types: int, int64, uint, uint64, float64, string, bool and type that implement flag.Value interface")
)

var (
	flagValueType = reflect.TypeOf((*flag.Value)(nil)).Elem()
)

// Parse properties of struct to flag,
// use default flag set, which is flag.CommandLine.
//
// Data type of field in struct must be supported by flag package:
// int, int64, uint, uint64, float64, string, bool
func Parse(i interface{}) error {
	return ParseByFlagSet(i, flag.CommandLine, os.Args[1:])
}

// ParseByFlagSet parse given flag set and arguments into struct
func ParseByFlagSet(i interface{}, fs *flag.FlagSet, args []string) error {
	if !isStructPointer(i) {
		return ErrNotPointer
	}

	if err := registerFlagStruct(i, fs); err != nil {
		return err
	}

	return fs.Parse(args)
}

// registerFlagStruct parse struct field, and register with flag set
func registerFlagStruct(i interface{}, fs *flag.FlagSet) error {
	if fs.Parsed() {
		return ErrFlagParsed
	}

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
			if !reflect.PtrTo(field.Type).Implements(flagValueType) {
				return ErrUnsupportType
			}

			fieldValue := reflect.NewAt(field.Type, fieldPtr)
			switch value := fieldValue.Interface().(type) {
			case flag.Value:
				fs.Var(value, flagName, flagUsage)
			}
		}
	}

	return nil
}
