package main

import (
	"os"
	"testing"
)

func TestFileStore_WriteRead(t *testing.T) {
	logger := NewConsoleLogger()
	store := NewFileStore("./testdata", logger)
	defer os.RemoveAll("./testdata")

	user := User{
		Name: "John",
		Age:  "25",
	}

	err := store.Write("users", "John", user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	var result User
	err = store.Read("users", "John", &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Name != "John" || result.Age != "25" {
		t.Fatalf("unexpected result, got %+v", result)
	}
}
