package database

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Storage 介面，定義資料存取的基本操作
type Storage interface {
	Write(collection, resource string, v interface{}) error
	Read(collection, resource string, v interface{}) error
	ReadAll(collection string) ([]string, error)
	Delete(collection, resource string) error
}

// FileStore 負責具體的文件操作
type FileStore struct {
	dir string
	log Logger
}

// NewFileStore 初始化一個新的 FileStore，負責文件存取
func NewFileStore(dir string, log Logger) *FileStore {
	return &FileStore{
		dir: filepath.Clean(dir),
		log: log,
	}
}

// Write 實作將數據寫入文件的操作
func (fs *FileStore) Write(collection, resource string, v interface{}) error {
	dir := filepath.Join(fs.dir, collection)
	fnlPath := filepath.Join(dir, resource+".json")
	tmpPath := fnlPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// 將 interface{} 序列化為 JSON
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	// 將結果寫入臨時文件
	if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	// 確保寫入過程的原子性，重命名臨時文件
	return os.Rename(tmpPath, fnlPath)
}

// Read 實作從文件中讀取數據
func (fs *FileStore) Read(collection, resource string, v interface{}) error {
	path := filepath.Join(fs.dir, collection, resource+".json")
	if _, err := os.Stat(path); err != nil {
		return err
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// 將文件內容解析並填充到傳入的 v 變量
	return json.Unmarshal(b, v)
}

// ReadAll 實作讀取集合中的所有文件
func (fs *FileStore) ReadAll(collection string) ([]string, error) {
	dir := filepath.Join(fs.dir, collection)

	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var records []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		b, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}
		records = append(records, string(b))
	}
	return records, nil
}

// Delete 實作刪除指定文件
func (fs *FileStore) Delete(collection, resource string) error {
	path := filepath.Join(fs.dir, collection, resource+".json")
	return os.Remove(path)
}
