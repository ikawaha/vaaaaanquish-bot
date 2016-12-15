package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("vaaaaanquish-bot", func() {
	Title("vaaaaanquish-bot")
	Description("vaaaaanquish さんの名前を間違って発言すると訂正してくれる slack bot です")
	Scheme("http")
	Host("localhost:8080")
})
