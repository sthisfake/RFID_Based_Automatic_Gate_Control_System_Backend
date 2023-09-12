import time
import requests
import subprocess
import socket
from pocketbase import PocketBase  # Client also works the same
from pocketbase.client import FileUpload
from datetime import datetime
import pytz


allowed_ids = {"94d2faf4", "a47f0304", "90056026"}


client = PocketBase('http://192.168.232.224:8090')
file_path = "C:\\Users\\Pouya\\Desktop\\ardu\\scripts\\rfid_data.txt"


def check_internet_connection():
    try:
        socket.create_connection(("www.google.com", 80), timeout=1)
        return True
    except:
        return False


def transfer_data_to_server(data):
    if data[0]["Tag_ID"] in allowed_ids:
        tag_id = data[0]["Tag_ID"]
        time = data[0]["Timestamp"]
        try:
            client.admins.auth_with_password(
                "gate@gmail.com", "gategate47")

            result = client.collection('authorized_person').get_first_list_item(
                f'tag_id= "{tag_id}"')

            # check if the guy is in the building or its new entery
            timestamp_str = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
            timestamp = datetime.strptime(timestamp_str, "%Y-%m-%d %H:%M:%S")

            local_time = datetime(
                timestamp.year, timestamp.month, timestamp.day, 0, 0, 0)

            local_timezone = pytz.timezone('Iran')
            utc_time = local_timezone.localize(local_time).astimezone(
                pytz.utc).strftime("%Y-%m-%d %H:%M:%S")
            check = ""
            try:
                check = client.collection('monitoring').get_list(
                    1, 1,
                    {
                        "filter": f'user_id = "{result.id}" && entery_time > "{utc_time}"',
                        "sort": '-created'
                    }

                )
                check = check.items[0]
            except:
                data = {
                    "entery_time": f"{time}",
                    "user_id": result.id,
                }
                result2 = client.collection('monitoring').create(data)
                return True

            if check:

                if check.__dict__['exit_time'] == "":
                    data = {
                        "exit_time": f"{time}",
                    }
                    result2 = client.collection(
                        'monitoring').update(check.id, data)
                    return True
                else:
                    data = {
                        "entery_time": f"{time}",
                        "user_id": result.id,
                    }
                    result2 = client.collection('monitoring').create(data)
                    return True
            else:
                data = {
                    "entery_time": f"{time}",
                    "user_id": result.id,
                }
                result2 = client.collection('monitoring').create(data)
                return True

        except:
            print("Error creating")
            return False
    else:
        return True


def read_first_3_lines_data(file_path):
    data = []
    try:
        with open(file_path, "r") as file:
            for _ in range(3):  # Read the first 3 lines
                line = file.readline().strip()
                if line.startswith("Tag ID: "):
                    tag_id = line[len("Tag ID: "):]
                if line.startswith("Timestamp: "):
                    timestamp = line[len("Timestamp: "):]
                if not line:
                    None
            data.append({"Tag_ID": tag_id, "Timestamp": timestamp})

    except:
        None

    return data


def delete_first_3_lines(file_path):
    try:
        with open(file_path, "r") as file:
            lines = file.readlines()

        lines = lines[3:]

        with open(file_path, "w") as file:
            for line in lines:
                file.write(line)

    except:
        pass


while True:
    if check_internet_connection():

        print("connection established")
        data = read_first_3_lines_data(file_path)
        if data:
            result = transfer_data_to_server(data)
            if result:
                delete_first_3_lines(file_path)

    else:
        print("no internet")

    time.sleep(0.1)
