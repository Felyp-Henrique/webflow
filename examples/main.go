package main

import (
	"Felyp-Henrique/webflow/dom"
	"syscall/js"
)

func main() {
	js.Global().Set("GoApp", GoApp())
	<-make(chan bool)
}

func GoApp() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		document := dom.NewDocument()
		h1 := dom.NewElement(document.Create("h1"))
		h1.SetContent("Hello")
		h2 := dom.NewElement(document.Create("h2"))
		h2.SetContent("World!")
		div := dom.NewElement(document.Create("div"))
		div.Add(h1.Object)
		div.Add(h2.Object)
		return div.GetValue()
	})
}
