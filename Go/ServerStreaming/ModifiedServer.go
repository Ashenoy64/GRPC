package main

import (
	"ServerStreaming/NPrime"
	"log"
	"net"
	"os"
	"time"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	NPrime.UnimplementedNPrimeServer
}

func (s *server) GetNPrime(req *NPrime.NPrimeRequest, stream NPrime.NPrime_GetNPrimeServer) error {
	n := req.GetNumber()
	for i := int32(2); i <= n; i++ {
		if isPrime(i) {
			res := &NPrime.NPrimeResponse{Prime: i}
			if err := stream.Send(res); err != nil {
				return err
			}
		}
	}
	return nil
}

func isPrime(n int32) bool {
	if n < 2 {
		return false
	}
	for i := int32(2); i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	time.Sleep(time.Second*60)
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Read environment variables
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		log.Fatal("SERVER_PORT is not set")
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	NPrime.RegisterNPrimeServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
