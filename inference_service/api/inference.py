from main import app

@app.route("/inference")
def get_inference():
    return {"response": "Inference Endpoint"}
