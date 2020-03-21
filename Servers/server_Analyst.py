import time
from concurrent import futures
from clickhouse_driver import Client
import grpc
import AnalystService_pb2
import AnalystService_pb2_grpc
import telebot
from telebot import apihelper

bot = telebot.TeleBot('1101313253:AAG-3hK95_Ojk6G1yfj_7ogK5NyPjI2AXUk')
chatids = []

proxies = {
'http': 'http://52.179.231.206:80',
'https': 'https://167.172.191.252:8118'
}

apihelper.proxy = proxies

class Listener(AnalystService_pb2_grpc.AnalystServiceServicer):
    """The listener function implemests the rpc call as described in the .proto file"""

    def __init__(self):
        self.counter = 0
        self.last_print_time = time.time()
        self.client = Client(host='78.140.223.19', password='12345')

    def __str__(self):
        return self.__class__.__name__

    def Analyse(self, request, context):
        data = request.message
        print('message-', request.message)

        critical_v = self.client.execute('SELECT * FROM test.critical')[0]

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

        if flag:
            self.client.execute(f'INSERT INTO test.warnings (Pressure, Humidity, TemperatureR, TemperatureA, pH, FlowRate, CO, EventTime) VALUES ({",".join(map(str, result))})')
            names = ['Pressure', 'Humidity', 'TemperatureR', 'TemperatureA', 'pH', 'FlowRate', 'CO', 'EventTime']
            for i in chatids:
                bot.send_message(i, 'Warning on sensors: ' + ",".join(map(str, [names[j] for j in range(len(names)) if result[i] == 1])))

        return AnalystService_pb2.Out()


def serve():

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=1))
    AnalystService_pb2_grpc.add_AnalystServiceServicer_to_server(Listener(), server)
    server.add_insecure_port("[::]:9999")
    server.start()
    bot.polling()
    try:
        while True:
            pass
    except KeyboardInterrupt:
        print("KeyboardInterrupt  sss")
        server.stop(0)


@bot.message_handler(commands=['start'])
def start_message(message):
    bot.send_message(message.chat.id, 'Привет, пиши /enable что бы подписаться на рассылку предупреждений,'
                                      'и /disable для того чтобы отписатся')


@bot.message_handler(commands=['enable'])
def enable(message):
    chatids.append(message.chat.id)
    bot.send_message(message.chat.id, 'Подписка оформлена')


@bot.message_handler(commands=['disable'])
def disable(message):
    chatids.remove(message.chat.id)
    bot.send_message(message.chat.id, 'Подписка отменена')


if __name__ == "__main__":
    serve()
