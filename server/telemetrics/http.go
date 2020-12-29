package telemetrics

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/KPI-KMD/Lab-3/server/tools"
)

// Channels HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of channels HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListTelemetries(r, rw, store)
		} else if r.Method == "POST" {
			handleSendData(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleListTelemetries(r *http.Request, rw http.ResponseWriter, store *Store) {
	var idTabl int
	if err := json.NewDecoder(r.Body).Decode(&idTabl); err != nil {
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	res, err := store.ListTelemetries(idTabl)
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}

func handleSendData(r *http.Request, rw http.ResponseWriter, store *Store) {
	var tt SendData

	if err := json.NewDecoder(r.Body).Decode(&tt); err != nil {
		log.Printf("Error decoding channel input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}

	currentTabletID, err := store.FindID(tt.Name)
	if err != nil {
		log.Printf("Error finding current ID record: %s", err)
		tools.WriteJsonInternalError(rw)
	}

	previousTelemetryList, err := store.ListTelemetries(currentTabletID)
	if err != nil {
		log.Printf("Error finding previous telemetry list: %s", err)
		tools.WriteJsonBadRequest(rw, "Invalid tablet id")
		return
	}

	previousTelemetryID, err := store.maxTelemetryID(currentTabletID)
	if err != nil {
		log.Printf("Error finding previous ID record: %s", err)
		tools.WriteJsonInternalError(rw)
	}

	lastTime := time.Unix(0, 0)
	if len(previousTelemetryList.Telemetry) > 0 {
		if parsed, err := time.Parse(time.RFC3339, previousTelemetryList.Telemetry[previousTelemetryID-1].ServerTime); err != nil {
			log.Printf("Error parsing time: %s", err)
			tools.WriteJsonInternalError(rw)
			return
		} else {
			lastTime = parsed
		}
	}

	currentTime := time.Now().UTC()
	differenceTime := currentTime.Sub(lastTime)
	if differenceTime.Seconds() < 10 {
		log.Printf(`The request has been ignored: %v passed since the previous request.
		Time to insert: %v. 
		Previous insertion ID: %v.`, differenceTime, lastTime.Add(10), previousTelemetryID)
		tools.WriteJsonResult(rw, "The server is updated every 10 seconds. Wait")
		return
	}

	if err := store.sendData(&tt); err != nil {
		log.Printf("Inserting error: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	} else {
		tools.WriteJsonOk(rw, "ok")
	}

}
