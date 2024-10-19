package main

import "encoding/json"

// Address 結構表示使用者的地址信息
type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

// User 結構表示使用者信息
type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}
