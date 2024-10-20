package main_test

import (
	"os"
	"testing"

	"github.com/Mahopanda/golang-database/database"
)

func TestFileStore_WriteRead(t *testing.T) {
	logger := database.NewConsoleLogger()
	store := database.NewFileStore("./testdata", logger)
	defer os.RemoveAll("./testdata") // 測試結束後刪除測試資料夾

	// 使用 map[string]interface{} 來表示動態數據
	user := map[string]interface{}{
		"Name": "John",
		"Age":  "25",
	}

	// 將用戶數據寫入文件
	err := store.Write("users", "John", user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// 讀取用戶數據
	var result map[string]interface{}
	err = store.Read("users", "John", &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// 驗證結果
	if result["Name"] != "John" || result["Age"] != "25" {
		t.Fatalf("unexpected result, got %+v", result)
	}
}
