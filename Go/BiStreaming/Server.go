package main

import (
	"BiStreaming/Inverse"
	"io"
	"log"

	"net"
	"os"

	"github.com/joho/godotenv"
	"gonum.org/v1/gonum/mat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	Inverse.UnimplementedInverseMatrixServer
}

func (s *server) GetInverseMatrix(stream Inverse.InverseMatrix_GetInverseMatrixServer) error {
	var matrix [][]float64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err == grpc.ErrServerStopped {
			break
		}
		if err != nil {
			return err
		}
		var row []float64
		for _, val := range req.GetRows() {
			row = append(row, val.Val...)
		}
		matrix = append(matrix, row)
	}
	invMatrix, err := invertMatrix(matrix)
	if err != nil {
		return err
	}

	for _, row := range invMatrix {
		realRow := &Inverse.RealRows{Val: row}
		res := &Inverse.InverseMatrixResponse{}
		res.Rows = append(res.Rows, realRow)
		if err := stream.Send(res); err != nil {
			return err
		}
	}
	return nil
}

func invertMatrix(matrix [][]float64) ([][]float64, error) {
	matA := mat.NewDense(len(matrix), len(matrix[0]), nil)
	for i := range matrix {
		matA.SetRow(i, matrix[i])
	}

	// Compute the inverse
	var matInv mat.Dense
	if err := matInv.Inverse(matA); err != nil {
		return nil, err
	}

	// Convert the result back to a 2D slice
	rows, cols := matInv.Dims()
	invMatrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		invMatrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			invMatrix[i][j] = matInv.At(i, j)
		}
	}

	return invMatrix, nil
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "50051" // Default port
	}

	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	Inverse.RegisterInverseMatrixServer(s, &server{})
	reflection.Register(s)

	log.Printf("Server listening at %v", serverPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
