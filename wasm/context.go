package wasm

import (
	"Felyp-Henrique/webflow"
	"fmt"
	"syscall/js"
)

type Context struct {
	_manager  *ContextManager
	_document js.Value
	_window   js.Value
	_el       js.Value
}

func (c *Context) GetId() string {
	return c._el.Get("id").String()
}

func (c *Context) SetId(id string) {
	c._el.Set("", id)
}

func (c *Context) GetClass() string {
	return c._el.Get("class").String()
}

func (c *Context) SetClass(class string) {
	c._el.Set("class", class)
}

func (c *Context) GetStyle(style string) string {
	return c._el.Get(fmt.Sprintf("style.%s", style)).String()
}

func (c *Context) SetStyle(style string, value any) {
	c._el.Set(fmt.Sprintf("style.%s", style), value)
}

func (c *Context) GetAttr(attr string) string {
	return c._el.Get(attr).String()
}

func (c *Context) SetAttr(attr string, value any) {
	c._el.Set(attr, value)
}

func (c *Context) GetText() string {
	return c._el.Get("innerText").String()
}

func (c *Context) SetText(text string) {
	c._el.Set("innerText", text)
}

func (c *Context) El(el string, scope webflow.ContextScope) {
	ctxCopy := &Context{
		_manager:  c._manager,
		_document: c._document,
		_window:   c._window,
		_el:       c._el,
	}
	child := c._manager.CreateContextChild(ctxCopy, el)
	scope(child)
}

type ContextManager struct {
	_document js.Value
	_window   js.Value
}

func (c *ContextManager) CreateContextRoot(selector string) webflow.Context {
	return &Context{
		_document: c._document,
		_window:   c._window,
		_el:       c._document.Call("querySelector", selector),
	}
}

func (c *ContextManager) CreateContextChild(ctx webflow.Context, el string) webflow.Context {
	parent := ctx.(*Context)
	target := &Context{
		_manager:  c,
		_document: c._document,
		_window:   c._window,
		_el:       c._document.Call("createElement", el),
	}
	parent._el.Call("append", target._el)
	return target
}

func New() *ContextManager {
	return &ContextManager{
		_document: js.Global().Get("document"),
		_window:   js.Global().Get("window"),
	}
}
