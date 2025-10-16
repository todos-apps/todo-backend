package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todo-backend/models"
)

// Handler holds DB reference
type Handler struct {
    DB *gorm.DB
}

// NewHandler returns a Handler
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{DB: db}
}

// CreateTodo POST /todos
func (h *Handler) CreateTodo(c *gin.Context) {
    var input struct {
        Title       string `json:"title" binding:"required"`
        Description string `json:"description"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    todo := models.Todo{
        Title:       input.Title,
        Description: input.Description,
    }

    if err := h.DB.Create(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to create todo"})
        return
    }

    c.JSON(http.StatusCreated, todo)
}

// ListTodos GET /todos
func (h *Handler) ListTodos(c *gin.Context) {
    var todos []models.Todo
    if err := h.DB.Order("created_at desc").Find(&todos).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch todos"})
        return
    }
    c.JSON(http.StatusOK, todos)
}

// GetTodo GET /todos/:id
func (h *Handler) GetTodo(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var todo models.Todo
    if err := h.DB.First(&todo, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch todo"})
        return
    }

    c.JSON(http.StatusOK, todo)
}

// UpdateTodo PUT /todos/:id
func (h *Handler) UpdateTodo(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    var payload struct {
        Title       *string `json:"title"`
        Description *string `json:"description"`
        Completed   *bool   `json:"completed"`
    }
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var todo models.Todo
    if err := h.DB.First(&todo, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusNotFound, gin.H{"error": "todo not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to fetch todo"})
        return
    }

    // apply changes only if fields are present (pointer fields)
    if payload.Title != nil {
        todo.Title = *payload.Title
    }
    if payload.Description != nil {
        todo.Description = *payload.Description
    }
    if payload.Completed != nil {
        todo.Completed = *payload.Completed
    }

    if err := h.DB.Save(&todo).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update todo"})
        return
    }
    c.JSON(http.StatusOK, todo)
}

// DeleteTodo DELETE /todos/:id
func (h *Handler) DeleteTodo(c *gin.Context) {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
        return
    }

    if err := h.DB.Delete(&models.Todo{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete todo"})
        return
    }
    c.Status(http.StatusNoContent)
}
