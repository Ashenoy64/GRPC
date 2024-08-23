from concurrent import futures
import grpc
import EigenVector_pb2
import EigenVector_pb2_grpc
import os
import dotenv
import numpy as np
import time

dotenv.load_dotenv()

class EigenVectorServicer(EigenVector_pb2_grpc.EigenVectorServicer):
    def GetEigenVector(self, request_iterator, context):
        matrices = []
        try:
            # Collect all requests and convert to numpy arrays
            for request in request_iterator:
                matrix = np.array([row.row for row in request.rows], dtype=np.float64)
                matrices.append(matrix)
            
            for matrix in matrices:
                # Ensure matrix is square
                if matrix.shape[0] != matrix.shape[1]:
                    context.set_details("Non-square matrix provided.")
                    context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
                    return
                
                eigenvalues, eigenvectors = np.linalg.eig(matrix)
                
                # Convert eigenvectors to protobuf format
                rows = []
                for eigenvector in eigenvectors.T:
                    complex_numbers = [
                        EigenVector_pb2.ComplexNumber(real=num.real, imaginary=num.imag) for num in eigenvector
                    ]
                    rows.append(EigenVector_pb2.ComplexRows(row=complex_numbers))
                
                response = EigenVector_pb2.EigenVectorResponse(rows=rows)
                yield response

        except Exception as e:
            context.set_details(f'Error occurred: {str(e)}')
            context.set_code(grpc.StatusCode.INTERNAL)

def serve(ip, port):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    EigenVector_pb2_grpc.add_EigenVectorServicer_to_server(EigenVectorServicer(), server)
    server_address = f'{ip}:{port}'
    server.add_insecure_port(server_address)
    server.start()
    print(f"Server started on {server_address}")
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    server_ip = os.getenv("SERVER_IP")
    server_port = os.getenv("SERVER_PORT")
    serve(server_ip, server_port)
