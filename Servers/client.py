import grpc
import AnalystService_pb2
import AnalystService_pb2_grpc
import time


def Analyse():
    data = [3, 1, 5, 4, 5, 6, 7, 2323]
    print('message-', data)

    critical_v = [2, 2, 2, 2, 2, 2, 2]

    print('criticals- ', critical_v)

    result = []
    flag = False

    counter = 0
    for i in range(len(data[:-1])):
        if data[i] < critical_v[i]:
            result.append(1)
            flag = True
        else:
            result.append(0)
    result.append(data[-1])
    print(result)

Analyse()
