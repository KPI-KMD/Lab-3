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

    client.get50LastRecords(2)
    .then((listOfTelemetries) => {
        console.log('class-1-tablet-2 before insertion:');
        console.log(listOfTelemetries);
    })

// Scenario 2: send data
client.sendData("class-1-tablet-2", '89', "currentVideo", new Date())
    .then(resp => {
        console.log('=== Scenario 2 ===');
        console.log('Added new telemetry status:', resp)
        return client.get50LastRecords(2)
        .then((listOfTelemetries) => {
            console.log('class-1-tablet-2 after insertion:');
            console.log(listOfTelemetries);
        })
    })

    .catch(err => {
            console.log('=== Scenario 2 ===');
            console.log(`Error trying to set telemetry: ${err}`)
    });
