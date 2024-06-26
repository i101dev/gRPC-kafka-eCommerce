package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/i101dev/gRPC-kafka-eCommerce/auth"
	"github.com/i101dev/gRPC-kafka-eCommerce/kafka"
	pb "github.com/i101dev/gRPC-kafka-eCommerce/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port     string
	adminKey string

	orderSrv    string
	orderConn   *grpc.ClientConn
	orderClient pb.OrderServiceClient

	productSrv    string
	productConn   *grpc.ClientConn
	productClient pb.ProductServiceClient

	userSrv    string
	userConn   *grpc.ClientConn
	userClient pb.UserServiceClient

	KAFKA_URI string
)

func loadENV() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file")
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	adminKey = os.Getenv("ADMIN_KEY")
	if adminKey == "" {
		log.Fatal("Invalid [ADMIN_KEY] - not found in [.env]")
	}

	KAFKA_URI = os.Getenv("KAFKA_URI")
	if KAFKA_URI == "" {
		log.Fatal("Invalid [KAFKA_URI] - not found in [.env]")
	}

	orderSrv = os.Getenv("ORDER_SRV")
	if orderSrv == "" {
		log.Fatal("Invalid [ORDER_SRV] - not found in [.env]")
	}
	productSrv = os.Getenv("PRODUCT_SRV")
	if productSrv == "" {
		log.Fatal("Invalid [PRODUCT_SRV] - not found in [.env]")
	}
	userSrv = os.Getenv("USER_SRV")
	if userSrv == "" {
		log.Fatal("Invalid [USER_SRV] - not found in [.env]")
	}
}

func loadGRPC() {

	// --------------------------------------------------------------------------
	// order service
	//
	orderConn, err := grpc.Dial(orderSrv, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial the [order] server %v", err)
	} else {
		orderClient = pb.NewOrderServiceClient(orderConn)
	}

	// --------------------------------------------------------------------------
	// inventory service
	//
	productConn, err := grpc.Dial(productSrv, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial the [product] server %v", err)
	} else {
		productClient = pb.NewProductServiceClient(productConn)
	}

	// --------------------------------------------------------------------------
	// user service
	//
	userConn, err := grpc.Dial(userSrv, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial the [user] server %v", err)
	} else {
		userClient = pb.NewUserServiceClient(userConn)
	}

	connectServices()
}

func connectServices() {

	if userClient != nil && productClient != nil && orderClient != nil {

		// --------------------------------------------------------------------------
		// User service
		userConnReq := &pb.UserConnReq{
			ProductSrv: productSrv,
			OrderSrv:   orderSrv,
		}
		userConnRes, userConnErr := userClient.UserConn(context.Background(), userConnReq)
		if userConnErr != nil {
			fmt.Printf("\n*** >>> User service failed to dial others - %v", userConnErr)
		} else {
			fmt.Printf("\n*** >>> %+v", userConnRes)
		}

		// --------------------------------------------------------------------------
		// Product service
		productConnReq := &pb.ProductConnReq{
			UserSrv:  userSrv,
			OrderSrv: orderSrv,
		}
		productConnRes, productConnErr := productClient.ProductConn(context.Background(), productConnReq)
		if productConnErr != nil {
			fmt.Printf("\n*** >>> Product service failed to dial others - %v", productConnErr)
		} else {
			fmt.Printf("\n*** >>> %+v", productConnRes)
		}

		// // --------------------------------------------------------------------------
		// Order service
		orderConnReq := &pb.OrderConnReq{
			ProductSrv: productSrv,
			UserSrv:    userSrv,
		}
		orderConnRes, orderConnErr := orderClient.OrderConn(context.Background(), orderConnReq)
		if orderConnErr != nil {
			fmt.Printf("\n*** >>> Order service failed to dial others - %v", orderConnErr)
		} else {
			fmt.Printf("\n*** >>> %+v", orderConnRes)
		}

	} else {
		fmt.Println("*** >>> One or more of the services is [nil]")
	}
}

func fiberApp() *fiber.App {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Fiber error": err.Error(),
			})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	return app
}

func main() {

	loadENV()
	loadGRPC()

	defer orderConn.Close()
	defer productConn.Close()
	defer userConn.Close()

	app := fiberApp()

	quitCh := make(chan os.Signal, 1)

	go kafka.StartConsumer("orders", KAFKA_URI, quitCh)
	go kafka.StartConsumer("products", KAFKA_URI, quitCh)
	go kafka.StartConsumer("users", KAFKA_URI, quitCh)

	// --------------------------------------------------------------------------
	// Routers
	app.Get("/", GET_index)
	app.Post("/order/test", POST_order_test)
	app.Post("/product/test", POST_product_test)
	app.Post("/user/test", POST_user_test)

	app.Post("/auth", POST_UserAuth)
	app.Post("/join", POST_UserJoin)

	// --------------------------------------------------------------------------
	// User API
	api_user := app.Group("/user")
	api_user.Use(auth.ValidateJWT)
	api_user.Get("/products", auth.RequireRole("cust"), GET_products)
	api_user.Get("/orders", auth.RequireRole("cust"), GET_orders)

	// --------------------------------------------------------------------------
	// Admin API
	api_admin := app.Group("/admin")
	api_admin.Use(auth.ValidateJWT)
	api_admin.Get("/users", auth.RequireRole("admin"), GET_users)
	api_admin.Get("/orders", auth.RequireRole("admin"), GET_orders)

	// --------------------------------------------------------------------------
	// Launch server
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
