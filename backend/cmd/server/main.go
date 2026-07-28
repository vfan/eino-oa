package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuhaitao/eino-oa/backend/config"
	"github.com/wuhaitao/eino-oa/backend/internal/agent"
	"github.com/wuhaitao/eino-oa/backend/internal/handler"
	"github.com/wuhaitao/eino-oa/backend/internal/model"
	"github.com/wuhaitao/eino-oa/backend/internal/repository"
	"github.com/wuhaitao/eino-oa/backend/internal/service"
)

func main() {
	cfg := config.Load()
	ctx := context.Background()

	db, err := model.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	studentRepo := repository.NewStudentRepo(db)
	courseRepo := repository.NewCourseRepo(db)
	gradeRepo := repository.NewGradeRepo(db)

	studentSvc := service.NewStudentService(studentRepo)
	courseSvc := service.NewCourseService(courseRepo)
	gradeSvc := service.NewGradeService(gradeRepo)
	sessionSvc := service.NewSessionService(repository.NewSessionRepo(db))

	studentHandler := handler.NewStudentHandler(studentSvc)
	courseHandler := handler.NewCourseHandler(courseSvc)
	gradeHandler := handler.NewGradeHandler(gradeSvc)

	r := gin.Default()

	api := r.Group("/api")
	studentHandler.Register(api)
	courseHandler.Register(api)
	gradeHandler.Register(api)

	if cfg.DeepSeekKey != "" {
		chatAgent, err := agent.NewChatAgent(ctx, cfg.DeepSeekKey, cfg.DeepSeekModel, cfg.DeepSeekBase, studentSvc, courseSvc, gradeSvc)
		if err != nil {
			log.Printf("AI Agent 初始化失败: %v", err)
		} else {
			chatHandler := handler.NewChatHandler(chatAgent, sessionSvc)
			chatHandler.Register(r)
			log.Println("AI Agent 已启用")
		}
	} else {
		log.Println("未设置 DEEPSEEK_API_KEY，AI Agent 未启用")
		r.POST("/api/chat", func(c *gin.Context) {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "AI 助手未启用，请设置 DEEPSEEK_API_KEY 环境变量",
			})
		})
	}

	log.Printf("服务启动在 :%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("服务启动失败: %v", err)
	}
}
