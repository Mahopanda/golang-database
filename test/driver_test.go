package main_test

import (
	"testing"

	"github.com/Mahopanda/golang-database/database"
)

func TestDriver_WriteRead(t *testing.T) {
	logger := database.NewConsoleLogger()

	// 初始化 JSONSerializer 並傳遞給 NewFileStore
	serializer := &database.JSONSerializer{}
	store := database.NewFileStore("./testdata", serializer)

	// 初始化鎖管理器
	lockManager := database.NewLockManager()

	// 傳入 store、logger 和 lockManager 初始化 driver
	driver := database.NewDriver(store, lockManager, logger)

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
