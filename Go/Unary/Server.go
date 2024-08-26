package main

import (
	"Unary/Unary"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	Unary.UnimplementedPrimeServiceServer
}

func (s *server) GetNthPrime(ctx context.Context, req *Unary.PrimeRequest) (*Unary.PrimeResponse, error) {
	n := req.GetN()
	prime := nthPrime(int(n))
	return &Unary.PrimeResponse{Prime: int64(prime)}, nil
}

func nthPrime(n int) int {
	count := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			count++
			if count == n {
				return i
			}
		}
	}
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ip := os.Getenv("SERVER_IP")
	port := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf("%s:%s", ip, port)

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	Unary.RegisterPrimeServiceServer(grpcServer, &server{})

	fmt.Printf("gRPC server is running on %s\n", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}