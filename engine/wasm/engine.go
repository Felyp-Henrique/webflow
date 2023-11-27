package wasm

import (
	"Felyp-Henrique/webflow/engine"
	"syscall/js"
)

type Engine struct {
	_factory engine.ElementFactory
	_handler *engine.Stack[engine.ElementHandler]
	_window  engine.WindowHandler
}

func (e *Engine) Factory() engine.ElementFactory {
	return e._factory
}

func (e *Engine) Handler() engine.ElementHandler {
	return e._handler.Peek()
}

func (e *Engine) Window() engine.WindowHandler {
	return e._window
}

func (e *Engine) ActionInlineCreate(el string) engine.ElementHandler {
	return e.Factory().Create(el)
}

func (e *Engine) ActionBlockCreate(el string, scope engine.Scope) engine.ElementHandler {
	var (
		elementParent engine.ElementHandler = e._handler.Peek()
		elementNew    engine.ElementHandler = e.Factory().CreateChild(elementParent, el)
	)
	e._handler.Push(elementNew)
	scope(elementNew)
	e._handler.Pop()
	return elementNew
}

func New() engine.Engine {
	return &Engine{
		_factory: &ElementFactory{
			_document: js.Global().Get("document"),
		},
		_handler: engine.NewStack[engine.ElementHandler](),
		_window: &WindowHandler{
			_value: js.Global().Get("window"),
		},
	}
}

type ElementHandler struct {
	_parent engine.ElementHandler
	_value  js.Value
}

func (e *ElementHandler) Parent() engine.ElementHandler {
	return e._parent
}

func (e *ElementHandler) SetId(id string) {
	e._value.Set("id", id)
}

func (e *ElementHandler) SetClass(class string) {
	e._value.Set("class", class)
}

func (e *ElementHandler) SetContent(content string) {
	e._value.Set("innerText", content)
}

type ElementFactory struct {
	_document js.Value
}

func (e *ElementFactory) Create(el string) engine.ElementHandler {
	return &ElementHandler{
		_value: e._document.Call("createElement", el),
	}
}

func (e *ElementFactory) CreateChild(parent engine.ElementHandler, el string) engine.ElementHandler {
	var (
		elementNew    js.Value        = e._document.Call("createElement", el)
		elementParent *ElementHandler = parent.(*ElementHandler)
	)
	elementParent._value.Call("append", elementNew)
	return &ElementHandler{
		_parent: elementParent,
		_value:  elementNew,
	}
}

type WindowHandler struct {
	_value js.Value
}

func (w *WindowHandler) Alert(message string) {
	w._value.Call("alert", message)
}
