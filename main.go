package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gastrader/gohouse/config"
	"github.com/gastrader/gohouse/ent"
	"github.com/gastrader/gohouse/ent/migrate"
	"github.com/gastrader/gohouse/handlers"
	"github.com/gastrader/gohouse/middleware"
	"github.com/gastrader/gohouse/routes"
	"github.com/gastrader/gohouse/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	conf := config.New()
	//connection to DB with ENT
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", conf.Database.Host, conf.Database.Port, conf.Database.User, conf.Database.Name, conf.Database.Password))
	if err != nil {
		utils.Fatalf("DB connection failed: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),

	)

	if err != nil {
		utils.Fatalf("migration failed: %v", err)
	}

	app := fiber.New()

	middleware.SetMiddleware(app)

	handler := handlers.NewHandlers(client, conf)

	routes.SetupApiV1(app, handler)

	port := "8080"

	addr := flag.String("addr", port, "http service address")
	flag.Parse()
	log.Fatal(app.Listen(":" + *addr))

}
