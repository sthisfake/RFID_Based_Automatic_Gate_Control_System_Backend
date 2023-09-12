import os
import time
import serial.tools.list_ports


arduino_port = 'COM9'

python_script = "C:\\Users\\Pouya\\Desktop\\ardu\\scripts\\read_data.py"


while True:

    connected_ports = [
        port.device for port in serial.tools.list_ports.comports()]

    if arduino_port in connected_ports:
        print("connected")
        os.system(f'python {python_script}')
    else:
        print("not  connected")
    time.sleep(0.1)
