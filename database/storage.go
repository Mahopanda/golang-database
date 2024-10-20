package database

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Reader 介面，定義讀取操作
type Reader interface {
	Read(collection, resource string, v interface{}) error
	ReadAll(collection string) ([]string, error)
}

// Writer 介面，定義寫入操作
type Writer interface {
	Write(collection, resource string, v interface{}) error
	Delete(collection, resource string) error
}

// Storage 接口
type Storage interface {
	Reader
	Writer
}

// Serializer 定義序列化接口
type Serializer interface {
	Serialize(v interface{}) ([]byte, error)
	Deserialize(data []byte, v interface{}) error
}

// JSONSerializer 負責 JSON 格式的序列化
type JSONSerializer struct{}

func (j *JSONSerializer) Serialize(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

func (j *JSONSerializer) Deserialize(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// FileStore 負責具體的文件操作
type FileStore struct {
	dir        string
	serializer Serializer
}

// NewFileStore 初始化 FileStore
func NewFileStore(dir string, serializer Serializer) *FileStore {
	return &FileStore{
		dir:        filepath.Clean(dir),
		serializer: serializer,
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

	b, err := fs.serializer.Serialize(v)
	if err != nil {
		return err
	}

	if err := os.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmpPath, fnlPath)
}

// Read 實作從文件中讀取數據
func (fs *FileStore) Read(collection, resource string, v interface{}) error {
	path := filepath.Join(fs.dir, collection, resource+".json")
	if _, err := os.Stat(path); err != nil {
		return err
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return fs.serializer.Deserialize(b, v)
}

// ReadAll 實作讀取集合中的所有文件
func (fs *FileStore) ReadAll(collection string) ([]string, error) {
	dir := filepath.Join(fs.dir, collection)

	if _, err := os.Stat(dir); err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var records []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dir, file.Name()))
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
