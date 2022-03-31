package main

import (
	"net/http"
	"test_swag2/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /api/v1
func main() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Swagger Example API"

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/hello/:name", HelloHandler)
	}
	r.Run("localhost:8080")
}

type User struct {
	Id   int    `json:"id", example:"1"`     // 유저ID
	Name string `json:"name" example:"John"` // 이름
	Age  int    `json:"age" example:"10"`    // 나이
}

// HelloHandler godoc
// @Summary 요약 기재
// @Description 상세 설명 기재
// @name get-string-by-int
// @Accept json
// @Produce json
// @Param name path string true "User name"
// @Router /api/v1/hello/{name} [get]
// @Success 200 {object} User
// @Failure 400
func HelloHandler(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"user": ""})
	} else {
		user := User{Id: 1, Name: name, Age: 20}
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}
