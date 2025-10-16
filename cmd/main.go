package main

import (
	"log"
	"net"
	"os"

	"github.com/P04KA/statistics/internal/app"
	"github.com/P04KA/statistics/pkg/stats"
	"google.golang.org/grpc"
)

func main() {
	application := app.New()

	port := os.Getenv("5052")

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("failed to listen:", err)
	}
	defer application.Close()

	grpcServer := grpc.NewServer()

	statsHandler := application.GetStatsHandler()
	stats.RegisterUserStatsServiceServer(grpcServer, statsHandler)

	log.Printf("statistics start on")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("fail:", err)
	}
}
