package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "path/to/prime"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPrimeServiceServer
}

func (s *server) get_nth_prime(ctx context.Context, req *pb.PrimeRequest) (*pb.PrimeResponse, error) {
	n := req.GetN()
	prime := nth_prime(int(n))
	return &pb.PrimeResponse{Prime: int64(prime)}, nil
}

func nth_prime(n int) int {
	count := 0
	for i := 2; ; i++ {
		if is_prime(i) {
			count++
			if count == n {
				return i
			}
		}
	}
}

// is_prime checks if a number is prime
func is_prime(num int) bool {
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
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpc_server := grpc.NewServer()
	pb.RegisterPrimeServiceServer(grpc_server, &server{})

	fmt.Println("gRPC server is running on port 50051")
	if err := grpc_server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
