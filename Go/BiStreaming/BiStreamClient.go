package main

import (
	"BiStreaming/Inverse"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	serverIP := os.Getenv("SERVER_IP")
	serverPort := os.Getenv("SERVER_PORT")
	biStreamNumberStr := os.Getenv("BI_STREAM_NUMBER")
	biStreamNumber, err := strconv.Atoi(biStreamNumberStr)
	if err != nil {
		log.Fatalf("Error parsing BI_STREAM_NUMBER: %v", err)
	}

	conn, err := grpc.Dial(serverIP+":"+serverPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := Inverse.NewInverseMatrixClient(conn)

	matrix := generateRandomMatrix(biStreamNumber)

	stream, err := c.GetInverseMatrix(context.Background())
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	for _, row := range matrix {
		realRow := &Inverse.RealRows{Val: row}
		req := &Inverse.InverseMatrixRequest{}
		req.Rows = append(req.Rows, realRow)
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error sending matrix row: %v", err)
		}
	}
	stream.CloseSend()

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			// End of the stream, no more data expected
			log.Println("Stream ended")
			break
		}
		if err != nil {
			log.Fatalf("Error receiving response: %v", err)
		}
		fmt.Println("Received row:", resp.GetRows())
	}
}

func generateRandomMatrix(size int) [][]float64 {
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		row := make([]float64, size)
		for j := 0; j < size; j++ {
			row[j] = rand.Float64() * 100 // Random values
		}
		matrix[i] = row
	}
	return matrix
}
