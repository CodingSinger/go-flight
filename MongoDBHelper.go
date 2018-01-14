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


func (self MongoClient)updateTickets(flightId string,database string,collection string,opType int) int{

	self.session = self.initSession()
	c := self.session.DB(database).C(collection)

	var temp string
	if opType == 1{
		temp = "SellTickets"
	}else{

		temp = "BookTickets"
	}
	err := c.Update(bson.M{"Fid":flightId},bson.M{"$inc":bson.M{"Remainer":-1,temp:1}})

	if err != nil{
		return 1
	}
	return 0
}


func (self MongoClient) CloseClient(){
	self.session.Close()

}