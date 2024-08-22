import grpc
from concurrent import futures
import time
import os
from dotenv import load_dotenv

import median_pb2
import median_pb2_grpc

load_dotenv()

ip = os.getenv("IP")  
port = os.getenv("PORT")  

arr = []

class FindMedianService(median_pb2_grpc.FindMedianServiceServicer):
    def get_median(self, request_iterator, context):
        global arr
        arr=[]
        for request in request_iterator:
            number = request.number
            arr.append(number)
            
            for i in range(len(arr)-1,0,-1):
                if arr[i]<arr[i - 1]:
                    arr[i],arr[i-1]=arr[i-1],arr[i]
                else:
                    break
        size = len(arr)
        if size % 2 == 0:
            median = (arr[size // 2] + arr[size // 2 - 1]) / 2.0
        else:
            median = arr[size // 2]
        return median_pb2.Response(median=median)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    median_pb2_grpc.add_FindMedianServiceServicer_to_server(FindMedianService(), server)
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
