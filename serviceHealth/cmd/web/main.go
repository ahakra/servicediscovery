package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"

	"github.com/ahakra/servicediscovery/pkg/config"
	"github.com/ahakra/servicediscovery/serviceHealth/internal/controller"
	repository "github.com/ahakra/servicediscovery/serviceHealth/internal/repo"
	"github.com/ahakra/servicediscovery/serviceHealth/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

var GRPCClient repository.ClientInfoRepo

func main() {
	conf := config.NewFromJson("config.json")
	GRPCClient = repository.NewClientInfo(conf)
	controller.GRPCClient = &GRPCClient
	// Initialize standard Go html template engine
	rootPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: %v\n", err)

	}

	fmt.Printf("Working Directory: %s\n", rootPath)
	engine := html.New(path.Join(rootPath, "serviceHealth", "internal", "views"), ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.FiberRoutes(app)

	go func() {
		if err := app.Listen(":8001"); err != nil {
			log.Fatalf("Error: %v", err)
		}
	}()

	// Handle termination signals
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	if err := app.Shutdown(); err != nil {
		log.Fatalf("Error: %v", err)
	}

}
