from opcua import Server
from random import randint
from random import random
import time

server = Server()

url = "opc.tcp://127.0.0.99:48400"
server.set_endpoint(url)

name = "OPCUA_SIMULATION_SERVER"
addspace = server.register_namespace(name)

node = server.get_objects_node()

Param = node.add_object(addspace, "Parameters")

sim_vars = {
    "Temperature1": range(0, 14200000),
    "Temperature2": range(0, 14200000),
    "Pressure": range(0, 1000),
    "Mass": range(0, 100000),
    "FluidFlow": range(0, 1000),
    "pH": range(0, 1, -1),
    "co2": range(0, 1, -1)
}

server_vars = [(key, Param.add_variable(addspace, key, 0)) for key in sim_vars.keys()]

for var in server_vars:
    var[1].set_writable()

server.start()
print(f"Server started at {url}")

while True:
    for var in server_vars:
        var_range = sim_vars[var[0]]
        randN = randint(var_range.start, var_range.stop)
        if var_range.step == -1:
            randN = random()
        var[1].set_value(randN)
        print(str(var[1]) + " set to " + str(var[1].get_data_value().Value.Value))

    try:
        time.sleep(10)
    except:
        break
   

server.stop()