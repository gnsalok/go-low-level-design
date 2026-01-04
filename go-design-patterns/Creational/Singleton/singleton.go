package main

import "sync"

// NOTE: Singletons are rarely necessary in Go. Prefer passing dependencies
// explicitly. When needed, sync.Once provides safe lazy initialization.

type DatabaseConnection struct{}

func (d *DatabaseConnection) Query(sql string) {
	// Database operations
	_ = sql
}

var (
	dbInstance *DatabaseConnection
	once       sync.Once
)

func GetDatabaseConnection() *DatabaseConnection {
	once.Do(func() {
		dbInstance = &DatabaseConnection{}
	})
	return dbInstance
}

// Usage:
// db := GetDatabaseConnection()
// db.Query("SELECT * FROM users")
