# app/routes/predict.py
from fastapi import APIRouter, HTTPException, Request, Depends
from app.core.model_manager import ModelManager
from app.model import schema

router = APIRouter()

def get_model(request: Request) -> ModelManager:
    model_manager = request.app.state.model_manager
    if model_manager is None:
        raise RuntimeError("ML model is not initialized!")
    return model_manager

@router.post("/predict")
async def predict(request: Request, model: ModelManager = Depends(get_model)):
    try:
        input_data = await request.json()
        input = schema.PredictionInput(**input_data)

        result = model.predict(input)
        msg = ""

        if result["confidence"] < 0.5:
            msg = "The model is not confident enough to make a prediction."
        elif result["confidence"] > 0.8:
            msg = "The model is very confident in its prediction."
        else:
            msg = "The model is moderately confident in its prediction. Needs more validation."

        return {
            "data": {
            "prediction": "No hypertension" if result["prediction"] == 0 else "Hypertension",
            "confidence": f"{result['confidence']}%",
            "notes": msg,
            "schema": {
                "version": "1.0",
            }
            }
        }
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
