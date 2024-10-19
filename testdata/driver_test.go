package main

import (
	"testing"
)

func TestDriver_WriteRead(t *testing.T) {
	logger := NewConsoleLogger()
	store := NewFileStore("./testdata", logger)
	driver := NewDriver(store, logger)

	user := User{
		Name: "John",
		Age:  "25",
	}

	err := driver.Write("users", "John", user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result User
	err = driver.Read("users", "John", &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Name != "John" || result.Age != "25" {
		t.Fatalf("unexpected result, got %+v", result)
	}
}
