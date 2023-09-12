# from datetime import datetime
# import pytz

# timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
# print("here ok")
# local_time = datetime(
#     timestamp.year, timestamp.month, timestamp.day, 0, 0, 0)
# import datetime

# # Create a datetime object
# my_datetime = datetime.datetime(2023, 9, 12, 15, 30, 0)

# # Extract year, month, and day
# year = my_datetime.year
# month = my_datetime.month
# day = my_datetime.day

# # Print the results
# print("Year:", year)
# print("Month:", month)
# print("Day:", day)
from datetime import datetime

# Get the current timestamp as a formatted string
timestamp_str = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
timestamp = datetime.strptime(timestamp_str, "%Y-%m-%d %H:%M:%S")

# Now you can access the year, month, and day attributes
year = timestamp.year
month = timestamp.month
day = timestamp.day

print("Year:", year)
print("Month:", month)
print("Day:", day)
