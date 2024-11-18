package main

import (
	"go-api-project-seed/internal/api/v1"
	"go-api-project-seed/internal/middleware"
	"go-api-project-seed/internal/repository"
	"go-api-project-seed/internal/service"
	"go-api-project-seed/internal/utils"
	"go-api-project-seed/pkg/redis"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置
	utils.InitConfig()

	// 初始化日志
	logger := utils.InitLogger()

	// 初始化数据库
	db := repository.InitDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 初始化 Redis
	rdb := redis.InitRedis()
	defer rdb.Close()

	// 初始化 Gin
	router := gin.Default()
	router.Use(middleware.LoggerMiddleware(logger))
	router.Use(middleware.CORSMiddleware())

	// 初始化服务
	sampleService := service.NewSampleService(repository.NewSampleRepository(db))
	// 注册路由
	apiV1 := router.Group("/api/v1")
	v1.RegisterSampleRoutes(apiV1, sampleService)

	// 启动服务
	port := viper.GetString("server.port")
	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
