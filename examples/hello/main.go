package main

import (
	"Felyp-Henrique/webflow"
	"Felyp-Henrique/webflow/wasm"
	"syscall/js"
)

func main() {
	js.Global().Set("App", js.FuncOf(func(this js.Value, args []js.Value) any {
		webflow.Start(wasm.New(), func(c webflow.Context) {
			c.El("div", func(div webflow.Context) {
				div.El("h1", func(h1 webflow.Context) {
					h1.SetText("Hello World!")
				})
				div.El("h2", func(h2 webflow.Context) {
					h2.SetText("Testing web assembly Golang!")
				})
			})
		})
		return nil
	}))
	<-make(chan bool)
}
