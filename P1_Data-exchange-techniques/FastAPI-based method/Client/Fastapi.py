from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

app = FastAPI()

class AddRequest(BaseModel):
    a: int
    b: int

class AddResponse(BaseModel):
    result: int

@app.post("/add", response_model=AddResponse)
async def add_numbers(req: AddRequest):
    # 简单校验（其实Pydantic会保证类型）
    if req.a is None or req.b is None:
        raise HTTPException(status_code=400, detail="Missing parameter 'a' or 'b'")
    return AddResponse(result=req.a + req.b)
