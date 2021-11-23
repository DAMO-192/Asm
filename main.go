package main

import (
	"Asm/api"
	"Asm/databases"
)

func main() {
	databases.InitDB()
	api.StartService()
}
