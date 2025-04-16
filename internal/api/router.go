package api

import (
	"swift-parser/internal/database"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db *database.DB
}

func NewRouter(db *database.DB) *Router {
	return &Router{db: db}
}

func (r *Router) Setup() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/v1/swift-codes")
	{
		v1.GET("/:swiftCode", r.GetSWIFTCode)
		v1.GET("/country/:countryISO2", r.GetSWIFTCodesByCountry)
		v1.POST("", r.PostSWIFTCode)
		v1.DELETE("/:swiftCode", r.DeleteSWIFTCode)
	}

	return router
}
