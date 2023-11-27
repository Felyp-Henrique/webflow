package main

import (
	w "Felyp-Henrique/webflow"
	e "Felyp-Henrique/webflow/engine"
	"Felyp-Henrique/webflow/engine/wasm"
)

func main() {
	engine := wasm.New()
	webflow := w.NewWebflow(engine)
	webflow.Start(App)
}

func App(ctx *w.Context) {
	HelloComponent(ctx)
}

func HelloComponent(ctx *w.Context) {
	w.Container(ctx, func(eh e.ElementHandler) {
		eh.SetId("")
		eh.SetClass("")

		w.Text(ctx, "Hello World!")

		w.Button(ctx, "Alert", func() {
			ctx.Window().Alert("Hello World!")
		})
	})
}
