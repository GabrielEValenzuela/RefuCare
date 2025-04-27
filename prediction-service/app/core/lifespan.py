# app/core/lifespan.py
from fastapi import FastAPI
import asyncio

from contextlib import asynccontextmanager

from app.core.model_manager import ModelManager
from app.core.eureka import eureka

@asynccontextmanager
async def lifespan(app: FastAPI):
    app.state.model_manager = ModelManager(
        model_path="app/core/data/model.pkl",
    )
    print("[DEBUG] Model manager initialized.")

    await eureka.register()  # Register on startup
    asyncio.create_task(eureka.send_heartbeat())  # Start heartbeat
    print("[DEBUG] Eureka registration and heartbeat started.")
    yield
