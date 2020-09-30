package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	product "github.com/jekhedhri/tuto006/application/domain/product"
	autowired "github.com/jekhedhri/tuto006/application/infra/product"
)

//FindProduct is FindProduct
func FindProduct(c *gin.Context) {
	productService, _ := autowired.InjectProductService()
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	result, _ := productService.Find(idInt)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}

//CreateProduct is FindProduct
func CreateProduct(c *gin.Context) {
	productService, _ := autowired.InjectProductService()
	product := product.Product{0, "Clavier", "13.5"}
	result := productService.Create(&product)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": result})
}
