package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	// ...
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=bogdan password=2310zavbdJ dbname=youtube_go_test_db sslmode=disable port=5432"
	}

	os.Exit(m.Run())
}
