package router

import (
	"fmt"
	"server/internal/user"
	"server/internal/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
	r.GET("/user/:email", userHandler.GetUserByEmail)

	r.POST("/ws/create-room", wsHandler.CreateRoom)
	r.GET("/ws/get-room", wsHandler.GetRoom)

	fmt.Println("jalan init touter boss")

}

func Start(addr string) error {
	return r.Run(addr)
}
