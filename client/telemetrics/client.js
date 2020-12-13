const http = require("../common/http.js");

const Client = baseURL => {
    const client = http.Client(baseURL);

    return {
        get50LastRecords:(tabletID) => client.get(tabletID, "/telemetrics"),
        sendData:(name, battery, currentvideo, devicetime) => 
        client.post( "/telemetrics", {name, battery, currentvideo, devicetime})
    };
}

module.exports = {Client};
