package controllers

import (
	"context"
	"net/http"
	"socialapp/models"
	"socialapp/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection = utils.GetCollection("posts")
var commentCollection *mongo.Collection = utils.GetCollection("comments")
var likeCollection *mongo.Collection = utils.GetCollection("likes")

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := postCollection.InsertOne(context.Background(), post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while inserting post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post_id": res.InsertedID})
}

func GetPost(c *gin.Context) {
	postID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(postID)

	var post models.Post
	err := postCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	commentCount, err := commentCollection.CountDocuments(context.Background(), bson.M{"post_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting comments"})
		return
	}

	likeCount, err := likeCollection.CountDocuments(context.Background(), bson.M{"post_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting likes"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post":         post,
		"commentCount": commentCount,
		"likeCount":    likeCount,
	})
}

func UpdatePost(c *gin.Context) {
	postID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(postID)

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := postCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": post})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating post"})
		return
	}

	if res.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}

func DeletePost(c *gin.Context) {
	postID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(postID)

	res, err := postCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting post"})
		return
	}

	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func ListPosts(c *gin.Context) {
	cursor, err := postCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching posts"})
		return
	}
	defer cursor.Close(context.Background())

	var posts []models.Post
	for cursor.Next(context.Background()) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error iterating posts"})
		return
	}

	var enrichedPosts []gin.H
	for _, post := range posts {
		commentCount, _ := commentCollection.CountDocuments(context.Background(), bson.M{"post_id": post.ID})
		likeCount, _ := likeCollection.CountDocuments(context.Background(), bson.M{"post_id": post.ID})

		enrichedPosts = append(enrichedPosts, gin.H{
			"post":         post,
			"commentCount": commentCount,
			"likeCount":    likeCount,
		})
	}

	c.JSON(http.StatusOK, enrichedPosts)
}
