package main

import (
	"net/http"
	"log"

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

	params := r.Form["passenger"][0]

	t2 := r.Form["time"][0]
	var byt = make([]byte, 0)
	if params != "" {
		byt = []byte(params)
	}
	p := Passenger{}
	json.Unmarshal(byt,&p)
	t := time.Now().Format("2016-01-02")
	p.Time,_ = time.Parse("2016-01-02",t)

	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}

	state := mongo.insert(p,database,passengerCollection)
	response := Response{
		State:state,
	}

	flytime,_ := time.Parse("2006-01-02",t2)
	err := mongo.updateTickets(p.FlightId,flytime,database,flightCollection,2)

	if err != nil{
		log.Fatal(nil)
		response.State = 1
	}


	js,_ := json.Marshal(response)

	ResponseWithJSON(w,js,200)






}

func buyTickets(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	params := r.Form["passenger"][0]
	t2 := r.Form["time"][0]



	var byt = make([]byte, 0)
	if params != "" {
		byt = []byte(params)
	}
	p := Passenger{}
	json.Unmarshal(byt,&p)
	//预定时间
	t := time.Now().Format("2006-01-02")
	p.Time,_ = time.Parse("2006-01-02",t)

	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}

	state := mongo.insert(p,database,passengerCollection)

	flytime,_ := time.Parse("2006-01-02",t2)
	mongo.updateTickets(p.FlightId,flytime,database,flightCollection,1)
	response := Response{
		State:state,
	}

	js,_ := json.Marshal(response)

	ResponseWithJSON(w,js,200)





}



func queryData(w http.ResponseWriter,r *http.Request){
	r.ParseForm()




	queryType := r.Form["type"]





	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}


	var f []Flight

	if len(queryType) != 0 &&queryType[0]== "num"  {

		//按航班号查询
		fid := r.Form["fid"][0]


		var erro error
		f,erro = mongo.FindFlightById(fid,database,flightCollection)

		if erro != nil{
			log.Fatal(erro)

		}





	}else{
		queryStr := r.Form["queryStr"][0]
		f = mongo.FindFlights(queryStr,database,flightCollection)
	}

	response := Response{}




	response.Include = f

	js,err := json.Marshal(&response)

	if err != nil{
		log.Fatal(err)

	}


	ResponseWithJSON(w,js,200)



}


func queryPassengers(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fid := r.Form["fid"][0]
	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}

	state := Response{
		State:0,
	}
	result,err := mongo.queryPassengers(fid,database,passengerCollection)


	state.Include = result
	if err != nil{
		log.Fatal(err)
		state.State = 1
	}




	s,err := json.Marshal(state)
	if err != nil {
		log.Fatal(err)
		state.State = 1
	}
	ResponseWithJSON(w,s,200)




}

var guestCollection string = "admin"

func login(w http.ResponseWriter,r *http.Request){



	r.ParseForm()

	username := r.Form["username"][0]
	password := r.Form["password"][0]





	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}
	state := Response{}
	result,err := mongo.queryGuest(username,database,guestCollection)


	if err != nil {

		if err.Error() == "not found"{
			state.State = 1
		}
	}else if result.password == password{
		state.State = 0

	}else{
		state.State = 2
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

	http.HandleFunc("/fly/book",bookTickets)
	http.HandleFunc("/fly/queryPassengers",queryPassengers)

	err := http.ListenAndServe(":9090",nil)
	if err != nil{
		log.Fatal("sas")
	}
}