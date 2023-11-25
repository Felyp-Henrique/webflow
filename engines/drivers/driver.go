package drivers

import "Felyp-Henrique/webflow/engines"

type Processor interface {
	Render(*engines.Element)
}

type ProcessorStyle interface {
	Render(engines.Style)
}

type ProcessorAttributes interface {
	Render(engines.Attributes)
}
