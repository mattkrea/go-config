package config

import (
	"testing"
)

type Configuration struct {
	Server   map[string]interface{} `json:"server"`
	Database map[string]interface{} `json:"database"`
}

func TestOpeningConfig(t *testing.T) {

	var c Configuration

	err := Open(&c)
	if err != nil {
		t.FailNow()
	}

	if c.Server["port"] != ":8080" {
		t.FailNow()
	}

	if c.Database["host"] != "localhost" {
		t.FailNow()
	}
}
