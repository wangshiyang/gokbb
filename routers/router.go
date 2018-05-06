package routers

import (
	"github.com/gin-gonic/gin"
	"shawn/gokbb/common/setting"
	"shawn/gokbb/routers/v1"
	"shawn/gokbb/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/auth", v1.GetAuth)

	apiv1 := r.Group("/api/v1")

	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)

		apiv1.POST("/tags", v1.AddTags)

		apiv1.PUT("/tags/:id", v1.EditTags)

		apiv1.DELETE("/tags/:id", v1.DeleteTags)

		apiv1.GET("/articles", v1.GetArticles)

		apiv1.GET("/articles/:id", v1.GetArticle)

		apiv1.POST("/articles", v1.AddArticle)

		apiv1.PUT("/articles/:id", v1.EditArticle)

		apiv1.DELETE("/articles/:id", v1.DeletedArticle)
	}

	return r
}
