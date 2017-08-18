//go:generate goagen bootstrap -d github.com/konchanSS/Todoz/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/konchanSS/Todoz/app"
)

func main() {
	// Create service
	service := goa.New("todoz")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "todos" controller
	c := NewTodosController(service)
	app.MountTodosController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
