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

	// 初始化文件存儲
	fileStore := database.NewFileStore("storage/", logger)

	// 初始化資料庫驅動
	db := database.NewDriver(fileStore, logger)

	// 初始化使用者數據
	employee := []database.User{
		{"John", "25", "1234567890", "Google", database.Address{"Bangalore", "Karnataka", "India", "560001"}},
		{"Doe", "30", "1234567890", "Microsoft", database.Address{"Hyderabad", "Telangana", "India", "500001"}},
		{"Smith", "35", "1234567890", "Amazon", database.Address{"Chennai", "Tamilnadu", "India", "600001"}},
		{"Tom", "40", "1234567890", "Facebook", database.Address{"Mumbai", "Maharashtra", "India", "400001"}},
		{"Jerry", "45", "1234567890", "Apple", database.Address{"Pune", "Maharashtra", "India", "411001"}},
		{"Mickey", "50", "1234567890", "Tesla", database.Address{"Kolkata", "West Bengal", "India", "700001"}},
	}

	// 將使用者數據寫入資料庫
	for _, emp := range employee {
		if err := db.Write("users", emp.Name, emp); err != nil {
			logger.Error("Error writing user: %s, %v", emp.Name, err)
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
		var user database.User
		if err := json.Unmarshal([]byte(record), &user); err != nil {
			logger.Error("Error unmarshalling user: %v", err)
		}
		fmt.Printf("User: %+v\n", user)
	}

}
