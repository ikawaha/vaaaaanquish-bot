// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/ikawaha/vaaaaanquish-bot/design
// --force=true
// --out=$(GOPATH)/src/github.com/ikawaha/vaaaaanquish-bot
// --version=v1.1.0-dirty
//
// API "vaaaaanquish-bot": Application User Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import "github.com/goadesign/goa"

// slackMessage user type.
type slackMessage struct {
	// Channel ID
	ChannelID *string `form:"channel_id,omitempty" json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// Channel Name
	ChannelName *string `form:"channel_name,omitempty" json:"channel_name,omitempty" xml:"channel_name,omitempty"`
	// Team Domain
	TeamDomain *string `form:"team_domain,omitempty" json:"team_domain,omitempty" xml:"team_domain,omitempty"`
	// Team ID
	TeamID *string `form:"team_id,omitempty" json:"team_id,omitempty" xml:"team_id,omitempty"`
	// Message Text
	Text *string `form:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// Timestamp
	TimeStamp *string `form:"time_stamp,omitempty" json:"time_stamp,omitempty" xml:"time_stamp,omitempty"`
	// Slack Token
	Token *string `form:"token,omitempty" json:"token,omitempty" xml:"token,omitempty"`
	// Trigger Word
	TriggerWord *string `form:"trigger_word,omitempty" json:"trigger_word,omitempty" xml:"trigger_word,omitempty"`
	// User Name
	UserID *string `form:"user_id,omitempty" json:"user_id,omitempty" xml:"user_id,omitempty"`
}

// Validate validates the slackMessage type instance.
func (ut *slackMessage) Validate() (err error) {
	if ut.Token == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	if ut.TeamID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "team_id"))
	}
	if ut.TeamDomain == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "team_domain"))
	}
	if ut.ChannelID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "channel_id"))
	}
	if ut.ChannelName == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "channel_name"))
	}
	if ut.TimeStamp == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "time_stamp"))
	}
	if ut.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user_id"))
	}
	if ut.Text == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "text"))
	}
	return
}

// Publicize creates SlackMessage from slackMessage
func (ut *slackMessage) Publicize() *SlackMessage {
	var pub SlackMessage
	if ut.ChannelID != nil {
		pub.ChannelID = *ut.ChannelID
	}
	if ut.ChannelName != nil {
		pub.ChannelName = *ut.ChannelName
	}
	if ut.TeamDomain != nil {
		pub.TeamDomain = *ut.TeamDomain
	}
	if ut.TeamID != nil {
		pub.TeamID = *ut.TeamID
	}
	if ut.Text != nil {
		pub.Text = *ut.Text
	}
	if ut.TimeStamp != nil {
		pub.TimeStamp = *ut.TimeStamp
	}
	if ut.Token != nil {
		pub.Token = *ut.Token
	}
	if ut.TriggerWord != nil {
		pub.TriggerWord = ut.TriggerWord
	}
	if ut.UserID != nil {
		pub.UserID = *ut.UserID
	}
	return &pub
}

// SlackMessage user type.
type SlackMessage struct {
	// Channel ID
	ChannelID string `form:"channel_id" json:"channel_id" xml:"channel_id"`
	// Channel Name
	ChannelName string `form:"channel_name" json:"channel_name" xml:"channel_name"`
	// Team Domain
	TeamDomain string `form:"team_domain" json:"team_domain" xml:"team_domain"`
	// Team ID
	TeamID string `form:"team_id" json:"team_id" xml:"team_id"`
	// Message Text
	Text string `form:"text" json:"text" xml:"text"`
	// Timestamp
	TimeStamp string `form:"time_stamp" json:"time_stamp" xml:"time_stamp"`
	// Slack Token
	Token string `form:"token" json:"token" xml:"token"`
	// Trigger Word
	TriggerWord *string `form:"trigger_word,omitempty" json:"trigger_word,omitempty" xml:"trigger_word,omitempty"`
	// User Name
	UserID string `form:"user_id" json:"user_id" xml:"user_id"`
}

// Validate validates the SlackMessage type instance.
func (ut *SlackMessage) Validate() (err error) {
	if ut.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	if ut.TeamID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "team_id"))
	}
	if ut.TeamDomain == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "team_domain"))
	}
	if ut.ChannelID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "channel_id"))
	}
	if ut.ChannelName == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "channel_name"))
	}
	if ut.TimeStamp == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "time_stamp"))
	}
	if ut.UserID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "user_id"))
	}
	if ut.Text == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "text"))
	}
	return
}