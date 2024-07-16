package main

import (
	"go-api-starter/storage"
	"go-api-starter/types"
)

func main(){
	db := storage.NewSqlStorage().GetDB()

	db.AutoMigrate(&types.User{})
}