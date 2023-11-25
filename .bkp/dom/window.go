package dom

import "syscall/js"

type Window struct {
	value js.Value
}

func (w *Window) Alert(text string) {
	w.value.Call("alert", text)
}

func NewWindow() *Window {
	return &Window{
		value: js.Global().Get("window"),
	}
}
