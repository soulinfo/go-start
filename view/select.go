package view

import "github.com/ungerik/go-start/utils"

///////////////////////////////////////////////////////////////////////////////
// Select

type Select struct {
	ViewBaseWithId
	Model    SelectModel
	Name     string
	Size     int // 0 shows all items, 1 shows a dropdownbox, other values show size items
	Class    string
	Disabled bool
}

func (self *Select) IterateChildren(callback IterateChildrenCallback) {
	self.Model.IterateChildren(self, callback)
}

func (self *Select) Render(context *Context, writer *utils.XMLWriter) (err error) {
	writer.OpenTag("select").Attrib("id", self.id).AttribIfNotDefault("class", self.Class)
	writer.Attrib("name", self.Name)
	if self.Disabled {
		writer.Attrib("disabled", "disabled")
	}

	size := self.Size

	if self.Model != nil {
		numOptions := self.Model.NumOptions()
		if size == 0 {
			size = numOptions
		}
		writer.Attrib("size", size)

		for i := 0; i < numOptions; i++ {
			writer.OpenTag("option")
			writer.AttribIfNotDefault("value", self.Model.Value(i))
			if self.Model.Selected(i) {
				writer.Attrib("selected", "selected")
			}
			if self.Model.Disabled(i) {
				writer.Attrib("disabled", "disabled")
			}
			err = self.Model.RenderItem(i, context, writer)
			if err != nil {
				return err
			}
			writer.CloseTag() // option
		}
	} else {
		writer.Attrib("size", size)
	}

	writer.CloseTag() // select
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// Model

type SelectModel interface {
	NumOptions() int
	Value(index int) string
	Selected(index int) bool
	Disabled(index int) bool
	RenderItem(index int, context *Context, writer *utils.XMLWriter) (err error)
	IterateChildren(parent *Select, callback func(parent View, child View) (next bool))
}

///////////////////////////////////////////////////////////////////////////////
// StringsSelectModel

type StringsSelectModel struct {
	Options        []string
	SelectedOption string
}

func (self *StringsSelectModel) NumOptions() int {
	return len(self.Options)
}

func (self *StringsSelectModel) Value(index int) string {
	return self.Options[index]
}

func (self *StringsSelectModel) Selected(index int) bool {
	return self.Options[index] == self.SelectedOption
}

func (self *StringsSelectModel) Disabled(index int) bool {
	return false
}

func (self *StringsSelectModel) RenderItem(index int, context *Context, writer *utils.XMLWriter) (err error) {
	writer.Content(self.Options[index])
	return nil
}

func (self *StringsSelectModel) IterateChildren(parent *Select, callback func(parent View, child View) (next bool)) {
}

///////////////////////////////////////////////////////////////////////////////
// IndexedStringsSelectModel

type IndexedStringsSelectModel struct {
	Options []string
	Index   int
}

func (self *IndexedStringsSelectModel) NumOptions() int {
	return len(self.Options)
}

func (self *IndexedStringsSelectModel) Value(index int) string {
	return self.Options[index]
}

func (self *IndexedStringsSelectModel) Selected(index int) bool {
	return index == self.Index
}

func (self *IndexedStringsSelectModel) Disabled(index int) bool {
	return false
}

func (self *IndexedStringsSelectModel) RenderItem(index int, context *Context, writer *utils.XMLWriter) (err error) {
	writer.Content(self.Options[index])
	return nil
}

func (self *IndexedStringsSelectModel) IterateChildren(parent *Select, callback func(parent View, child View) (next bool)) {
}
