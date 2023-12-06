package main

import (
	"Felyp-Henrique/webflow"
	"Felyp-Henrique/webflow/els"
	"Felyp-Henrique/webflow/engine"
	"Felyp-Henrique/webflow/engine/wasm"
)

func main() {
	engine := wasm.New()
	webflow := webflow.NewWebflow(engine)
	webflow.Start(App)
}

func App(ctx *webflow.Context) {
	HelloComponent(ctx)
}

func HelloComponent(ctx *webflow.Context) {
	els.Section(ctx, func(eh engine.ElementHandler) {
		els.Title(ctx, els.Title1, "Hello World!")

		els.Text(ctx, "Hello World!")

		els.Container(ctx, func(eh engine.ElementHandler) {
			els.Span(ctx, "icon_smile")

			els.Button(ctx, "Click to Hello!", func() {
				ctx.Window().Alert("Hello World!")
			})
		})
	})
}
