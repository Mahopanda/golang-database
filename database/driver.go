package database

import (
	"encoding/json"
	"fmt"
)

// Driver 結構用於管理資料庫
type Driver struct {
	lockManager LockManager
	store       Storage
	log         Logger
}

// NewDriver 初始化一個新的 Driver
func NewDriver(store Storage, log Logger, lockManager LockManager) *Driver {
	return &Driver{
		store:       store,
		log:         log,
		lockManager: lockManager,
	}
}

// Write 實作將數據寫入資料庫
func (d *Driver) Write(collection, resource string, v interface{}) error {
	mutex := d.lockManager.GetLock(collection)
	mutex.Lock()
	defer mutex.Unlock()

	return d.store.Write(collection, resource, v)
}

// Read 實作從資料庫中讀取數據
func (d *Driver) Read(collection, resource string, v interface{}) error {
	return d.store.Read(collection, resource, v)
}

// ReadAll 實作讀取資料庫中的所有數據
func (d *Driver) ReadAll(collection string) ([]string, error) {
	return d.store.ReadAll(collection)
}

// Delete 實作從資料庫中刪除數據
func (d *Driver) Delete(collection, resource string) error {
	mutex := d.lockManager.GetLock(collection)
	mutex.Lock()
	defer mutex.Unlock()

	return d.store.Delete(collection, resource)
}

// Query 函數從指定集合中查詢符合條件的數據
func (d *Driver) Query(collection string, filter func(map[string]interface{}) bool) ([]map[string]interface{}, error) {
	records, err := d.store.ReadAll(collection)
	if err != nil {
		return nil, err
	}

	var results []map[string]interface{}
	for _, record := range records {
		var data map[string]interface{}
		err := json.Unmarshal([]byte(record), &data)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling record: %v", err)
		}
		if filter(data) {
			results = append(results, data)
		}
	}

	return results, nil
}
