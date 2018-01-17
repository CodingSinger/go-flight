package main

import (
	"fmt"


	"time"
	"testing"

	"strconv"
)


func maisn() {

	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}
	queryStr := "{\"destination\":\"北京\",\"time\":\"2018-01-13\"}";
	var f []Flight
	f = mongo.FindFlights(queryStr,database,flightCollection)
	for _,e := range f {
		fmt.Println(e)
	}
}

func Test2(t *testing.T) {



	fmt.Println( strconv.Itoa(34355100))



	s := "按时函数"

	fmt.Print(s[0:1])

	s3 := []rune(s)
	fmt.Println(string(s3[0:1]))

	s2 := "sadsds"
	fmt.Println(s2[0:3])
}

func Test_findByPk(ts *testing.T){

	mongo := MongoClient{
		MongoUrl:mongoUrl,
	}
	t ,err:= time.Parse("2006-01-02","2018-01-14")
	if err != nil{
		panic(err)
	}
	a := Flight{
		Destination:"北京",
		Departure:"杭州",
		Remainer:10,
		SellTickets:5,
		BookTickets:3,
		Price:300,
		Fid:"fdf",

		Time:t,
	}

	mongo.insert(a,database,flightCollection)
}

func Test_def(ts *testing.T){
	var temp int
	defer func(){
		if temp == 1{
			fmt.Printf("1")
		}else{
			fmt.Println("2")
		}
	}()
	defer fmt.Printf("sdd")

	defer fmt.Println("b")

	temp = 2
	temp = 1

}

func Test_time(t *testing.T) {

	timestamp := time.Now().Unix()

	fmt.Println(timestamp)



	//格式化为字符串,tm为Time类型



	fmt.Println(time.Now().Format("2006-01-02"))

}

