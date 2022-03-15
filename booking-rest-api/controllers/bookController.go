package controllers

import (
	"strconv"
	"webapi-with-go/database"
	"webapi-with-go/models"

	"github.com/gin-gonic/gin"
)

func ShowBook(c *gin.Context) {

	bookId := c.Param("id")

	newId, err := strconv.Atoi(bookId)

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Id has to be integer.",
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, newId).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}
