package data

import (
	"time"
)

type Comparable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

type Equatable interface {
	Comparable | bool | time.Time
}

type Type string

const (
	IntDataType      Type = "int"
	DecimalDataType  Type = "decimal"
	BoolDataType     Type = "bool"
	StringDataType   Type = "string"
	RuneDataType     Type = "rune"
	DatetimeDataType Type = "datetime"
)

func GetType(value interface{}) Type {
	switch value.(type) {
	case int8, int16, int, int64, uint8, uint16, uint32, uint64:
		return IntDataType
	case float32, float64:
		return DecimalDataType
	case bool:
		return BoolDataType
	case string:
		return StringDataType
	case rune:
		return RuneDataType
	case time.Time:
		return DatetimeDataType
	default:
		return StringDataType
	}
}
