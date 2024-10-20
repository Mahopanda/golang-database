package main_test

import (
	"testing"

	"github.com/Mahopanda/golang-database/database"
)

func TestDriver_WriteRead(t *testing.T) {
	logger := database.NewConsoleLogger()
	store := database.NewFileStore("./testdata", logger)
	driver := database.NewDriver(store, logger)

	// 使用 map[string]interface{} 來表示動態數據
	user := map[string]interface{}{
		"Name": "John",
		"Age":  "25",
	}

	// 將用戶數據寫入數據庫
	err := driver.Write("users", "John", user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// 讀取用戶數據
	var result map[string]interface{}
	err = driver.Read("users", "John", &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// 驗證結果
	if result["Name"] != "John" || result["Age"] != "25" {
		t.Fatalf("unexpected result, got %+v", result)
	}
}
