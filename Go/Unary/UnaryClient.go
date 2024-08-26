package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"Unary/Unary"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ip := os.Getenv("SERVER_IP")
	port := os.Getenv("SERVER_PORT")
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

	client := Unary.NewPrimeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &Unary.PrimeRequest{N: int32(n)}
	res, err := client.GetNthPrime(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetNthPrime: %v", err)
	}

	fmt.Printf("The %dth prime number is %d\n", n, res.GetPrime())
}
