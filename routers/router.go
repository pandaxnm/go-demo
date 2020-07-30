package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pandaxnm/go-gin-example/controllers/v1"
	"github.com/pandaxnm/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	tag := r.Group("/v1")
	{
		tag.GET("/tags", v1.GetTags)
		tag.POST("/tags", v1.AddTag)
		tag.PUT("/tags/:id", v1.EditTag)
		tag.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}
