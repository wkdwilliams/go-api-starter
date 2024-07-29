package main

import (
	"go-api-starter/internal/storage"
	"go-api-starter/internal/types"
)

func main(){
	db := storage.NewSqlStorage().GetDB()

	db.AutoMigrate(&types.User{})
}