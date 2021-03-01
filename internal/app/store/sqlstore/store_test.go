package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

// TestMain вызываеться один раз в конкретном пакете и это
// отличное место чтоб сделать разовые манипуляции
func TestMain(m *testing.M)  {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=root password=root dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
