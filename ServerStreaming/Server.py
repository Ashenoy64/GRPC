import grpc
from concurrent import futures
import time
import dotenv
import os

import NPime_pb2
import NPime_pb2_grpc

dotenv.load_dotenv()


def isPrime(n):
    if n == 1 or n == 0:
        return False
    for i in range(2, int(n**(1/2)) + 1):
        if n % i == 0:
            return False
    return True

class NPimeServicer(NPime_pb2_grpc.NPimeServicer):
    def GetNPime(self, request, context):
        number = request.number
        print("Received number:", number)
        for i in range(1, number + 1):
            if isPrime(i):
                yield NPime_pb2.NPimeResponse(prime=i)

def serve(ip, port):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    NPime_pb2_grpc.add_NPimeServicer_to_server(NPimeServicer(), server)
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
