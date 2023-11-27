package main

import (
	"Felyp-Henrique/webflow"
	"Felyp-Henrique/webflow/elements/box"
	"Felyp-Henrique/webflow/elements/btn"
	"Felyp-Henrique/webflow/elements/txt"
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
	box.Container(ctx, func(eh engine.ElementHandler) {
		eh.SetId("")
		eh.SetClass("")

		txt.Text(ctx, "Hello World!")

		box.El(ctx, "section", func(eh engine.ElementHandler) {
			txt.El(ctx, "span", "icon_smile")

			btn.Button(ctx, "Click to Hello!", func() {
				ctx.Window().Alert("Hello World!")
			})
		})
	})
}
