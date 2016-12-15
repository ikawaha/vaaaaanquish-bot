package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("public", func() {
	Files("/", "./templates/index.tmpl.html")
	Files("/static/*filepath", "./static/")
})
