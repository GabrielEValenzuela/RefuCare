from fastapi import FastAPI
import uvicorn

from app.routes.predict import router as predict_router
from app.core.lifespan import lifespan

app = FastAPI(lifespan=lifespan)

app.include_router(predict_router,prefix="/api", tags=["predict"])

@app.get("/health")
async def health_check():
    return {"status": "ok"}

if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8000)
