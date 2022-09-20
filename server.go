package gin_server

import (
	"github.com/Inf85/gin_server/helpers"
	"github.com/Inf85/gin_server/log"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"os"
)

func StartServer() *gin.Engine {
	_ = godotenv.Load()
	var err error

	r := gin.Default()

	port := helpers.GetEnvDefault("PORT", "3000")
	if err = r.Run(":" + port); err != nil {
		log.GetLogger().Errorf("Can not start server: %s", err.Error())
		os.Exit(1)
	}

	return r
}
