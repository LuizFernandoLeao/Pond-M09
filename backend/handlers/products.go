package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "secure-web-app/backend/models" 
    "secure-web-app/backend/utils" 
)

func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := utils.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create product"})
        return
    }

    c.JSON(http.StatusCreated, product)
}

func GetProducts(c *gin.Context) {
    var products []models.Product
    if err := utils.DB.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve products"})
        return
    }

    c.JSON(http.StatusOK, products)
}

func UpdateProduct(c *gin.Context) {
    var product models.Product
    if err := utils.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := utils.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update product"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
    if err := utils.DB.Delete(&models.Product{}, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}