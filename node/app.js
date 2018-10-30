const express = require('express');
const app = express();

const projectID = 'thegrinch-project';
const serviceName = 'profiler-node';

require("@google-cloud/profiler").start({
    serviceContext: {
        projectID: projectID,        
        service: serviceName,
        version: "0.0.1"
    },
    logLeveL: 3,
});

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

function getRandomInt(max) {
    return Math.floor(Math.random() * Math.floor(max));
  }

app.get('/', (req, res) => {
    console.log('request made');
    const delay = getRandomInt(5000);
    blockCpuFor(delay);
    res.send('Blocked CPU for ' + delay);
})

app.listen(8080, () => console.log(`Example app listening on port 8080!`))