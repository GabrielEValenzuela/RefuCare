import joblib
import pandas as pd
from app.model.schema import PredictionInput

FEATURES = ["age", "blood_pressure_sys", "blood_pressure_dias", "is_active"]

class ModelManager:
    def __init__(self, model_path: str):
        self.model = joblib.load(model_path)

    def predict(self, input_data: PredictionInput) -> dict:
        df = pd.DataFrame([input_data.model_dump()])

        prediction = self.model.predict(df[FEATURES])[0]
        confidence = round(max(self.model.predict_proba(df[FEATURES])[0]) * 100, 2)

        return {
            "prediction": int(prediction),
            "confidence": confidence
        }
