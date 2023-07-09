package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type album struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Artist string `json:"artist"`
  Price float64 `json:"price"`
}

var albums = []album{
  {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
  {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
  {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
  router := gin.Default()
  router.GET("/albums", get_albums)
  router.POST("/albums", post_albums)
  router.GET("/albums/:id", get_album_by_id)
  router.Run("localhost:8080")
}

func get_albums(c *gin.Context){
  c.IndentedJSON(http.StatusOK, albums)
}


func post_albums(c *gin.Context) {
  var newAlbum album
  if err := c.BindJSON(&newAlbum); err != nil {
    return
  }
  albums = append(albums, newAlbum)
  c.IndentedJSON(http.StatusCreated, newAlbum)
}

func get_album_by_id(c *gin.Context) {
  id := c.Param("id")
  for _, a := range albums {
    if a.ID == id {
      c.IndentedJSON(http.StatusOK, a)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
