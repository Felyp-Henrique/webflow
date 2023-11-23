package dom

import "syscall/js"

type Object struct {
	value js.Value
}

// Find one element by selector.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (o *Object) Get(selector string) {
	o.value.Call("querySelector", selector)
}

// Find some elements by selector.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (o *Object) GetAll(selector string) {
	o.value.Call("querySelectorAll", selector)
}

// Add one object into object.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Document/append
func (o *Object) Add(object Object) {
	o.value.Call("append", object.value)
}

// Remove on object into object.
//
// @see https://developer.mozilla.org/en-US/docs/Web/API/Element/remove
func (o *Object) Remove(object Object) {
	o.value.Call("remove", object.value)
}
