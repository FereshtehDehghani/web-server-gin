package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64  `json:"price"`

}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context){
c.IndentedJSON(http.StatusOK,albums) 

}

func postAlbums(c *gin.Context){
    var newAlbum album

    if err:=c.BindJSON(&newAlbum);err != nil{
        return
    }

    albums=append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated,newAlbum)

}

func getAlbumById(c *gin.Context){
   id := c.Param("id");

  for _,value:= range albums {
       if value.ID == id {
 c.IndentedJSON(http.StatusFound,value)
       }
  }
   c.IndentedJSON(http.StatusNotFound,gin.H{"message":"Album not found"})
}

func main(){
// Create a Gin router with default middleware (logger and recovery)
 r := gin.Default()
  r.GET("/albums", getAlbums)
   r.POST("/albums", postAlbums)
   r.GET("/albums/:id",getAlbumById)

  // Start server on port 8080 (default)
  // Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
  r.Run("localhost:8080")

}