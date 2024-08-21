import grpc
from concurrent import futures
import time
import os
from dotenv import load_dotenv

import prime_pb2
import prime_pb2_grpc

# Load environment variables from .env file
load_dotenv()

ip = os.getenv("IP")  # Default to localhost if not set
port = os.getenv("PORT")  # Default port if not set

# Function to check if a number is prime
def is_prime(n):
    if n <= 1:
        return False
    for i in range(2, int(n**0.5) + 1):
        if n % i == 0:
            return False
    return True

# Function to find the nth prime number
def nth_prime(n):
    count = 0
    num = 2
    while True:
        if is_prime(num):
            count += 1
            if count == n:
                return num
        num += 1

# Implement the gRPC service
class PrimeNumberService(prime_pb2_grpc.PrimeNumberServiceServicer):
    def get_prime_number(self, request, context):
        n = request.position
        prime_number = nth_prime(n)
        return prime_pb2.PrimeResponse(number=prime_number)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    prime_pb2_grpc.add_PrimeNumberServiceServicer_to_server(PrimeNumberService(), server)
    server.add_insecure_port(f'{ip}:{port}')
    server.start()
    print(f"Server started on {ip}:{port}")
    try:
        while True:
            time.sleep(86400) 
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()
