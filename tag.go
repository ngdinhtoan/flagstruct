package flagstruct

import (
	"strconv"
	"strings"
)

// tagData contains information about flag name, default value, usage string
type tagData struct {
	name     string // name of flag option
	defValue string // default value of flag option
	usage    string // usage of flag option
}

func (td tagData) intValue() int {
	if td.defValue == "" {
		return 0
	}

	i64, _ := strconv.ParseInt(td.defValue, 0, 32)
	return int(i64)
}

func (td tagData) int64Value() int64 {
	if td.defValue == "" {
		return 0
	}

	i64, _ := strconv.ParseInt(td.defValue, 0, 64)
	return i64
}

func (td tagData) uintValue() uint {
	if td.defValue == "" {
		return 0
	}

	ui64, _ := strconv.ParseUint(td.defValue, 0, 32)
	return uint(ui64)
}

func (td tagData) uint64Value() uint64 {
	if td.defValue == "" {
		return 0
	}

	ui64, _ := strconv.ParseUint(td.defValue, 0, 64)
	return ui64
}

func (td tagData) stringValue() string {
	return td.defValue
}

func (td tagData) float64Value() float64 {
	if td.defValue == "" {
		return 0.0
	}

	f64, _ := strconv.ParseFloat(td.defValue, 64)
	return f64
}

func (td tagData) boolValue() bool {
	if td.defValue == "" {
		return false
	}

	b, _ := strconv.ParseBool(td.defValue)
	return b
}

func parseTag(tag string) tagData {
	td := tagData{}

	tagOpt := strings.Split(tag, ",")
	td.name = tagOpt[0]

	if len(tagOpt) > 2 {
		td.usage = tagOpt[2]
		td.defValue = tagOpt[1]
	} else if len(tagOpt) > 1 {
		td.defValue = tagOpt[1]
	}

	return td
}
