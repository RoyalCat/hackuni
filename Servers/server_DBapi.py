import os
import time
from concurrent import futures
import logging
import random
import string
import threading
import grpc
import test_pb2
import test_pb2_grpc


class Listener(test_pb2_grpc.TestServiceServicer):
    """The listener function implemests the rpc call as described in the .proto file"""

    def __init__(self):
        self.counter = 0
        self.last_print_time = time.time()

    def __str__(self):
        return self.__class__.__name__

    def ping(self, request, context):
        print(request.mes)
        return test_pb2.O()


def serve():

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=1))
    test_pb2_grpc.add_TestServiceServicer_to_server(Listener(), server)
    server.add_insecure_port("[::]:8888")
    server.start()
    try:
        while True:
            pass
    except KeyboardInterrupt:
        print("KeyboardInterrupt")
        server.stop(0)


if __name__ == "__main__":
    serve()
