package main

import (
	"github.com/gin-gonic/gin"

	"test_swag/docs"
	"test_swag/model"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/welcome/:name", welcomePathParam)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// Welcome godoc
// @Summary 요약
// @Description 설명
// @name get-string-by-int
// @Accept  json
// @Produce  json
// @Param name path string true "User name"
// @Router /welcome/{name} [get]
// @Success 200 {object} welcomeModel
func welcomePathParam(c *gin.Context) {
	name := c.Param("name")
	welcomeMessage := model.User{ID: 1, Name: name}

	c.JSON(200, gin.H{"message": welcomeMessage})
}

func main() {
	// programatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"

	r := setupRouter()

	r.Run()
}
