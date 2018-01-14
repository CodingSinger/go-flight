package main

import (
	"net/http"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"

	"time"
)
var mongoUrl string = "localhost:27017"
var database string = "test"
var flightCollection string = "flight"
var passengerCollection string = "passenger"
var errStatus Response = Response{
	State:1,
	Msg:"server error",
}

func bookTickets(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	params := r.Form["map"][0]
	var byt = make([]byte, 0)
	if params != "" {
		byt = []byte(params)
	}
	p := Passenger{}
	json.Unmarshal(byt,&params)
	t := time.Now().Format("2016-01-02")
	p.Time,_ = time.Parse("2016-01-02",t)

	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}
	defer mongo.CloseClient()
	state := mongo.insert(p,database,passengerCollection)
	mongo.updateTickets(p.FlightId,database,flightCollection,2)
	response := Response{
		State:state,
	}

	js,_ := json.Marshal(response)

	ResponseWithJSON(w,js,200)





}

func buyTickets(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	params := r.Form["passenger"][0]
	var byt = make([]byte, 0)
	if params != "" {
		byt = []byte(params)
	}
	p := Passenger{}
	json.Unmarshal(byt,&p)
	t := time.Now().Format("2006-01-02")
	p.Time,_ = time.Parse("2006-01-02",t)

	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}
	defer mongo.CloseClient()
	state := mongo.insert(p,database,passengerCollection)

	mongo.updateTickets(p.FlightId,database,flightCollection,1)
	response := Response{
		State:state,
	}

	js,_ := json.Marshal(response)

	ResponseWithJSON(w,js,200)





}
func queryData(w http.ResponseWriter,r *http.Request){
	r.ParseForm()


	queryStr := r.Form["queryStr"][0]

	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}
	//defer 	mongo.CloseClient()
	var f []Flight

	f = mongo.FindFlights(queryStr,database,flightCollection)

	response := Response{}




	response.Include = f

	js,err := json.Marshal(&response)

	if err != nil{
		log.Fatal(err)

	}


	ResponseWithJSON(w,js,200)



}

func login(w http.ResponseWriter,r *http.Request){



	r.ParseForm()

	username := r.Form["username"][0]
	password := r.Form["password"][0]




	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")

	state := Response{}
	result := Guest{}
	err = c.Find(bson.M{"name": username}).One(&result)

	if err != nil {

		if err.Error() == "not found"{
			state.State = 1
		}
	}else if result.password == password{
		state.State = 0

	}else{
		state.State = 1
	}

	s,err := json.Marshal(state)
	if err != nil {
		log.Fatal(err)
	}
	ResponseWithJSON(w,s,200)





}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func main(){

	http.HandleFunc("/login",login)

	http.HandleFunc("/fly/query",queryData)

	http.HandleFunc("/fly/buy",buyTickets)

	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("sas")
	}
}