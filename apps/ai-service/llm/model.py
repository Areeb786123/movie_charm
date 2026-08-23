import os 

from dotenv import load_dotenv
from langchain_huggingface import HuggingFaceEndpoint

load_dotenv()
from dotenv import load_dotenv
from langchain_huggingface import HuggingFaceEndpoint, ChatHuggingFace

load_dotenv()

llm_endpoint = HuggingFaceEndpoint(
    repo_id="meta-llama/Llama-3.1-8B-Instruct",
    provider="featherless-ai",   # cheap, confirmed-working provider for this model
    task="conversational",
    max_new_tokens=100,
    temperature=0.1,
)
llm = ChatHuggingFace(llm = llm_endpoint)
