import grpc
import AnalystService_pb2
import AnalystService_pb2_grpc


def run():
    with grpc.insecure_channel("localhost:9999") as channel:
        stub = AnalystService_pb2_grpc.AnalystServiceStub(channel)
        while True:
            try:
                mess = input()
                stub.Analyse(AnalystService_pb2.Enter(message=str(mess)))
            except KeyboardInterrupt:
                print("KeyboardInterrupt")
                channel.unsubscribe(close)
                exit()



def close(channel):
    channel.close()


if __name__ == "__main__":
    run()
