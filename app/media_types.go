// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/ikawaha/vaaaaanquish-bot/design
// --out=$(GOPATH)/src/github.com/ikawaha/vaaaaanquish-bot
// --version=v1.1.0-dirty
//
// API "vaaaaanquish-bot": Application Media Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import "github.com/goadesign/goa"

// VaaaaanquishBotMessage media type (default view)
//
// Identifier: application/vnd.vaaaaanquish.bot.message+json; view=default
type VaaaaanquishBotMessage struct {
	// Message Text
	Text string `form:"text" json:"text" xml:"text"`
}

// Validate validates the VaaaaanquishBotMessage media type instance.
func (mt *VaaaaanquishBotMessage) Validate() (err error) {
	if mt.Text == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "text"))
	}
	return
}
