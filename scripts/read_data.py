import serial

ser = serial.Serial('COM9', 9600)

file_path = "C:\\Users\\Pouya\\Desktop\\ardu\\scripts\\rfid_data.txt"

while True:
    print("here")
    try:
        line = ser.readline().decode('utf-8').strip()

        if line.startswith("Tag ID: "):
            tag_id = line[len("Tag ID: "):]

            from datetime import datetime
            timestamp = datetime.utcnow().strftime("%Y-%m-%d %H:%M:%S")

            with open(file_path, "a") as file:
                file.write(f"Tag ID: {tag_id}\n")
                file.write(f"Timestamp: {timestamp}\n\n")

    except KeyboardInterrupt:
        break

ser.close()
