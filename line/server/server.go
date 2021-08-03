package server

import (
	route "github.com/Filmptz/whalealert_linebot/line/routes"
	"github.com/gin-gonic/gin"
)

func setupServer(s *gin.RouterGroup) error{

	s.GET("/", route.GetStatusOK())
	s.POST("/webhook",route.UpdateMessages())

	return nil
}