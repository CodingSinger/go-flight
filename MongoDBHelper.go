package main

import (
	"gopkg.in/mgo.v2"


	"encoding/json"

	"time"

	"gopkg.in/mgo.v2/bson"

)



type MongoClient struct{
	MongoUrl string
	session  *mgo.Session

}

func (self MongoClient) initSession() *mgo.Session{

	if self.session == nil{
		session, err := mgo.Dial(mongoUrl)
		session.SetMode(mgo.Monotonic, true)
		return session.Clone()

		if err != nil {
			panic(err)
		}
	}
	return self.session

}
func (self MongoClient) insert(obj interface{},database string,collection string) int{
	self.session = self.initSession()
	c := self.session.DB(database).C(collection)
	err := c.Insert(&obj)
	if err != nil{
		return 1
	}
	return 0

}


func (self MongoClient) FindFlightById(fid string,database string,collection string)([]Flight,error){
	self.session = self.initSession()

	c := self.session.DB(database).C(collection)

	var result []Flight

	err := c.Find(bson.M{"fid":fid}).All(&result)

	return result,err




}
func (self MongoClient) FindFlights(queryStr string,database string,collection string) []Flight{


	self.session = self.initSession()


	// Optional. Switch the session to a monotonic behavior.


	c := self.session.DB(database).C(collection)

	var condition map[string]interface{}
	var byt = make([]byte, 0)
	if queryStr != "" {
		byt = []byte(queryStr)
	}

	var result []Flight

	json.Unmarshal(byt, &condition)


	if condition["time"] != nil{
		s := condition["time"].(string);
		t ,err:= time.Parse("2006-01-02",s)

		if err != nil{
			panic(err)
		}

		condition["time"] = t
	}
	err := c.Find(condition).All(&result)


	if err != nil {

		if err.Error() == "not found"{
			return result
		}



		}
	return	result

}


func (self MongoClient)updateTickets(flightId string,time time.Time,database string,collection string,opType int) error {

	self.session = self.initSession()
	c := self.session.DB(database).C(collection)

	var temp string
	if opType == 1 {
		temp = "selltickets"
	} else {

		temp = "booktickets"
	}
	err := c.Update(bson.M{"fid": flightId,"time":time}, bson.M{"$inc": bson.M{"remainer": -1, temp: 1}})
	return err
}
func (self MongoClient) queryPassengers(flightId string,database string,collection string)([]Passenger,error){
	self.session = self.initSession()


	// Optional. Switch the session to a monotonic behavior.


	c := self.session.DB(database).C(collection)

	var passengers []Passenger

	err := c.Find(bson.M{"flightid":flightId}).All(&passengers)

	return passengers,err




}


func (self MongoClient) queryGuest(username string,database string,collection string) (Guest,error){

	self.session = self.initSession()


	// Optional. Switch the session to a monotonic behavior.


	c := self.session.DB(database).C(collection)


	var r Guest
	err := c.Find(bson.M{"username":username}).One(&r)

	return r,err
}

func (self MongoClient) CloseClient(){
	self.session.Close()

}