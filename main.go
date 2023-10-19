package main

import (
	"context"
	"flag"
	"log"

	"github.com/Jieqstu/go-hotel-reservation/api"
	"github.com/Jieqstu/go-hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"

var config = fiber.Config {
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", "5010", "The listrn address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	// user := types.User{
	// 	FirstName: "max",
	// 	LastName: "verstappen",
	// }
	// newStore := db.NewMongoUserStore(client)
	// res, err := newStore.GetColl().InsertOne(context.Background(), user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(res)

	// handlers initialization
	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

    app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUserById)

	// dereference the pointer to string
    app.Listen(*listenAddr)

}