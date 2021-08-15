package main

import (
	"net/http"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
)

// getRestaurants responds with the list of all restaurant as JSON.
func getRestaurants(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, restaurants)
}

// postRestaurants adds a var waitingList = []WaitingList{} from JSON received in the request body.
func postRestaurants(c *gin.Context) {
	id := c.Param("id")
	var request WaitingListRequest
	var newWaitingList WaitingList

	// Call BindJSON to bind the received JSON to newRestaurant.
	if err := c.BindJSON(&request); err != nil {
		return
	}

	newWaitingListId := "1"
	nextCheckinNumber := 0
	if len(waitingList) > 0 {
		newWaitingListId = strconv.Itoa(len(waitingList) + 1)
		waitingListForRestaurant := filter(waitingList, func(waitingList WaitingList) bool { return waitingList.RestaurantId == id })
		nextCheckinNumber = len(waitingListForRestaurant) + 1
	}
	newWaitingList = WaitingList{
		ID: newWaitingListId,
		UserId: request.UserId,
		RestaurantId: id,
		Date: time.Now().String(),
		Number: nextCheckinNumber,
		WaitingAt: "",
		CheckinAt: "",
		CancelAt: "",
		FinishAt: "",
	}

	// Add the new restaurant to the slice.
	waitingList = append(waitingList, newWaitingList)
	c.IndentedJSON(http.StatusCreated, newWaitingList)
}

func filter(arr []WaitingList, predicate func(WaitingList) bool) []WaitingList {
    out := make([]WaitingList, 0)
 
    for _, e := range arr {
        if predicate(e) {
            out = append(out, e)
        }
    }
 
    return out
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
	router.POST("/restaurants/:id/waitingList", postRestaurants)

	router.Run("localhost:8080")
}
