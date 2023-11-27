package box

import (
	"Felyp-Henrique/webflow"
	"Felyp-Henrique/webflow/engine"
)

func El(ctx *webflow.Context, el string, scope engine.Scope) {
	ctx.Engine().ActionBlockCreate(el, scope)
}

func Container(ctx *webflow.Context, scope engine.Scope) {
	ctx.Engine().ActionBlockCreate("div", scope)
}
