import PocketBase from "pocketbase";

const pb = new PocketBase("http://127.0.0.1:8090");

await pb.admins.authWithPassword("admingateapp@gmail.com", "adminadmin45");

const result = await pb.send("/overall_states/today", "GET");

console.log(result);
