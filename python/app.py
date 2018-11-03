from flask import Flask
import logging
import random
import string
import time
import datetime
app = Flask(__name__)

def blockCPU(delay):
    logging.warn("delaying for " + str(delay))
    startTime = datetime.datetime.now()
    endTime = startTime + datetime.timedelta(seconds = delay)
    result = 0
    logging.warn("start time is " + str(startTime))
    logging.warn("end time is " + str(endTime)) 
    while True: 
        result += random.random() * random.random()
        if (datetime.datetime.now() > endTime):
            logging.warn("exiting loop")
            return
 
 

@app.route('/')
def doProfile():
    # generate random number to delay for:
    timeToBlock = random.randint(0,10)
    blockCPU(timeToBlock)
    return ("CPU blocked for " + str(timeToBlock))

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0')