import os
import time
from concurrent import futures
import logging
import random
import string
import threading
import grpc
import AnalystService_pb2
import AnalystService_pb2_grpc
import test_pb2
import test_pb2_grpc


# import Next Service pb2/pb2_grpc


class Listener(AnalystService_pb2_grpc.AnalystServiceServicer):
    """The listener function implemests the rpc call as described in the .proto file"""

    def __init__(self):
        self.counter = 0
        self.last_print_time = time.time()

    def __str__(self):
        return self.__class__.__name__

    def Analyse(self, request, context):
        print(request.message)

        process = Process(request.message)

        test_pb2_grpc.TestServiceStub(grpc.insecure_channel("localhost:8888")).ping(test_pb2.Ent(mes=process))
        print('sended!!!')

        return AnalystService_pb2.Out()

def Process(inpt):
    res = "f".join(inpt)
    return res


def serve():

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=1))
    AnalystService_pb2_grpc.add_AnalystServiceServicer_to_server(Listener(), server)
    server.add_insecure_port("[::]:9999")
    server.start()
    try:
        while True:
            pass
    except KeyboardInterrupt:
        print("KeyboardInterrupt")
        server.stop(0)


if __name__ == "__main__":
    serve()
