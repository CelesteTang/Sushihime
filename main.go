package main

import (
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// getRestaurants responds with the list of all restaurant as JSON.
func getRestaurants(c *gin.Context) {

	var vm RestaurantViewModel
	for _, r := range restaurants {
		vm = append(vm, RestaurantViewModel{ID: r.ID, name: r.name} )
	}
	c.IndentedJSON(http.StatusOK, vm)
}

// postWaitingList adds a var waitingList = []WaitingList{} from JSON received in the request body.
func postWaitingList(c *gin.Context) {
	id := c.Param("id")
	var request WaitingListRequest
	var newWaitingList WaitingList

	// Call BindJSON to bind the received JSON to newRestaurant.
	if err := c.BindJSON(&request); err != nil {
		return
	}

	newWaitingListId := "1"
	nextCheckinNumber := 0
	if len(waitingLists) > 0 {
		newWaitingListId = strconv.Itoa(len(waitingLists) + 1)
		waitingListForRestaurant := filter(waitingLists, func(waitingList WaitingList) bool { return waitingList.RestaurantId == id })
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
	waitingLists = append(waitingLists, newWaitingList)
	c.IndentedJSON(http.StatusCreated, WaitingListViewModel{ WaitingListId: newWaitingListId, Number: nextCheckinNumber})
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
	
	var restaurant Restaurant
	var currentWaitingList []WaitingList
	var currentCheckingList []WaitingList

	for _, r := range restaurants {
		if r.ID == id {
			restaurant = r
			for _, w := range waitingLists {
				if w.RestaurantId == r.ID {
					if w.CheckinAt == "" && w.CancelAt == "" {
						currentWaitingList = append(currentWaitingList, w)
					} else if w.CheckinAt != "" && w.FinishAt == "" {
						currentCheckingList = append(currentCheckingList, w)
					}
				}
			}
			break
		}
	}
	sort.SliceStable(currentCheckingList, func(i, j int) bool {
		return currentWaitingList[i].Number < currentWaitingList[j].Number
	})

	nextCheckinNumber := 0
	if len(currentCheckingList) > 0 {
		nextCheckinNumber = currentCheckingList[0].Number
	}

	vm := RestaurantDetailViewModel {
		Name: restaurant.Name,
		IsWaitlineOpen: restaurant.IsWaitlineOpen,
		WaitingLimit: restaurant.WaitingLimit,
		WaitingCount: len(currentWaitingList),
		CheckinCount: len(currentCheckingList),
		CheckinNumber: restaurant.CheckinNumber,
		NextCheckinNumber: nextCheckinNumber,
	}
	
	c.IndentedJSON(http.StatusOK, vm)
}

func main() {
	router := gin.Default()
	router.GET("/restaurants", getRestaurants)
	router.GET("/restaurants/:id", getRestaurantByID)
	router.POST("/restaurants/:id/waitingList", postWaitingList)

	router.Run("localhost:8080")
}
