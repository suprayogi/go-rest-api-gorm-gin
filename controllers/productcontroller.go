package productcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suprayogi/go-rest-api/models"
)

func Index(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)
	// models.DB.Debug().Where("isdeleted<>? or isdeleted is null", 1).Select([]string{"id", "namaproduct"}).Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"msg":      "Product List",
		"products": products,
	})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBind(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusCreated, gin.H{
		"msg":     "Product Created",
		"product": product,
	})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBind(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	if models.DB.Model(&product).Where("id=?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"msg": "Cannot update product",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg":     "Product Updated",
		"product": product,
	})
}

func Delete(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBind(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if models.DB.Model(&product).Where("id=?", id).Delete(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"msg": "Cannot update product",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"msg":     "Product Deleted",
		"product": product,
	})
}
