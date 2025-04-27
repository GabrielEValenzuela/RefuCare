from pydantic import BaseModel

class PredictionInput(BaseModel):
    age: int
    blood_pressure_sys: int
    blood_pressure_dias: int
    is_active: bool

class PredictionResponse(BaseModel):
    data: dict
