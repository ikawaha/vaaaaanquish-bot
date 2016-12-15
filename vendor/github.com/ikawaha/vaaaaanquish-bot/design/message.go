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
	RequiredAttribute("token", String, "Slack Token")
	RequiredAttribute("team_id", String, "Team ID")
	RequiredAttribute("team_domain", String, "Team Domain")
	RequiredAttribute("channel_id", String, "Channel ID")
	RequiredAttribute("channel_name", String, "Channel Name")
	RequiredAttribute("time_stamp", String, "Timestamp")
	RequiredAttribute("user_id", String, "User ID")
	RequiredAttribute("user_id", String, "User Name")
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
		RequiredAttribute("token", String, "Slack Token")
		RequiredAttribute("channel_id", String, "Channel ID")
		RequiredAttribute("channel_name", String, "Channel ID")
		RequiredAttribute("text", String, "Message Text")
		Attribute("user_name", String, "My bot")
		Attribute("link_names", Integer, "Find and link channel names and usernames", func() {
			Enum(1)
		})
	})
	View("default", func() {
		Attribute("token")
		Attribute("channel_id")
		Attribute("channel_name")
		Attribute("text")
		Attribute("user_name")
		Attribute("link_names")
	})
})
