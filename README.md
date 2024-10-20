# Go Key-Value Database

This project is a simple key-value store database implemented in Go, inspired by NoSQL databases like MongoDB. It supports dynamic data storage using `map[string]interface{}` and provides basic CRUD (Create, Read, Update, Delete) operations, along with query functionality.

## Features

- **Dynamic Schema**: Stores data without requiring predefined models or structures.
- **CRUD Operations**: Supports Create, Read, Update, Delete operations for collections.
- **Query Functionality**: Query collections using a custom filter function to search for records.
- **File-based Storage**: Data is stored in JSON format in the local file system.
- **Concurrency Safe**: Uses `sync.Mutex` to manage concurrent access to collections.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Mahopanda/golang-database.git
   ```
2. Navigate to the project directory:
   
  ```bash
  cd golang-database
  ```

3. Install dependencies (if any):
   ```base
   go mod tidy
   ```

## Usage
### To use the key-value store in your project, follow these steps:

1. Initialize the Logger, Storage, and Driver:

```go
logger := database.NewConsoleLogger()
store := database.NewFileStore("./storage", logger)
db := database.NewDriver(store, logger)
```

2. Insert Data:
```go
user := map[string]interface{}{
    "Name":    "John",
    "Age":     30,
    "Contact": "john@example.com",
}

err := db.Write("users", "john", user)
if err != nil {
    logger.Error("Failed to write user:", err)
}
```

3. Read Data:

```go
var result map[string]interface{}
err := db.Read("users", "john", &result)
if err != nil {
    logger.Error("Failed to read user:", err)
}
```

4. Query Data:
```go   
results, err := db.Query("users", func(data map[string]interface{}) bool {
    return data["Age"] == 30
})
```

5. Delete Data:
```go
err := db.Delete("users", "john")
if err != nil {
    logger.Error("Failed to delete user:", err)
}
```

## Running Tests
To run the unit tests for the project, execute the following command:
```bash
go test ./...
```

## Folder Structure
```bash
├── database/           # Core logic for the key-value database
│   ├── driver.go       # Handles high-level operations and concurrency control
│   ├── storage.go      # File-based storage implementation
│   ├── lock_manager.go # 包含 LockManager 的邏輯
│   └── logger.go       # Logger implementation
│   └── serializer.go   # 包含 Serializer 的邏輯
├── test/               # Test cases for the database
└── main.go             # Example usage of the key-value database
```
