package dom

import "syscall/js"

// Document object abstraction.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Document
type Document struct {
	Value js.Value
}

// Find one element by selector.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
func (d *Document) Get(selector string) *Element {
	return &Element{
		Value: d.Value.Call("querySelector", selector),
	}
}

// Create a new value(element).
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (d *Document) Create(name string) *Element {
	return &Element{
		Value: d.Value.Call("createElement", name),
	}
}

func NewDocument() *Document {
	return &Document{
		Value: js.Global().Get("document"),
	}
}
