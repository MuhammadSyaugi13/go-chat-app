package router

import (
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
	r.GET("/ws/get-rooms", wsHandler.GetRooms)
	r.GET("/ws/get-clients", wsHandler.GetClients)
	r.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)

}

func Start(addr string) error {
	return r.Run(addr)
}
