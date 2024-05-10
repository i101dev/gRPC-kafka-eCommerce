package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"

	"github.com/i101dev/gRPC-kafka-eCommerce/auth"
)

func main() {

	if err := godotenv.Load("../config/.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// --------------------------------------------------------------------------
	// INITIALIZE FIBER
	//
	app := fiber.New(fiber.Config{
		ErrorHandler: fiberErrorHandler,
	})

	// --------------------------------------------------------------------------
	// MIDDLEWARES
	//
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// --------------------------------------------------------------------------
	// ROUTES
	//

	app.Get("/", Handle_get_index)
	app.Post("/login", Handle_post_login)

	// --------------------------------------------------------------------------
	// Customer API
	//
	api_cust := app.Group("/api")
	api_cust.Use(auth.ValidateJWT)
	api_cust.Get("/products", Handle_get_products)
	api_cust.Get("/inventory", Handle_get_inventory)

	// --------------------------------------------------------------------------
	// Admin API
	//
	api_admin := app.Group("/admin")
	api_admin.Use(auth.ValidateJWT)
	api_admin.Get("/users", auth.RequireRole("admin"), Handle_get_users)
	api_admin.Get("/orders", auth.RequireRole("admin"), haandl_get_orders)

	// --------------------------------------------------------------------------
	// Launch server
	//
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

func fiberErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"Fiber error": err.Error(),
	})
}
