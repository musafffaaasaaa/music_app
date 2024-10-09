package main

import (
	"music-app/docs"
	"music-app/handlers"
	"music-app/utils"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)


func main() {
	utils.LoadEnv()       
	utils.InitLogger()    
	utils.ConnectDB()     
	defer utils.CloseDB() 

	r := gin.Default()

	
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Добро пожаловать в API Music App!"})
	})
	r.GET("/songs", handlers.GetLibrary)
	r.POST("/songs", handlers.AddSong)
	r.GET("/songs/:id/lyrics", handlers.GetSongText)
	r.PUT("/songs/:id", handlers.UpdateSong)
	r.DELETE("/songs/:id", handlers.DeleteSong)

	
	r.Run(":8080")
}
