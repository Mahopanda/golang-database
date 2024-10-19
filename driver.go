package main

import (
	"sync"
)

// Driver 結構用於管理資料庫，包含互斥鎖和日誌
type Driver struct {
	mutex  sync.Mutex
	mutexs map[string]*sync.Mutex
	store  Storage
	log    Logger
}

// NewDriver 初始化一個新的 Driver，負責資料庫管理
func NewDriver(store Storage, log Logger) *Driver {
	return &Driver{
		mutexs: make(map[string]*sync.Mutex),
		store:  store,
		log:    log,
	}
}

// Write 實作將數據寫入資料庫
func (d *Driver) Write(collection, resource string, v interface{}) error {
	// 獲取或創建集合的互斥鎖
	// 這裡需要保證，當多個協程同時訪問或創建集合的互斥鎖時，不會發生競爭條件
	// 因此，我們會先使用全局互斥鎖來鎖定對 mutexs map的操作
	mutex := d.getOrCreateMutex(collection)
	// 鎖定這個集合專用的互斥鎖，保護該集合內的操作，確保同一時間只有一個協程可以寫入該集合
	mutex.Lock()// 開始鎖定集合的互斥鎖，保證接下來的操作是安全的
	defer mutex.Unlock()// 在寫入操作完成後解鎖，允許其他協程訪問該集合

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
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()

	return d.store.Delete(collection, resource)
}

// getOrCreateMutex 函數用於獲取或創建指定集合的互斥鎖
// 使用全局互斥鎖保護對 mutexs 地圖的訪問，以避免多個協程同時訪問或創建同一個集合的互斥鎖時出現競爭條件
/*
全局互斥鎖 (d.mutex)：這個互斥鎖是用來保護對 mutexs 地圖的讀寫操作的。因為 mutexs map是共享的，
如果多個協程同時嘗試創建或訪問該地圖，可能會發生資料競爭，因此需要鎖定。

集合專用互斥鎖 (d.mutexs[collection])：每個集合都有自己專用的互斥鎖，用來保護該集合內的操作
（例如讀取、寫入和刪除）。這樣可以保證，當有多個協程同時對同一個集合進行操作時，不會出現數據競爭，確保數據的一致性。
*/
func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	m, ok := d.mutexs[collection]
	if !ok {
		m = &sync.Mutex{}
		d.mutexs[collection] = m
	}
	return m
}
