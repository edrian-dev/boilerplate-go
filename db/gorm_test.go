package db

import "testing"

func TestNewConnection(t *testing.T) {
	tests := []struct {
		name string
	}{{
		name: "Prueba conexión con Postgres",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewConnection()
		})
	}
}
