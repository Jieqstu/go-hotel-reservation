package main

import (
	"flag"

	"github.com/Jieqstu/go-hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", "5010", "The listrn address of the API server")
	flag.Parse()
	
    app := fiber.New()

	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user:id", api.HandleGetUserById)

	// dereference the pointer to string
    app.Listen(*listenAddr)

}