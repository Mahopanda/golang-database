package main_test

import (
	"testing"

	"github.com/Mahopanda/golang-database/database"
)

func TestDriver_WriteRead(t *testing.T) {
	logger := database.NewConsoleLogger()
	store := database.NewFileStore("./testdata", logger)
	driver := database.NewDriver(store, logger)

	user := database.User{
		Name: "John",
		Age:  "25",
	}

	err := driver.Write("users", "John", user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result database.User
	err = driver.Read("users", "John", &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Name != "John" || result.Age != "25" {
		t.Fatalf("unexpected result, got %+v", result)
	}
}
