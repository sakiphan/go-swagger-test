package main

import (
	"net/http"

	_ "goSwagger/docs" // Swaggo tarafından oluşturulan belgeleri dahil etmek için

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// pingEndpoint fonksiyonu için Swagger yorum satırları
// @Summary Ping example
// @Schemes
// @Description Does a ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /ping [get]
func pingEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// helloEndpoint fonksiyonu için Swagger yorum satırları
// @Summary Hello example
// @Description Returns a greeting
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /hello [get]
func helloEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	router := gin.Default()

	// Swagger UI'ı sunmak için route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// "/ping" ve "/hello" endpoint'lerini router'a ekleyin
	router.GET("/ping", pingEndpoint)
	router.GET("/hello", helloEndpoint)

	// Sunucuyu başlat
	router.Run(":8080")
}
