const http = require("../common/http.js");

const Client = baseURL => {
    const client = http.Client(baseURL);

    return {
        get50LastRecords:() => client.get("/telemetrics"),
        sendTelemeticValues: values => client.post("/telemetrics", values)
    };
}

module.exports = {Client};