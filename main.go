//go:generate goagen bootstrap -d github.com/konchanSS/Todoz/design

package main

import (
	"flag"
	"log"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/konchanSS/Todoz/app"
	"github.com/konchanSS/Todoz/controller"
	"github.com/konchanSS/Todoz/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create service
	service := goa.New("todoz")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	var (
		port  = flag.String("port", ":8080", "addr to bind")
		env   = flag.String("env", "development", "application envirionment (production, development etc.)")
		dbrun = flag.Bool("dbrun", false, "database run mode")
	)
	flag.Parse()

	// Mount "todos" controller
	if *dbrun {
		cs, err := database.NewConfigsFromFile("dbconfig.yml")
		if err != nil {
			log.Fatalf("cannot open database configuration. exit. %s", err)
		}
		dbcon, err := cs.Open(*env)
		if err != nil {
			log.Fatalf("database initialization failed: %s", err)
		}
		c := controller.NewTodosController(service, dbcon)
		app.MountTodosController(service, c)
	}

	// Start service
	if err := service.ListenAndServe(*port); err != nil {
		service.LogError("startup", "err", err)
	}

}
