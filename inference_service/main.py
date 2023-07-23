import os
import sys

print(sys.path)

from flask import Flask

# from inference_service.inference import get_inference

app = Flask(__name__)

@app.route("/")
def hello_world():
    return {"response": "inference service"}

@app.route("/inference", method=["POST"])
def generate_inference():
    return "<p>Hello, World!</p>"

@app.route("/runtime", method=["POST"])
def set_runtime():
    return "<p>Hello, World!</p>"

if __name__ == '__main__':
    app.run(debug=True, use_reloader=True)
