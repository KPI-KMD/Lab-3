package db

import "testing"

func TestDbConnection_ConnectionURL(t *testing.T) {
	conn := &Connection{
		DbName:     "Multimedia_class",
		User:       "postgres",
		Password:   "admin",
		Host:       "localhost",
		DisableSSL: true,
	}
	if conn.ConnectionURL() != "postgres://postgres:admin@localhost/Multimedia_class?sslmode=disable" {
		t.Error("Unexpected connection string")
	}
}
