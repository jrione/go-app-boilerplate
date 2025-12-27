package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/jrione/go-app-boilerplate/plugin"
	"github.com/jrione/go-app-boilerplate/proto"
	"github.com/jrione/go-app-boilerplate/routes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Initialize config
	config := plugin.NewConfig()

	// Initialize logger
	logger := plugin.NewLogger()

	// Initialize database
	db := plugin.NewDatabase(config, logger)

	// Gin router for REST API
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, logger, db)

	// Start Gin server in a goroutine
	go func() {
		addr := fmt.Sprintf(":%s", config.GetString("api_port"))
		logger.Info("Starting REST API server on ", addr)
		if err := r.Run(addr); err != nil {
			logger.Fatal("Failed to start REST API server: ", err)
		}
	}()

	// gRPC server
	grpcPort := config.GetString("grpc_port")
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		logger.Fatal("Failed to listen for gRPC: ", err)
	}

	grpcServer := grpc.NewServer()
	// Register gRPC services here
	exampleServer := proto.NewExampleServer(logger)
	proto.RegisterExampleServiceServer(grpcServer, exampleServer)

	reflection.Register(grpcServer) // Enable reflection for debugging

	logger.Info("Starting gRPC server on :", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
