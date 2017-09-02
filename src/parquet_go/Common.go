package parquet_go

import (
//	"log"
	"parquet"
	"reflect"
	"strings"
)

func SizeOf(val reflect.Value) int64{
	switch val.Type().Kind() {
	case reflect.Int16:
		return 2
	case reflect.Int32:
		return 4
	case reflect.Int64:
		return 8
	case reflect.Float32:
		return 4
	case reflect.Float64:
		return 8
	case reflect.Bool:
		return 1
	case reflect.String:
		return int64(val.Len())
	case reflect.Slice:
		var size int64 = 0
		for i:=0; i<val.Len(); i++{
			size += SizeOf(val.Index(i))
		}
		return size
	case reflect.Struct:
		var size int64 = 0
		numField := TypeNumberField(val.Type())
		for i:=0; int32(i)<numField; i++{
			size += SizeOf(val.Field(i))
		}
		return size
	default:
		return 4
	}
}

func PathToStr(path []string) string {
	return strings.Join(path, ".")
}

func StrToPath(str string) []string {
	return strings.Split(str, ".")
}

func TypeNumberField(t reflect.Type) int32 {
	if t.Kind() == reflect.Struct {
		return int32(t.NumField())
	} else if t.Kind() == reflect.Slice {
		return 1
	} else {
		return 0
	}
}

func GoTypeToParquetType(goT reflect.Type) parquet.Type {
	switch goT.Kind() {
	case reflect.Bool:
		return parquet.Type_BOOLEAN
	case reflect.Int:
		return parquet.Type_INT64
	case reflect.Int32:
		return parquet.Type_INT32
	case reflect.Int64:
		return parquet.Type_INT64
	case reflect.Float32:
		return parquet.Type_FLOAT
	case reflect.Float64:
		return parquet.Type_DOUBLE
	case reflect.String:
		return parquet.Type_BYTE_ARRAY
	default:
		return parquet.Type_FIXED_LEN_BYTE_ARRAY
	}
}

