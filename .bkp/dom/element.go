package dom

import (
	"fmt"
	"syscall/js"
)

// Struct abstraction to JavaScript's objects based in Element class.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element
type Element struct {
	Value js.Value
}

// Get element's id.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (e *Element) GetId() string {
	return e.Value.Get("id").String()
}

// Set element's id.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (e *Element) SetId(id string) {
	e.Value.Set("id", id)
}

// Get element's class name.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (e *Element) GetClass() string {
	return e.Value.Get("className").String()
}

// Set element's class name.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (e *Element) SetClass(class string) {
	e.Value.Set("className", class)
}

// Find one element by selector.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (e *Element) Get(selector string) *Element {
	return &Element{
		Value: e.Value.Call("querySelector", selector),
	}
}

// Find some elements by selector.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (e *Element) GetAll(selector string) {
	e.Value.Call("querySelectorAll", selector)
}

// Add one element into element.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Document/append
func (e *Element) Add(element Element) {
	e.Value.Call("append", element.Value)
}

// Remove on element into element.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/remove
func (e *Element) Remove(element Element) {
	e.Value.Call("remove", element.Value)
}

// Get attribute value.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (e *Element) GetAttr(name string) string {
	return e.Value.Get(fmt.Sprintf("attributes.%s", name)).String()
}

// Set attribute value.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (e *Element) SetAttr(name string, value any) {
	e.Value.Set(fmt.Sprintf("attributes.%s", name), value)
}

// Get HTML.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (e *Element) GetContent() js.Value {
	return e.Value.Get("innerHTML")
}

// Set HTML.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (e *Element) SetContent(html any) {
	e.Value.Set("innerHTML", html)
}

// Add event to element.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener
func (e *Element) AddEvent(event string, handler js.Func) {
	e.Value.Call("addEventListener", event, handler)
}

// Return de JavaScript's object.
func (e *Element) GetValue() js.Value {
	return e.Value
}
