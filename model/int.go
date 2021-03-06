package model

import (
	"fmt"
	"strconv"
)

// Attributes:
// * min
// * max
// * label
type Int int64

func (self *Int) Get() int64 {
	return int64(*self)
}

func (self *Int) Set(value int64) {
	*self = Int(value)
}

func (self *Int) String() string {
	return strconv.FormatInt(self.Get(), 10)
}

func (self *Int) SetString(str string) error {
	value, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		self.Set(value)
	}
	return err
}

func (self *Int) IsEmpty() bool {
	return false
}

func (self *Int) Validate(metaData *MetaData) []*ValidationError {
	value := int64(*self)
	errors := NoValidationErrors
	min, ok, err := self.Min(metaData)
	if err != nil {
		errors = append(errors, &ValidationError{err, metaData})
	} else if ok && value < min {
		errors = append(errors, &ValidationError{&IntBelowMin{value, min}, metaData})
	}
	max, ok, err := self.Max(metaData)
	if err != nil {
		errors = append(errors, &ValidationError{err, metaData})
	} else if ok && value > max {
		errors = append(errors, &ValidationError{&IntAboveMax{value, max}, metaData})
	}
	return errors
}

func (self *Int) Min(metaData *MetaData) (min int64, ok bool, err error) {
	str, ok := metaData.Attrib("min")
	if !ok {
		return 0, false, nil
	}
	value, err := strconv.ParseInt(str, 10, 64)
	return value, err == nil, err
}

func (self *Int) Max(metaData *MetaData) (max int64, ok bool, err error) {
	str, ok := metaData.Attrib("max")
	if !ok {
		return 0, false, nil
	}
	value, err := strconv.ParseInt(str, 10, 64)
	return value, err == nil, err
}

type IntBelowMin struct {
	Value int64
	Min   int64
}

func (self *IntBelowMin) Error() string {
	return fmt.Sprintf("Int %d below minimum of %d", self.Value, self.Min)
}

type IntAboveMax struct {
	Value int64
	Max   int64
}

func (self *IntAboveMax) Error() string {
	return fmt.Sprintf("Int %d above maximum of %d", self.Value, self.Max)
}

func (self *Int) Hidden(metaData *MetaData) (hidden bool) {
	return metaData.BoolAttrib("hidden")
}
