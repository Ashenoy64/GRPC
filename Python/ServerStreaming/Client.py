import grpc
import NPime_pb2
import NPime_pb2_grpc

import os
import dotenv   

dotenv.load_dotenv()

server_ip = os.getenv("SERVER_IP")
server_port = os.getenv("SERVER_PORT")
primes = int(os.getenv("SERVER_STREAM_NUMBER"))

def run():
    with grpc.insecure_channel(f'{server_ip}:{server_port}') as channel:
        stub = NPime_pb2_grpc.NPimeStub(channel)
        response = stub.GetNPime(NPime_pb2.NPimeRequest(number=primes))
        for prime in response:
            print("Prime number:", prime.prime)

if __name__ == '__main__':
    run()
