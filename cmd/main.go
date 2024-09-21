package main

import (
	"fmt"
	"server/db"
	"server/helper"
	"server/internal/user"
	"server/internal/ws"
	"server/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn, err := db.NewDatabase()
	helper.PanicIfError(err, "Tidak bisa inisiasi database")

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	router.InitRouter(userHandler, wsHandler)

	fmt.Println("hub.Roomssss : ")

	router.Start("localhost:8090")

}
