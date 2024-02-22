package main

import (
	"ecommerceGo/internal/database"
	"ecommerceGo/internal/server"
)

func main() {
	database.InitDb()
	server.Run()
}
