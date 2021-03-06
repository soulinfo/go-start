package view

import "github.com/ungerik/go-start/utils"

///////////////////////////////////////////////////////////////////////////////
// Div

// Div represents a HTML div element.
type Div struct {
	ViewBaseWithId
	Class   string
	Content View
}

func (self *Div) IterateChildren(callback IterateChildrenCallback) {
	if self.Content != nil {
		callback(self, self.Content)
	}
}

func (self *Div) Render(context *Context, writer *utils.XMLWriter) (err error) {
	writer.OpenTag("div").Attrib("id", self.id).AttribIfNotDefault("class", self.Class)
	if self.Content != nil {
		err = self.Content.Render(context, writer)
	}
	writer.ExtraCloseTag()
	return err
}
