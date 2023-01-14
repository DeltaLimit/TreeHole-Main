package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

type TestUser struct {
	ID       uint
	EMail    string
	Password string
}

func TestConnectDB(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	HandleError(err)

	db.AutoMigrate(&TestUser{})
	db.Create(&TestUser{ID: 1, EMail: "123@test.cn", Password: "114514"})
}
