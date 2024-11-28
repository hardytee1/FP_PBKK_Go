package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/middleware"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	r.Run()
}