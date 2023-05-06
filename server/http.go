package server

import (
	"backend/module/router"

	"github.com/gin-gonic/gin"
)

func NewServer() {
	//PORT := os.Getenv("PORT")
	ctrl := NewSetup()
	rt := gin.Default()
	//init middleware

	rt.Use(gin.Recovery(), gin.Logger())
	router.NewRouter(rt, ctrl.GuruCtrl, ctrl.SiswaCtrl)
	rt.Run(":8080")
}
