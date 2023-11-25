package webflow

func Start(manager ContextManager, root Component) {
	app := manager.CreateContextRoot("div#app")
	root(app)
}
