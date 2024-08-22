import grpc
import median_pb2
import median_pb2_grpc
import os
from dotenv import load_dotenv
import random

load_dotenv()

ip = os.getenv("IP")
port = os.getenv("PORT")

arr=[]

channel = grpc.insecure_channel(f'{ip}:{port}')
stub = median_pb2_grpc.FindMedianServiceStub(channel)

def generate_random():
    return random.randint(1, 1000)

def generate_requests():
    for _ in range(200):
        number = generate_random()
        arr.append(number)
        yield median_pb2.Request(number=number)

def run():
    response= stub.get_median(generate_requests())
    print(f"The median of the array is {response.median}")

if __name__ == '__main__':
    run()
