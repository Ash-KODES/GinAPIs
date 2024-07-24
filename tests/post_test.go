package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "socialapp/controllers"
    "socialapp/models"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreatePost(t *testing.T) {
    router := gin.Default()
    router.POST("/posts", controllers.CreatePost)

    userID, err := primitive.ObjectIDFromHex("60d5ec49d6db2c3f387f10c9")
    assert.Nil(t, err)

    post := models.Post{
        Name:        "testing Post",
        Description: "this is a test post",
        UserID:      userID,
    }
    jsonValue, _ := json.Marshal(post)
    req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    // Record the response
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assertions
    assert.Equal(t, http.StatusCreated, w.Code)
    var responsePost models.Post
    err = json.Unmarshal(w.Body.Bytes(), &responsePost)
    assert.Nil(t, err)
    assert.Equal(t, post.Name, responsePost.Name)
    assert.Equal(t, post.Description, responsePost.Description)
    assert.NotNil(t, responsePost.ID)
}

func TestGetPost(t *testing.T) {
    // Initialize Gin router
    router := gin.Default()
    router.GET("/posts/:id", controllers.GetPost)

    // Create a new request to get a post
    req, _ := http.NewRequest("GET", "/posts/60d5ec49d6db2c3f387f10c9", nil)

    // Record the response
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // Assertions
    assert.Equal(t, http.StatusOK, w.Code)
    var responsePost models.Post
    err := json.Unmarshal(w.Body.Bytes(), &responsePost)
    assert.Nil(t, err)
    assert.Equal(t, "Test Post", responsePost.Name)
}

func TestUpdatePost(t *testing.T) {

    router := gin.Default()
    router.PUT("/posts/:id", controllers.UpdatePost)

    // string ID to objectID
    userID, err := primitive.ObjectIDFromHex("60d5ec49d6db2c3f387f10c9")
    assert.Nil(t, err)


    updatedPost := models.Post{
        Name:        "Updated Post",
        Description: "This is an updated post",
        UserID:      userID,
    }
    jsonValue, _ := json.Marshal(updatedPost)
    req, _ := http.NewRequest("PUT", "/posts/60d5ec49d6db2c3f387f10c9", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    // response store
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    // checks
    assert.Equal(t, http.StatusOK, w.Code)
    var responsePost models.Post
    err = json.Unmarshal(w.Body.Bytes(), &responsePost)
    assert.Nil(t, err)
    assert.Equal(t, updatedPost.Name, responsePost.Name)
    assert.Equal(t, updatedPost.Description, responsePost.Description)
}
