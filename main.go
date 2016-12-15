//go:generate goagen bootstrap -d github.com/ikawaha/vaaaaanquish-bot/design

package main

import (
	"flag"
	"os"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ikawaha/vaaaaanquish-bot/app"
)

var addr = flag.String("addr", defaultAddr(), "server address")

func defaultAddr() string {
	if s := os.Getenv("PORT"); s != "" {
		return ":" + s
	}
	return ":8080"
}

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
	if err := service.ListenAndServe(*addr); err != nil {
		service.LogError("startup", "err", err)
	}
}
