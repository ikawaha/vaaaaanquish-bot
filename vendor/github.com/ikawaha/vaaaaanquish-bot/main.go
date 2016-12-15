//go:generate goagen bootstrap -d github.com/ikawaha/vaaaaanquish-bot/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ikawaha/vaaaaanquish-bot/app"
)

func main() {
	// Create service
	service := goa.New("vaaaaanquish-bot")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "message" controller
	c := NewMessageController(service)
	app.MountMessageController(service, c)
	// Mount "public" controller
	c2 := NewPublicController(service)
	app.MountPublicController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
