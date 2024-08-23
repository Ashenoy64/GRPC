package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "Median/Median"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"math/rand"
)

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	sizeStr := os.Getenv("SIZE")

	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		log.Fatalf("Invalid SIZE value: %v", err)
	}

	address := ip + ":" + port
	fmt.Printf("Connecting to gRPC server at %s\n", address)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFindMedianServiceClient(conn)

	stream, err := client.GetMedian(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		number := int32(randomInt(1, 1000))
		// fmt.Printf("Sending number: %d\n", number)
		if err := stream.Send(&pb.Request{Number: number}); err != nil {
			log.Fatalf("Error sending number: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Printf("The median is %f\n", resp.GetMedian())
}
