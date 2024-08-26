package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"Median/Median"

	"math/rand"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ip := os.Getenv("SERVER_IP")
	port := os.Getenv("SERVER_PORT")
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

	client := Median.NewFindMedianServiceClient(conn)

	stream, err := client.GetMedian(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	rand.Seed(time.Now().UnixNano())
	var numbers []int32

	for i := 0; i < size; i++ {
		number := int32(randomInt(1, 1000))
		numbers = append(numbers, number)
		if err := stream.Send(&Median.Request{Number: number}); err != nil {
			log.Fatalf("Error sending number: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d ", numbers[i])
	}
	fmt.Printf("\n")
	fmt.Printf("The median is %f\n", resp.GetMedian())
}
