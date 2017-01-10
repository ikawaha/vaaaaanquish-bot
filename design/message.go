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
	Attribute("service_id", String, "Service ID")
	Attribute("timestamp", Number, "Timestamp")
	Attribute("user_id", String, "User ID")
	Attribute("user_name", String, "User Name")
	Attribute("trigger_word", String, "Trigger Word")

	RequiredAttribute("text", String, "Message Text")
})

var _ = Resource("message", func() {
	BasePath("/v1/slack")
	DefaultMedia(MessageMedia)
	Action("inbound", func() {
		Routing(POST("/inbound"))
		Payload(SlackMessage)
		Response(OK)
	})
})

var MessageMedia = MediaType("application/vnd.vaaaaanquish.bot.message+json", func() {
	Attributes(func() {
		RequiredAttribute("text", String, "Message Text")
		Attribute("icon_url", String, "ICON URL")
		Attribute("icon_emoji", String, "ICON Emoji")
		Attribute("username", String, "User Name")
		Attribute("channel", String, "Other Channel")
	})
	View("default", func() {
		Attribute("text")
		Attribute("icon_url")
		Attribute("icon_emoji")
		Attribute("username")
		Attribute("channel")
	})
})
