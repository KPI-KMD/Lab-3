const telemetric = require('./telemetrics/client');

const client = telemetric.Client('http://localhost:8080');

client.get50LastRecords(1)
    .then((listOfTelemetries) => {
        console.log('=== Scenario 1 ===');
        console.log(listOfTelemetries);
    })
        
    .catch(err => {
        console.log('=== Scenario 1 ===');
        console.log(`Error trying to get 50 last records: ${err}`)
    });

client.sendTelemetricValues({})
    .then(resp => {
        
    })