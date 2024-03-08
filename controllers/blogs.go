package controllers

import (
      //"fmt"
      "github.com/gin-gonic/gin"
      "net/http"
      "rest-api/models"
)

func FindBlogs(c *gin.Context){
    var blogs []models.Blog
    models.DB.Find(&blogs)
    c.JSON(http.StatusOK, gin.H{"data": blogs})
}

type createBlogSchema struct {
    Title string `json:"title" binding: "required"`
    Content string `json: "content" binding: "required"`
    Slug string `json: "slug" binding: "required"`
}
func CreateBlog(c *gin.Context){
    var inputBlog createBlogSchema;
    if err := c.ShouldBindJSON(&inputBlog); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    blog := models.Blog{Title: inputBlog.Title, Content: inputBlog.Content, Slug: inputBlog.Slug}
    models.DB.Create(&blog)
    c.JSON(http.StatusOK, gin.H{"data": blog})
}
func UpdateBlog(c *gin.Context){
    var blog models.Blog
    if err := models.DB.Where("id=?", c.Param("id")).First(&blog).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
        return
    }
    var inputBlog createBlogSchema
    if err:= c.ShouldBindJSON(&inputBlog); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return

    }
    models.DB.Model(&blog).Updates(&inputBlog)
    c.JSON(http.StatusOK, gin.H{"data": blog})
}
func DeleteBlog( c *gin.Context ){
    var blog models.Blog
    if err := models.DB.Where("id=?", c.Param("id")).First(&blog).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Not found!"})
        return
    }
    models.DB.Delete(&blog)
    c.JSON(http.StatusOK, gin.H{"status": true})
}
