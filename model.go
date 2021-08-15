package main

type Restaurant struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	WaitingLimit   int    `json:"waitingLimit"`
	IsWaitlineOpen bool   `json:"isWaitlineOpen"`
	CheckinNumber  int    `json:"checkinNumber"`
	WaitingNumber  int    `json:"waitingNumber"`
}

type WaitingList struct {
	ID           string `json:"id"`
	UserId       string `json:"userId"`
	RestaurantId string `json:"restaurantId"`
	Date         string `json:"waitingNumber"`
	Number       int    `json:"number"`
	WaitingAt    string `json:"waitingAt"`
	CheckinAt    string `json:"checkinAt"`
	CancelAt     string `json:"cancelAt"`
	FinishAt     string `json:"finishAt"`
}

var restaurants = []Restaurant{
	{ID: "1", Name: "北車旗艦店", WaitingLimit: 500, IsWaitlineOpen: true, CheckinNumber: 0, WaitingNumber: 0},
	{ID: "2", Name: "永春店", WaitingLimit: 300, IsWaitlineOpen: true, CheckinNumber: 0, WaitingNumber: 0},
	{ID: "3", Name: "南港店", WaitingLimit: 250, IsWaitlineOpen: true, CheckinNumber: 0, WaitingNumber: 0},
}
