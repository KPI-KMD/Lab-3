const http = require("../common/http.js");

const Client = baseURL => {
    const client = http.Client(baseURL);

    return {
        get50LastRecords:(tabletID) => client.get(tabletID, "/telemetrics"),
        sendTelemetricValues: values => client.post("/telemetrics", values)
    };
}

module.exports = {Client};