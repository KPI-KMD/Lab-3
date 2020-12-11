//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/KPI-KMD/Lab-3/server/telemetrics"
)

// ComposeApiServer will create an instance of TelemetryApiServer according to providers defined in this file.
func ComposeApiServer(port HttpPortNumber) (*TelemetryApiServer, error) {
	wire.Build(
		// DB connection provider (defined in main.go).
		NewDbConnection,
		// Add providers from telemetrics package.
		telemetrics.Providers,
		// Provide TelemetryApiServer instantiating the structure and injecting channels handler and port number.
		wire.Struct(new(TelemetryApiServer), "Port", "TelemetryHandler"),
	)
	return nil, nil
}