const express = require('express');
const app = express();

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
    console.log('request made');
    res.send('Hello World!');
})

app.listen(8080, () => console.log(`Example app listening on port 8080!`))