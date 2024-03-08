package main;
import (
  "github.com/gin-gonic/gin"
  "rest-api/models"
  "rest-api/controllers"
)

func main(){
    r := gin.Default()
    models.ConnectDatabase()
    r.GET("/blogs", controllers.FindBlogs)
    r.POST("/blog", controllers.CreateBlog)
    r.PATCH("/blog/:id", controllers.UpdateBlog)
    r.DELETE("/blog/:id", controllers.DeleteBlog)
    r.Run()
}
