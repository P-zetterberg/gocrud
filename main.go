package main

import (
	"go-crud/data"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
	Id   int    `json:"id"`
}

var users []User

func handleDeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	for index, user := range users {
		if user.Id == id {
			// Remove the user from the slice
			users = append(users[:index], users[index+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
func handleUpdateUser(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i := 0; i < len(users); i++ {
		if users[i].Id == user.Id {
			users[i].Name = user.Name
			c.JSON(http.StatusOK, users[i])
		}
	}
}

func main() {
	users = append(users, User{Name: "Pontus", Age: 29, City: "Västerås", Id: 1})
	users = append(users, User{Name: "Poze", Age: 1337, City: "OSRS", Id: 2})
	r := gin.Default()
	r.GET("/getUsers", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, users)
	})
	r.GET("/getUser/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		for i := 0; i < len(users); i++ {
			if users[i].Id == id {
				c.JSON(http.StatusOK, users[i])
			}
		}
	})

	r.POST("/addUser", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		users = append(users, user)
	})

	r.DELETE("/deleteUser/:id", handleDeleteUser)
	r.PUT("/updateUser", handleUpdateUser)
	data.InitDatabase()
	r.Run()
}
