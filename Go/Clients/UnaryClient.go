package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "Unary/Unary"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	nStr := os.Getenv("N")

	n, err := strconv.Atoi(nStr)
	if err != nil {
		log.Fatalf("Invalid value for N: %v", err)
	}

	address := fmt.Sprintf("%s:%s", ip, port)
	fmt.Printf("Connecting to gRPC server at %s\n", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPrimeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.PrimeRequest{N: int32(n)}
	res, err := client.GetNthPrime(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetNthPrime: %v", err)
	}

	fmt.Printf("The %dth prime number is %d\n", n, res.GetPrime())
}
