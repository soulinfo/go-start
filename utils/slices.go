package utils

import (
	"github.com/ungerik/go-start/errs"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func DeleteEmptySliceElementsVal(sliceVal reflect.Value) reflect.Value {
	if sliceVal.Kind() != reflect.Slice {
		panic("Argument is not a slice: " + sliceVal.String())
	}
	zeroVal := reflect.Zero(sliceVal.Type().Elem())
	for i := 0; i < sliceVal.Len(); i++ {
		elemVal := sliceVal.Index(i)
		if reflect.DeepEqual(elemVal.Interface(), zeroVal.Interface()) {
			before := sliceVal.Slice(0, i)
			after := sliceVal.Slice(i+1, sliceVal.Len())
			sliceVal = reflect.AppendSlice(before, after)
			i--
		}
	}
	return sliceVal
}

func DeleteEmptySliceElements(slice interface{}) interface{} {
	return DeleteEmptySliceElementsVal(reflect.ValueOf(slice)).Interface()
}

func SliceInsert(slice []interface{}, index int, count int, value interface{}) (result []interface{}) {
	switch {
	case count < 0:
		panic(errs.Format("Negative count %d not allowed", count))
	case count == 0:
		return slice
	}

	length := len(slice)
	errs.PanicIfIndexOutOfBounds("SliceInsert", index, length)

	result = make([]interface{}, length+count)
	copy(result, slice[:index])
	copy(result[index+count:], slice[index:])
	for i := index; i < index+count; i++ {
		result[i] = value
	}

	return result
}

func SliceDelete(slice []interface{}, index int, count int) (result []interface{}) {
	switch {
	case count < 0:
		panic(errs.Format("Negative count %d not allowed", count))
	case count == 0:
		return slice
	}

	length := len(slice)
	errs.PanicIfIndexOutOfBounds("SliceDelete", index, length)

	if index+count > length {
		count = length - index
	}

	return append(slice[:index], slice[index+count:]...)
}

// Implements sort.Interface
type Sortable struct {
	Slice    []interface{}
	LessFunc func(a, b interface{}) bool
}

func (self *Sortable) Len() int {
	return len(self.Slice)
}

func (self *Sortable) Less(i, j int) bool {
	return self.LessFunc(self.Slice[i], self.Slice[j])
}

func (self *Sortable) Swap(i, j int) {
	self.Slice[i], self.Slice[j] = self.Slice[j], self.Slice[i]
}

func (self *Sortable) Sort() {
	sort.Sort(self)
}

func Sort(slice []interface{}, lessFunc func(a, b interface{}) bool) {
	sortable := Sortable{slice, lessFunc}
	sortable.Sort()
}

/*
func CloneStringSlice(original []string) (clone []string) {
	if original != nil {
		clone = make([]string, len(original))
		for i := range original {
			clone[i] = original[i]
		}
	}
	return clone
}


func CloneByteSlice(original []byte) (clone []byte) {
	if original != nil {
		clone = make([]byte, len(original))
		for i := range original {
			clone[i] = original[i]
		}
	}
	return clone
}
*/

func MakeVersionTuple(fields ...int) VersionTuple {
	t := make(VersionTuple, len(fields))
	for i := range fields {
		t[i] = fields[i]
	}
	return t
}

func ParseVersionTuple(s string) (VersionTuple, error) {
	fields := strings.Split(s, ".")
	t := make(VersionTuple, len(fields))
	for i := range fields {
		value, err := strconv.ParseInt(fields[i], 10, 32)
		if err != nil {
			return nil, err
		}
		t[i] = int(value)
	}
	return t, nil
}

type VersionTuple []int

func (self VersionTuple) GreaterEqual(other VersionTuple) bool {
	for i := range other {
		var value int
		if i < len(self) {
			value = self[i]
		}
		if value > other[i] {
			return true
		} else if value < other[i] {
			return false
		}
	}
	return true
}

func (self VersionTuple) String() string {
	var sb StringBuilder
	for i := range self {
		if i > 0 {
			sb.Byte('.')
		}
		sb.Int(self[i])
	}
	return sb.String()
}
