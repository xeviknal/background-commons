package database

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	setupTestDatabase()
}

func teardown() {
	destroyDatabase()
}

func setupTestDatabase() {
	SetConnectionConfig("test", "test", "test")
}

func destroyDatabase() {
	Clean()
}
