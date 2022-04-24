from services import run_producer, run_consumer, run_quotes
from flask import Flask, request
from flask import Response
import json
import os


app = Flask(__name__)

@app.route('/quotes', methods=['GET'])
def event():
    return Response(json.dumps(run_quotes()), status=200)


if __name__ == '__main__':
    # run_producer()
    # run_consumer()
    
    app.run(host='0.0.0.0', threaded=True)

