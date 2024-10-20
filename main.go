package main

import (
	"encoding/json"
	"fmt"

	"github.com/Mahopanda/golang-database/database"
)

const Version = "0.0.1"

// 主函數
func main() {

	logger := database.NewConsoleLogger()
	serializer := &database.JSONSerializer{}
	// 初始化文件存儲
	store := database.NewFileStore("storage/", serializer)
	lockManager := database.NewLockManager()
	// 初始化資料庫驅動
	db := database.NewDriver(store, logger,lockManager)

	// 初始化使用者數據
	employee := []map[string]interface{}{
		{
			"Name":    "John",
			"Age":     25,
			"Contact": "1234567890",
			"Company": "Google",
			"Address": map[string]interface{}{
				"City":    "Bangalore",
				"State":   "Karnataka",
				"Country": "India",
				"Pincode": "560001",
			},
		},
		{
			"Name":    "Doe",
			"Age":     30,
			"Contact": "1234567890",
			"Company": "Microsoft",
			"Address": map[string]interface{}{
				"City":    "Hyderabad",
				"State":   "Telangana",
				"Country": "India",
				"Pincode": "500001",
			},
		},
		{
			"Name":    "Smith",
			"Age":     35,
			"Contact": "1234567890",
			"Company": "Amazon",
			"Address": map[string]interface{}{
				"City":    "Chennai",
				"State":   "Tamilnadu",
				"Country": "India",
				"Pincode": "600001",
			},
		},
		{
			"Name":    "Tom",
			"Age":     40,
			"Contact": "1234567890",
			"Company": "Facebook",
			"Address": map[string]interface{}{
				"City":    "Mumbai",
				"State":   "Maharashtra",
				"Country": "India",
				"Pincode": "400001",
			},
		},
		{
			"Name":    "Jerry",
			"Age":     45,
			"Contact": "1234567890",
			"Company": "Apple",
			"Address": map[string]interface{}{
				"City":    "Pune",
				"State":   "Maharashtra",
				"Country": "India",
				"Pincode": "411001",
			},
		},
		{
			"Name":    "Mickey",
			"Age":     50,
			"Contact": "1234567890",
			"Company": "Tesla",
			"Address": map[string]interface{}{
				"City":    "Kolkata",
				"State":   "West Bengal",
				"Country": "India",
				"Pincode": "700001",
			},
		},
	}

	// 將使用者數據寫入資料庫
	for _, emp := range employee {
		if err := db.Write("users", emp["Name"].(string), emp); err != nil {
			logger.Error("Error writing user: %s, %v", emp["Name"], err)
		}
	}

	// 刪除使用者
	if err := db.Delete("users", "John"); err != nil {
		logger.Error("Error deleting user John: %v", err)
	}

	// 讀取所有使用者數據
	records, err := db.ReadAll("users")
	if err != nil {
		logger.Error("Error reading users: %v", err)
	}

	// 輸出所有使用者數據
	for _, record := range records {
		var user map[string]interface{}
		if err := json.Unmarshal([]byte(record), &user); err != nil {
			logger.Error("Error unmarshalling user: %v", err)
		}
		fmt.Printf("User: %+v\n", user)
	}

	// 查詢所有年齡為 30 的使用者
	results, err := db.Query("users", func(data map[string]interface{}) bool {
		age, ok := data["Age"].(float64) // JSON 解析時，數值類型會被解釋為 float64
		return ok && age == 30
	})

	if err != nil {
		logger.Error("Error querying users: %v", err)
	}

	// 打印查詢結果
	fmt.Println("Users with Age 30:")
	for _, user := range results {
		fmt.Printf("User: %+v\n", user)
	}
}
