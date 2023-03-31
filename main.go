package main

import (
	"gituhub.com/snirkop89/goly/app/model"
	"gituhub.com/snirkop89/goly/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
