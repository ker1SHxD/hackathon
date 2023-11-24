package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"slices"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Fruit struct {
	Name string
}

var users = []user{
	{Username: "ker1SH", Password: "bababoi"},
	{Username: "rafaellmir", Password: "sussybaka"},
	{Username: "vladdyboy", Password: "mission13"},
}

func postUser(c *gin.Context) {
	var newUser user
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	if slices.Contains(users, newUser) {
		c.IndentedJSON(http.StatusOK, newUser)
	} else {
		users = append(users, newUser)
		c.IndentedJSON(http.StatusCreated, newUser)
	}
}csd

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.POST("/login", postUser)
	router.Run("192.168.43.2:8080")
}
