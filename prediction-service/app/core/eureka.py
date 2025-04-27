# app/core/eureka.py

import asyncio
import socket
import httpx
import os

EUREKA_SERVER = os.getenv("EUREKA_SERVER", "http://localhost:8761/eureka")
APP_NAME = os.getenv("EUREKA_APP_NAME", "PREDICTION-SERVICE")
INSTANCE_PORT = int(os.getenv("INSTANCE_PORT", 8000))

instance_id = f"{socket.gethostname()}:{APP_NAME}:{INSTANCE_PORT}"
eureka_url = f"{EUREKA_SERVER}/apps/{APP_NAME}"

async def register():
    payload = {
        "instance": {
            "instanceId": instance_id,
            "hostName": socket.gethostname(),
            "app": APP_NAME,
            "ipAddr": socket.gethostbyname(socket.gethostname()),
            "status": "UP",
            "port": {"$": INSTANCE_PORT, "@enabled": "true"},
            "dataCenterInfo": {
                "@class": "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
                "name": "MyOwn"
            }
        }
    }
    async with httpx.AsyncClient() as client:
        response = await client.post(eureka_url, json=payload, headers={"Content-Type": "application/json"})
        if response.status_code in (204, 200):
            print(f"[EUREKA] Successfully registered {APP_NAME}")
        else:
            print(f"[EUREKA] Registration failed: {response.status_code} {response.text}")

async def send_heartbeat():
    while True:
        async with httpx.AsyncClient() as client:
            response = await client.put(f"{eureka_url}/{instance_id}")
            if response.status_code in (204, 200):
                print("[EUREKA] Heartbeat sent.")
            else:
                print(f"[EUREKA] Heartbeat failed: {response.status_code} {response.text}")
        await asyncio.sleep(30)  # send heartbeat every 30 seconds
