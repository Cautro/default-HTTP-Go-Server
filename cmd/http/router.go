package http

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type UseCase struct {
	DB *sql.DB 
}

func NewRouter(u UseCase) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	registerPublicRouters(u, router)

	return router
}

func registerPublicRouters(u UseCase, r *gin.Engine) {
	//routes
}