package dom

import "syscall/js"

func RenderWasm(start func()) {
	js.Global().Set("WebflowApp", js.FuncOf(func(this js.Value, args []js.Value) any {
		start()
		return nil
	}))

	<-make(chan bool)
}
