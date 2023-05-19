package api

import (
	db "github.com/PigeonShogi/HR_system/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server 用以處理 HTTP 請求
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer 用以建立 HTTP 伺服器 並設定路由
func NewServer(store db.Store) *Server {
	router := gin.Default()
	server := &Server{store: store, router: router}

	// 找出一筆 employees 的記錄
	router.GET("/employees/:id", server.getEmployee)

	// 找出多筆 employees 的記錄
	router.GET("/employees", server.listEmployee)

	// employees 建立新記錄
	router.POST("/employees", server.createEmployee)

	return server
}

// 以特定位址啟用 HTTP 伺服器
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
