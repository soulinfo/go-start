package view

import "github.com/ungerik/go-start/utils"

///////////////////////////////////////////////////////////////////////////////
// If

type If struct {
	ViewBaseWithId
	Condition   bool
	Content     View
	ElseContent View
}

func (self *If) Init(thisView View) {
	self.ViewBaseWithId.Init(thisView)

	// ViewBaseWithId.Init() initializes the child reported by IterateChildren(),
	// we need to initialize the child for the other case of !self.Condition
	var child View
	if !self.Condition {
		child = self.Content
	} else {
		child = self.ElseContent
	}
	if child != nil {
		child.Init(child)
	}
}

func (self *If) IterateChildren(callback IterateChildrenCallback) {
	var child View
	if self.Condition {
		child = self.Content
	} else {
		child = self.ElseContent
	}
	if child != nil {
		callback(self, child)
	}
}

func (self *If) Render(context *Context, writer *utils.XMLWriter) (err error) {
	content := self.Content
	if !self.Condition {
		content = self.ElseContent
	}
	if content == nil {
		return nil
	}
	return content.Render(context, writer)
}
