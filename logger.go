package main

import (
	"github.com/jcelliott/lumber"
)

// Logger 介面，用於不同級別的日誌記錄
type Logger interface {
	Fatal(string, ...interface{})
	Error(string, ...interface{})
	Warn(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Trace(string, ...interface{})
}

// NewConsoleLogger 創建一個默認的 Console Logger
func NewConsoleLogger() Logger {
	return lumber.NewConsoleLogger(lumber.INFO)
}
