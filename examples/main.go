package main

import (
	"Felyp-Henrique/webflow/components"
	"Felyp-Henrique/webflow/dom"
	"syscall/js"
)

type CountComponent struct {
	Number int
}

func (c *CountComponent) Render() *dom.Element {
	var (
		pCounter *dom.Element = dom.El("p", dom.WithContent(0))
	)
	return dom.El("div", dom.WithClass("component-count"), dom.WithChildren(
		dom.El("button", dom.WithContent("+"), dom.WithEvent("click", js.FuncOf(func(this js.Value, args []js.Value) any {
			c.Number += 1
			pCounter.SetContent(c.Number)
			return nil
		}))),
		pCounter,
		dom.El("button", dom.WithContent("-"), dom.WithEvent("click", js.FuncOf(func(this js.Value, args []js.Value) any {
			c.Number -= 1
			pCounter.SetContent(c.Number)
			return nil
		}))),
	))
}

func main() {
	js.Global().Set("GoApp", GoApp())
	js.Global().Set("GoAppWithDom", GoAppWithDom())
	<-make(chan bool)
}

func GoApp() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		return components.Render("div#app", &CountComponent{})
	})
}

func GoAppWithDom() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		h1 := dom.El("h1",
			dom.WithId("hello"),
			dom.WithContent("Hello"))
		h2 := dom.El("h2",
			dom.WithId("world"),
			dom.WithContent("World!"))
		button := dom.El("button",
			dom.WithContent("Say hello!"),
			dom.WithEvent("click", js.FuncOf(func(this js.Value, args []js.Value) any {
				dom.Win().Alert("Hello Guys!")
				h1.SetContent("Hello World!")
				h2.SetContent("Hello Guys!")
				return true
			})))
		div := dom.El("div",
			dom.WithId("contents"),
			dom.WithChildren(h1, h2, button))
		return div.GetValue()
	})
}
