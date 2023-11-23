package dom

import "syscall/js"

// Document object abstraction.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Document
type Document struct {
	Object
}

// Create a new value(element).
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (d *Document) Create(name string) js.Value {
	return d.value.Call("createElement", name)
}

func NewDocument() *Document {
	return &Document{
		Object: Object{
			value: js.Global().Get("document"),
		},
	}
}
