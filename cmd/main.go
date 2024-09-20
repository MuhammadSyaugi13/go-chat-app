package main

import (
	"server/db"
	"server/helper"
	"server/internal/user"
	"server/router"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn, err := db.NewDatabase()
	helper.PanicIfError(err, "Tidak bisa inisiasi database")

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	router.InitRouter(userHandler)
	router.Start("localhost:8090")

}
