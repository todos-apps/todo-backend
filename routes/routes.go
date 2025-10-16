package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"todo-backend/handlers"
)

func Setup(db *gorm.DB) *gin.Engine {
    r := gin.Default()

	// Configure CORS
	//allow all origins
	r.Use(cors.Default())

    // r.Use(cors.New(cors.Config{
    //     AllowOrigins:     []string{"http://localhost:5173"}, // your Vue dev server
    //     AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    //     AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    //     ExposeHeaders:    []string{"Content-Length"},
    //     AllowCredentials: true,
    //     MaxAge:           12 * time.Hour,
    // }))

    h := handlers.NewHandler(db)
	
	//define routes
    api := r.Group("/api")
    {
        todos := api.Group("/todos")
        {
            todos.POST("", h.CreateTodo)
            todos.GET("", h.ListTodos)
            todos.GET("/:id", h.GetTodo)
            todos.PUT("/:id", h.UpdateTodo)
            todos.DELETE("/:id", h.DeleteTodo)
        }
    }

    // simple health check
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    return r
}
