package database

import (
	"sync"
)

// LockManager 介面，定義了鎖管理的操作
type LockManager interface {
	GetLock(collection string) *sync.Mutex
}

// DefaultLockManager 是一個基於 map 實現的鎖管理器
type DefaultLockManager struct {
	mutex  sync.Mutex
	mutexs map[string]*sync.Mutex
}

// NewLockManager 初始化一個新的鎖管理器
func NewLockManager() *DefaultLockManager {
	return &DefaultLockManager{
		mutexs: make(map[string]*sync.Mutex),
	}
}

// GetLock 獲取或創建集合的互斥鎖
func (lm *DefaultLockManager) GetLock(collection string) *sync.Mutex {
	lm.mutex.Lock()
	defer lm.mutex.Unlock()

	m, ok := lm.mutexs[collection]
	if !ok {
		m = &sync.Mutex{}
		lm.mutexs[collection] = m
	}
	return m
}
