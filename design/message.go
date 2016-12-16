package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

func RequiredAttribute(name string, args ...interface{}) {
	Attribute(name, args...)
	Required(name)
}

var SlackMessage = Type("SlackMessage", func() {
	Attribute("token", String, "Slack Token")
	Attribute("team_id", String, "Team ID")
	Attribute("team_domain", String, "Team Domain")
	Attribute("channel_id", String, "Channel ID")
	Attribute("channel_name", String, "Channel Name")
	Attribute("timestamp", String, "Timestamp")
	Attribute("user_id", String, "User ID")
	Attribute("user_name", String, "User Name")
	RequiredAttribute("text", String, "Message Text")
	Attribute("trigger_word", String, "Trigger Word")
})

var _ = Resource("message", func() {
	BasePath("/v1/slack")
	Action("inbound", func() {
		Routing(POST("/inbound"))
		Payload(SlackMessage)
		Response(OK, MessageMedia)
	})
})

var MessageMedia = MediaType("application/vnd.vaaaaanquish.bot.message+json", func() {
	Attributes(func() {
		RequiredAttribute("text", String, "Message Text")
	})
	View("default", func() {
		Attribute("text")
	})
})
