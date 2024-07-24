package controllers

import (
	"context"
	"net/http"
	"socialapp/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return	
	}

	res, err := commentCollection.InsertOne(context.Background(), comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment_id": res.InsertedID})
}

func GetComment(c *gin.Context) {
	commentID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(commentID)

	var comment models.Comment
	err := commentCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&comment)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func UpdateComment(c *gin.Context) {
	commentID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(commentID)

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := commentCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": comment})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating comment"})
		return
	}

	if res.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func DeleteComment(c *gin.Context) {
	commentID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(commentID)

	res, err := commentCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting comment"})
		return
	}

	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func ListComments(c *gin.Context) {
	cursor, err := commentCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching comments"})
		return
	}
	defer cursor.Close(context.Background())

	var comments []models.Comment
	for cursor.Next(context.Background()) {
		var comment models.Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
