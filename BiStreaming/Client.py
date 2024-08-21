import grpc
import EigenVector_pb2
import EigenVector_pb2_grpc
import os
import numpy as np
from dotenv import load_dotenv

load_dotenv()
server_ip = os.getenv("SERVER_IP")
server_port = os.getenv("SERVER_PORT")
matrix_size = int(os.getenv("MATRIX_SIZE", 3))

def generate_random_matrix(m, n):
    return np.random.rand(m, n)

def generate_requests():
    matrix = generate_random_matrix(matrix_size, matrix_size)
    rows = [EigenVector_pb2.RealRows(row=row.tolist()) for row in matrix]
    yield EigenVector_pb2.EigenVectorRequest(rows=rows)

def run():
    channel = grpc.insecure_channel(f'{server_ip}:{server_port}')
    stub = EigenVector_pb2_grpc.EigenVectorStub(channel)
    try:
        responses = stub.GetEigenVector(generate_requests())
        for response in responses:
            print("Received eigenvectors:")
            for row in response.rows:
                row_values = [f"{num.real} + {num.imaginary}i" for num in row.row]
                print(row_values)
    except grpc.RpcError as e:
        print(f"RPC error: {e.code()} - {e.details()}")

if __name__ == '__main__':
    run()
