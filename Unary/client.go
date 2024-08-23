package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "path/to/prime"

	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <n>", os.Args[0])
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid argument: %v", err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPrimeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.PrimeRequest{N: int32(n)}
	res, err := client.get_nth_prime(ctx, req)
	if err != nil {
		log.Fatalf("Error calling get_nth_prime: %v", err)
	}

	fmt.Printf("The %dth prime number is %d\n", n, res.GetPrime())
}
