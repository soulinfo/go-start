package model

import "github.com/ungerik/go-mail"

type Email string

func (self *Email) Get() string {
	return string(*self)
}

func (self *Email) Set(value string) (err error) {
	if value != "" {
		if value, err = email.ValidateAddress(value); err != nil {
			return err
		}
	}
	*self = Email(value)
	return nil
}

func (self *Email) IsEmpty() bool {
	return len(*self) == 0
}

func (self *Email) String() string {
	return self.Get()
}

func (self *Email) SetString(str string) (err error) {
	return self.Set(str)
}

func (self *Email) FixValue(metaData *MetaData) {
}

func (self *Email) Validate(metaData *MetaData) []*ValidationError {
	str := self.Get()
	if self.Required(metaData) || str != "" {
		if _, err := email.ValidateAddress(str); err != nil {
			return NewValidationErrors(err, metaData)
		}
	}
	return NoValidationErrors
}

func (self *Email) Required(metaData *MetaData) bool {
	return metaData.BoolAttrib("required")
}
