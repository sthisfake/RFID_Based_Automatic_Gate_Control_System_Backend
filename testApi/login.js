import PocketBase from "pocketbase";

const pb = new PocketBase("http://45.149.77.147:8090");

await pb.admins.authWithPassword("admingateapp@gmail.com", "adminadmin45");
