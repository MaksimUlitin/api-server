package store_test

import (
	"os"
	"testing"
)

var (
	DataBaseURL string
)

func TestMain(m *testing.M) {
	DataBaseURL = os.Getenv("DATABASE_URL")
	if DataBaseURL == "" {
		DataBaseURL = "host=localhost dbname=restapi sslmode=disamble"
	}
	os.Exit(m.Run())
}
