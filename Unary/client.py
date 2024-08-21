import grpc
import prime_pb2
import prime_pb2_grpc
import os
from dotenv import load_dotenv

load_dotenv()

ip = os.getenv("IP")
port = os.getenv("PORT")
position=int(os.getenv("POSITION"))

def run():
    channel = grpc.insecure_channel(f'{ip}:{port}')
    stub = prime_pb2_grpc.PrimeNumberServiceStub(channel)
    
    request = prime_pb2.PrimeRequest(position=position)
    
    try:
        response = stub.get_prime_number(request)
        print(f"The {position}th prime number is {response.number}.")
    except grpc.RpcError as e:
        print(f"An error occurred: {e}")

if __name__ == '__main__':
    run()
