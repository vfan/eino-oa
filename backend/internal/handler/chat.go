package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wuhaitao/eino-oa/backend/internal/agent"
	"github.com/wuhaitao/eino-oa/backend/internal/service"
)

type ChatHandler struct {
	chatAgent  *agent.ChatAgent
	sessionSvc *service.SessionService
}

func NewChatHandler(chatAgent *agent.ChatAgent, sessionSvc *service.SessionService) *ChatHandler {
	return &ChatHandler{chatAgent: chatAgent, sessionSvc: sessionSvc}
}

func (h *ChatHandler) Register(r *gin.Engine) {
	g := r.Group("/api/sessions")
	g.GET("", h.ListSessions)
	g.POST("", h.CreateSession)
	g.GET("/:id", h.GetSession)
	g.DELETE("/:id", h.DeleteSession)
	g.POST("/:id/chat", h.Chat)
}

func (h *ChatHandler) ListSessions(c *gin.Context) {
	sessions, err := h.sessionSvc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func (h *ChatHandler) CreateSession(c *gin.Context) {
	var req struct {
		Title string `json:"title"`
	}
	c.ShouldBindJSON(&req)
	if req.Title == "" {
		req.Title = "新对话"
	}

	session, err := h.sessionSvc.Create(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, session)
}

func (h *ChatHandler) GetSession(c *gin.Context) {
	session, err := h.sessionSvc.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "会话不存在"})
		return
	}
	c.JSON(http.StatusOK, session)
}

func (h *ChatHandler) DeleteSession(c *gin.Context) {
	if err := h.sessionSvc.Delete(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}

func (h *ChatHandler) Chat(c *gin.Context) {
	sessionID := c.Param("id")

	var req struct {
		Message string `json:"message"`
	}
	if err := c.ShouldBindJSON(&req); err != nil || req.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "消息不能为空"})
		return
	}

	session, err := h.sessionSvc.GetByID(sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "会话不存在"})
		return
	}

	// 如果是第一条消息，用消息内容作为会话标题
	if len(session.Messages) == 0 {
		title := req.Message
		if len([]rune(title)) > 30 {
			title = string([]rune(title)[:30]) + "..."
		}
		_ = h.sessionSvc.UpdateTitle(sessionID, title)
	}

	h.sessionSvc.AddMessage(sessionID, "user", req.Message)

	history := make([]agent.HistoryMessage, 0, len(session.Messages))
	for _, m := range session.Messages {
		history = append(history, agent.HistoryMessage{Role: m.Role, Content: m.Content})
	}

	chatReq := &agent.ChatRequest{
		Message: req.Message,
		History: history,
	}

	resp, err := h.chatAgent.Chat(c.Request.Context(), chatReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.sessionSvc.AddMessage(sessionID, "assistant", resp.Reply)

	c.JSON(http.StatusOK, resp)
}
