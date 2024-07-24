package controllers

import (
	"context"
	"net/http"
	"socialapp/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateLike(c *gin.Context) {
	var like models.Like
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := likeCollection.InsertOne(context.Background(), like)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"like_id": res.InsertedID})
}

func GetLike(c *gin.Context) {
	likeID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(likeID)

	var like models.Like
	err := likeCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&like)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Like not found"})
		return
	}

	c.JSON(http.StatusOK, like)
}

func DeleteLike(c *gin.Context) {
	likeID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(likeID)

	res, err := likeCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting like"})
		return
	}

	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Like not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like deleted successfully"})
}

func ListLikes(c *gin.Context) {
	cursor, err := likeCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching likes"})
		return
	}
	defer cursor.Close(context.Background())

	var likes []models.Like
	for cursor.Next(context.Background()) {
		var like models.Like
		cursor.Decode(&like)
		likes = append(likes, like)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating likes"})
		return
	}

	c.JSON(http.StatusOK, likes)
}
