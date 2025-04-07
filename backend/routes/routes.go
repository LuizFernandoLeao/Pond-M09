package routes

import (
    "github.com/gin-gonic/gin"
    "secure-web-app/backend/handlers"
)

func SetupRoutes(router *gin.Engine) {
    userRoutes := router.Group("/users")
    {
        userRoutes.POST("/", handlers.CreateUser)
        userRoutes.GET("/", handlers.GetUsers)
        userRoutes.PUT("/:id", handlers.UpdateUser)
        userRoutes.DELETE("/:id", handlers.DeleteUser)
    }

    productRoutes := router.Group("/products")
    {
        productRoutes.POST("/", handlers.CreateProduct)
        productRoutes.GET("/", handlers.GetProducts)
        productRoutes.PUT("/:id", handlers.UpdateProduct)
        productRoutes.DELETE("/:id", handlers.DeleteProduct)
    }
}