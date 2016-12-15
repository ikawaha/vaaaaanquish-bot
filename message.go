package main

import (
	"regexp"

	"github.com/goadesign/goa"
	"github.com/ikawaha/vaaaaanquish-bot/app"
)

// MessageController implements the message resource.
type MessageController struct {
	*goa.Controller
}

// NewMessageController creates a message controller.
func NewMessageController(service *goa.Service) *MessageController {
	return &MessageController{Controller: service.NewController("MessageController")}
}

var re = regexp.MustCompile(`\b[Vv]a+nquish\b`)

// Inbound runs the inbound action.
func (c *MessageController) Inbound(ctx *app.InboundMessageContext) error {
	m := re.FindAllString(ctx.Payload.Text, -1)
	ctx.Payload.Text = ""
	for _, t := range m {
		if t[1:] != "aaaaanquish" {
			ctx.Payload.Text = t[0:1] + "aaaaanquish です..."
			break
		}
	}
	res := toAppVaaaaanquishBotMessage(ctx)
	return ctx.OK(res)
}

func toAppVaaaaanquishBotMessage(ctx *app.InboundMessageContext) *app.VaaaaanquishBotMessage {
	one := 1
	botName := "VaaaaanquishBot"
	return &app.VaaaaanquishBotMessage{
		ChannelID:   ctx.Payload.ChannelID,
		ChannelName: ctx.Payload.ChannelName,
		LinkNames:   &one,
		Text:        ctx.Payload.Text,
		Token:       ctx.Payload.Token,
		UserName:    &botName,
	}
}
