package main

import (
	// "context"
	"io"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "Median/Median"
)

type server struct {
	pb.UnimplementedFindMedianServiceServer
}

func (s *server) GetMedian(stream pb.FindMedianService_GetMedianServer) error {
	var numbers []int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// End of stream
			break
		}
		if err != nil {
			return err
		}
		number := req.GetNumber()
		numbers = append(numbers, number)

		// Perform insertion sort
		for i := len(numbers) - 1; i > 0; i-- {
			if numbers[i] < numbers[i-1] {
				numbers[i], numbers[i-1] = numbers[i-1], numbers[i]
			} else {
				break
			}
		}
	}

	// Calculate the median
	median := calculateMedian(numbers)

	// Send the response
	return stream.SendAndClose(&pb.Response{Median: median})
}

func calculateMedian(numbers []int32) float64 {
	n := len(numbers)
	if n == 0 {
		return 0
	}
	if n%2 == 1 {
		return float64(numbers[n/2])
	}
	return float64(numbers[n/2-1]+numbers[n/2]) / 2.0
}

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ip := os.Getenv("IP")
	port := os.Getenv("PORT")
	address := ip + ":" + port

	// Start listening on the specified IP and PORT
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFindMedianServiceServer(grpcServer, &server{})

	log.Printf("gRPC server is running on %s\n", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
