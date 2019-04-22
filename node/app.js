const express = require('express');
const app = express();

const projectID = 'csp-testing';
const serviceName = "nodejs-profiler";

require("@google-cloud/profiler").start({
    serviceContext: {
        projectID: projectID,        
        service: serviceName,
        version: "eks"
    },
    logLeveL: 3,
});

function getRandomInt(max) {
    return Math.floor(Math.random() * Math.floor(max));
  }

function blockCpuFor(ms) {
	var now = new Date().getTime();
    var result = 0
    var shouldRun = true;
	while(shouldRun) {
		result += Math.random() * Math.random();
        if (new Date().getTime() > now +ms)
			return;
	}	
}

app.get('/', (req, res) => {
    const delay = getRandomInt(5000);
    console.log('request made');
    blockCpuFor(delay);
    res.send('Delayed for ' + delay);
    console.log('delayed for ' + delay);
})

app.listen(8081, () => console.log(`Example app listening on port 8081!`))