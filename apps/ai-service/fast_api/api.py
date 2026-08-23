from fastapi import FastAPI
from pydantic import BaseModel

from llm.graph.graph import graph

app = FastAPI()

class CommentRequest(BaseModel):
    comment:str
    
   
@app.get("/health")
def health():
    return{
        "status":"ok"
    }        
    
@app.post("/analyze")
def analyze_comment(request: CommentRequest):

    result = graph.invoke({
        "comment": request.comment
    })

    return {
        "message": result["message"],
        "sentiment_type": result["sentiment"]
    }   