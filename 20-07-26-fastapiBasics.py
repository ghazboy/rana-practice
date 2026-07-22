from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def read_root():
    return {"message": "Hello, Rana!"}

@app.get("/about")
def about_page():
    return {"message": "This is my practicing page for Rana"}