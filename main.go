package main

import (
    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "webdev-intern-assignment/docs" // swagger generated files
    "webdev-intern-assignment/database"
    "webdev-intern-assignment/controllers"
    "webdev-intern-assignment/middleware"
    "webdev-intern-assignment/models"
)

// @title Authentication API
// @version 1.0
// @description API for authentication with role-based access control.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
    database.Connect()
    
    r := gin.Default()

    // Swagger route
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // Public routes
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    // Protected routes
    protected := r.Group("/api")
    protected.Use(middleware.AuthMiddleware())
    {
        // Admin routes
        admin := protected.Group("/admin")
        admin.Use(middleware.RoleAuth(models.AdminRole))
        {
            // @Summary Get admin dashboard
            // @Description Get admin dashboard data
            // @Tags Admin
            // @Security BearerAuth
            // @Success 200 {object} map[string]interface{}
            // @Failure 401 {object} map[string]string
            // @Failure 403 {object} map[string]string
            // @Router /api/admin/dashboard [get]
            admin.GET("/dashboard", func(c *gin.Context) {
                c.JSON(200, gin.H{"message": "Admin dashboard"})
            })
        }

        // Editor routes
        editor := protected.Group("/editor")
        editor.Use(middleware.RoleAuth(models.AdminRole, models.EditorRole))
        {
            // @Summary Get editor content
            // @Description Get editor content data
            // @Tags Editor
            // @Security BearerAuth
            // @Success 200 {object} map[string]interface{}
            // @Failure 401 {object} map[string]string
            // @Failure 403 {object} map[string]string
            // @Router /api/editor/content [get]
            editor.GET("/content", func(c *gin.Context) {
                c.JSON(200, gin.H{"message": "Editor content"})
            })
        }

        // Reader routes
        reader := protected.Group("/reader")
        reader.Use(middleware.RoleAuth(models.AdminRole, models.EditorRole, models.ReaderRole))
        {
            // @Summary Get reader articles
            // @Description Get reader articles data
            // @Tags Reader
            // @Security BearerAuth
            // @Success 200 {object} map[string]interface{}
            // @Failure 401 {object} map[string]string
            // @Failure 403 {object} map[string]string
            // @Router /api/reader/articles [get]
            reader.GET("/articles", func(c *gin.Context) {
                c.JSON(200, gin.H{"message": "Reader articles"})
            })
        }
    }

    r.Run(":8080")
}