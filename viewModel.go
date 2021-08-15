package viewModel

type RestaurantViewModel struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
}

type WaitingListViewModel struct {
	ID           string `json:"waitingListId"`
	Number       int    `json:"number"`
}

type RestaurantDetailViewModel struct {
	Name              string `json:"name"`
	IsWaitlineOpen    bool   `json:"isWaitlineOpen"`
	WaitingLimit      int    `json:"waitingLimit"`
	WaitingCount      int    `json:"waitingCount"`
	CheckinCount      int    `json:"checkinCount"`
	CheckinNumber     int    `json:"checkinNumber"`
	NextCheckinNumber int    `json:"checkinNumber"`
}

