package dom

import (
	"fmt"
	"syscall/js"
)

// Element object abstraction.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element
type Element struct {
	Object
}

// Get element's id.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (e *Element) GetId() string {
	return e.value.Get("id").String()
}

// Set element's id.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (e *Element) SetId(id string) {
	e.value.Set("id", id)
}

// Get element's class name.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (e *Element) GetClass() string {
	return e.value.Get("className").String()
}

// Set element's class name.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (e *Element) SetClass(class string) {
	e.value.Set("className", class)
}

// Get attribute value.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (e *Element) GetAttr(name string) string {
	return e.value.Get(fmt.Sprintf("attributes.%s", name)).String()
}

// Set attribute value.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (e *Element) SetAttr(name string, value any) {
	e.value.Set(fmt.Sprintf("attributes.%s", name), value)
}

// Get HTML.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (e *Element) GetContent() js.Value {
	return e.value.Get("innerHTML")
}

// Set HTML.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (e *Element) SetContent(html any) {
	e.value.Set("innerHTML", html)
}

// Add event to element.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (e *Element) AddEvent(event string, handler js.Func) {
	e.value.Call("addEventListener", event, handler)
}

func (e *Element) GetValue() js.Value {
	return e.value
}

func NewElement(value js.Value) *Element {
	return &Element{
		Object: Object{
			value: value,
		},
	}
}
