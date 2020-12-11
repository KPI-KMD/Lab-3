// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/KPI-KMD/Lab-3/server/telemetrics"
)

// Injectors from modules.go:

// ComposeApiServer will create an instance of TelemetryApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*TelemetryApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := telemetrics.NewStore(db)
	httpHandlerFunc := telemetrics.HttpHandler(store)
	telemetryApiServer := &TelemetryApiServer{
		Port:             port,
		TelemetryHandler: httpHandlerFunc,
	}
	return telemetryApiServer, nil
}
