package main

import "time"

type Response struct{

	State int
	Msg string
	Include interface{}
}


type Guest struct {

	username string

	password string
	lastlogin time.Time

	fullname string
	email string
}

type Passenger struct {
	Username string
	CardNum string
	FlightId string
	Time time.Time
	OpType int

}

type Flight struct {

	Departure string
	Destination string
	Time time.Time
	Remainer int
	SellTickets int
	BookTickets int
	Fid string
	Price float32

}