package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	pb "github.com/i101dev/gRPC-kafka-eCommerce/proto"
)

var (
	inventoryDB *gorm.DB

	ADMIN_KEY          string
	INVENTORY_SRV_HOST string
	INVENTORY_SRV_PORT string
)

type InventoryServer struct {
	pb.InventoryServiceServer
}

func GetDB() *gorm.DB {
	return inventoryDB
}

func loadENV() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading [inventory-service] .env file")
	}

	ADMIN_KEY = os.Getenv("ADMIN_KEY")
	if ADMIN_KEY == "" {
		log.Fatal("Invalid [ADMIN_KEY] - not found in [.env]")
	}

	INVENTORY_SRV_HOST = os.Getenv("INVENTORY_SRV_HOST")
	if INVENTORY_SRV_HOST == "" {
		log.Fatal("Invalid [INVENTORY_SRV_HOST] - not found in [.env]")
	}

	INVENTORY_SRV_PORT = os.Getenv("INVENTORY_SRV_PORT")
	if INVENTORY_SRV_PORT == "" {
		log.Fatal("Invalid [INVENTORY_SRV_PORT] - not found in [.env]")
	}
}

func loadDB() {

	dbUser := os.Getenv("INVENTORY_DB_USER")
	dbPass := os.Getenv("INVENTORY_DB_PASS")
	dbName := os.Getenv("INVENTORY_DB_NAME")
	dbHost := os.Getenv("INVENTORY_DB_HOST")
	dbPort := os.Getenv("INVENTORY_DB_PORT")

	if dbUser == "" || dbPass == "" || dbName == "" || dbHost == "" || dbPort == "" {
		log.Fatalf("incomplete database connection parameters")
	}

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to [inventory-service] database: %+v", err)
	} else {
		inventoryDB = db
		InitModels(db)
	}
}

func loadSRV() {

	lis, err := net.Listen("tcp", ":"+INVENTORY_SRV_PORT)

	if err != nil {
		log.Fatalf("Failed to start the [inventory-gRPC] %+v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterInventoryServiceServer(grpcServer, &InventoryServer{})

	log.Printf("*** >>> [inventory-gRPC] server started at %+v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("*** >>> [inventory-gRPC] failed to start - %+v", err)
	}
}

func main() {
	loadENV()
	loadDB()
	loadSRV()
}