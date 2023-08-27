package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {

	var err error 

	//Initialize SQLite
	db, err = InitialzeSqlite()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return nil
}

func GetSQLiteDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Initialize Logger

	logger = NewLogger(p)
	return logger
}