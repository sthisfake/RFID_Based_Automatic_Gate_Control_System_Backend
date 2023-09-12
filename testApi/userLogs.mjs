import PocketBase from "pocketbase";

const pb = new PocketBase("http://127.0.0.1:8090");

await pb.admins.authWithPassword("admingateapp@gmail.com", "adminadmin45");

const result = await pb.send("/user_logs/nfwmxf8az3argf3/2/100/", "GET");

console.log(result);
