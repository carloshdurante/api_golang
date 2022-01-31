package controllers

import (
	"strconv"

	"github.com/carloshdurante/api_golang/database"
	"github.com/carloshdurante/api_golang/models"
	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := database.GetDB()

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)

	db := database.GetDB()

	err := db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, book)
}

func ShowAllBooks(c *gin.Context) {
	db := database.GetDB()

	var books []models.Book
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"books": books,
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	var book models.Book
	db := database.GetDB()

	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.BindJSON(&book)

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	db := database.GetDB()

	var book models.Book
	err = db.First(&book, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = db.Delete(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, book)
}
