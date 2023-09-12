import PocketBase from "pocketbase";

const pb = new PocketBase("http://127.0.0.1:8090");

await pb.admins.authWithPassword("admingateapp@gmail.com", "adminadmin45");

const result = await pb.send("/dashbord/1402-06-19/1/2/", "GET");

console.log(result);
