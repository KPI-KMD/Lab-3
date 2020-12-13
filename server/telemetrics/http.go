package telemetrics

import (
	"encoding/json"
	"github.com/KPI-KMD/Lab-3/server/tools"
	"log"
	"net/http"
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

func handleListTelemetries(r *http.Request, rw http.ResponseWriter, store *Store,) {
	var idTabl int64
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
 var sdata SendData
 if err := json.NewDecoder(r.Body).Decode(&sdata); err != nil {
  log.Printf("Error decoding channel input: %s", err)
  tools.WriteJsonBadRequest(rw, "bad JSON payload")
  return
 }
 err := store.sendData(&sdata)
 if err != nil {
  log.Printf("Error inserting record: %s", err)
  tools.WriteJsonInternalError(rw)
 } else {
  tools.WriteJsonOk(rw, "ok")
 }
 time.Sleep(10 * time.Second)
}