package webflow

type Context interface {
	GetId() string

	SetId(string)

	GetClass() string

	SetClass(string)

	GetStyle(string) string

	SetStyle(string, any)

	GetAttr(string) string

	SetAttr(string, any)

	GetText() string

	SetText(string)

	El(string, ContextScope)
}

type ContextScope func(Context)

type ContextManager interface {
	CreateContextRoot(string) Context

	CreateContextChild(Context, string) Context
}
