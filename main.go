package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type restaurant struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	WaitingLimit   int    `json:"waitingLimit"`
	IsWaitlineOpen bool   `json:"isWaitlineOpen"`
	CheckinNumber  int    `json:"checkinNumber"`
	WaitingNumber  int    `json:"waitingNumber"`
}

var restaurants = []restaurant{
	{ID: "1", Name: "北車旗艦店", WaitingLimit: 500, IsWaitlineOpen: true, CheckinNumber: 0, WaitingNumber: 0},
	{ID: "2", Name: "永春店", WaitingLimit: 300, IsWaitlineOpen: true, CheckinNumber: 0, WaitingNumber: 0},
	{ID: "3", Name: "南港店", WaitingLimit: 250, IsWaitlineOpen: true, CheckinNumber: 0, WaitingNumber: 0},
}

// getRestaurants responds with the list of all restaurant as JSON.
func getRestaurants(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, restaurants)
}

// postRestaurants adds a restaurant from JSON received in the request body.
func postRestaurants(c *gin.Context) {
	var newRestaurant restaurant

	// Call BindJSON to bind the received JSON to newRestaurant.
	if err := c.BindJSON(&newRestaurant); err != nil {
		return
	}

	// Add the new restaurant to the slice.
	restaurants = append(restaurants, newRestaurant)
	c.IndentedJSON(http.StatusCreated, newRestaurant)
}

// getRestaurantByID locates the restaurant whose ID value matches the id
// parameter sent by the client, then returns that restaurant as a response.
func getRestaurantByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// a restaurant whose ID value matches the parameter.
	for _, a := range restaurants {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "restaurant not found"})
}

func main() {
	router := gin.Default()
	router.GET("/restaurants", getRestaurants)
	router.GET("/restaurants/:id", getRestaurantByID)
	router.POST("/restaurants", postRestaurants)

	router.Run("localhost:8080")
}
