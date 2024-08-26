package main

import (
	"ServerStreaming/NPrime"
	"context"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read environment variables
	serverIP := os.Getenv("SERVER_IP")
	if serverIP == "" {
		log.Fatal("SERVER_IP is not set")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT is not set")
	}

	streamNumberStr := os.Getenv("SERVER_STREAM_NUMBER")
	if streamNumberStr == "" {
		log.Fatal("SERVER_STREAM_NUMBER is not set")
	}

	streamNumber, err := strconv.Atoi(streamNumberStr)
	if err != nil {
		log.Fatalf("Invalid SERVER_STREAM_NUMBER: %v", err)
	}

	conn, err := grpc.Dial(serverIP+":"+port, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NPrime.NewNPrimeClient(conn)

	req := &NPrime.NPrimeRequest{Number: int32(streamNumber)}

	stream, err := client.GetNPrime(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get primes: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			// End of the stream, no more data expected
			log.Println("Stream ended")
			break
		}
		if err != nil {
			log.Fatalf("error receiving from stream: %v", err)
		}
		log.Printf("Prime: %d", res.GetPrime())
	}
}
