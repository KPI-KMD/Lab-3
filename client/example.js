const { time } = require('console');
const { TIMEOUT } = require('dns');
const telemetric = require('./telemetrics/client');

const client = telemetric.Client('http://localhost:8080');

// Scenario 1: get 50 last records
client.get50LastRecords(1)
    .then((listOfTelemetries) => {
        console.log('=== Scenario 1 ===');
        console.log(listOfTelemetries);
    })
        
    .catch(err => {
        console.log('=== Scenario 1 ===');
        console.log(`Error trying to get 50 last records: ${err}`)
    });

// Scenario 2: send data
client.sendData("class-1-tablet-1", '89', "currentVideo", new Date())
    .then(resp => {
        console.log('=== Scenario 2 ===');
        console.log('Added new telemetry status:', resp)
    })

    .catch(err => {
            console.log('=== Scenario 2 ===');
            console.log(`Error trying to set telemetry: ${err}`)
    });

// Scenario 3: Insertion before pass 10 seconds
client.sendData("class-1-tablet-1", '32', "vidosik", new Date())
    .then(resp => {
        console.log('=== Scenario 3 ===');
        console.log('Added new telemetry status:', resp)
    })

    .catch(err => {
            console.log('=== Scenario 3 ===');
            console.log(`Error trying to set telemetry: ${err}`)
    });

