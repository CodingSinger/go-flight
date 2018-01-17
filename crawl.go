package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"

	"sync"

	"time"

	"strconv"
	"math/rand"
)

var areas = []string{"PEK","CAN","PVG","SHA","CTU","SZX","KMG","XIY","CKG","HGH","XMN","NKG","WUH","SHE","TAO","TSN"}
var pageurl string = "http://www.variflight.com/flight/"
var wg sync.WaitGroup  //定义一个同步等待的组

var client = MongoClient{MongoUrl:mongoUrl}
var r = rand.New(rand.NewSource(time.Now().UnixNano()))


func maina() {



	begin,_ := time.Parse("20060102","20180117")
	var flight Flight
	flight.Time = begin
	flight.Remainer = 100
	flight.SellTickets = 0
	flight.BookTickets = 0
	for index,area := range areas {
		for _,area1 := range areas[index+1:]{
			wg.Add(2)
			time.Sleep(10*1e9)
			 run(pageurl+area+"-"+area1+".html?AE71649A58c77&fdate=",flight,0,20180117)
			time.Sleep(10*1e9)
			 run(pageurl+area1+"-"+area+".html?AE71649A58c77&fdate=",flight,0,20180117)
		}
	}






	//run("http://www.variflight.com/flight/CAN-PEK.html?AE71649A58c77",flight,0)

	wg.Wait()
}

func run(url string,flight Flight,depth int,date int){

	if depth >= 3{
		return
	}

	time.Sleep(10*1e9)
	doc, err := goquery.NewDocument(url+ strconv.Itoa(date))


	if err != nil {
		log.Fatal(err)
	}

	 doc.Find("ul#list").Find("li").Each(func(k int, selection *goquery.Selection) {
		selection.Find("span").Each(func(i int, selection *goquery.Selection) {



			switch i {
			case 0:
				selection.Find("b").Find("a").Each(func(j int, selection *goquery.Selection) {

					if j == 1{
						flight.Fid = selection.Text()
					}

				})

				break
			case 3:

				temp := []rune(strFormat(selection.Text()))
				flight.Departure = string(temp[0:2])
				break
			case 6:
				temp := []rune(strFormat(selection.Text()))
				flight.Destination = string(temp[0:2])
				break

			}





		})

		flight.Price = float32(r.Intn(2000)+1000)



		client.insert(flight,database,flightCollection)
	})




	flight.Time,_ = time.Parse("20060102",strconv.Itoa(date+1))
	run(url,flight,depth+1,date+1)


}


func strFormat(str string)string{
	return strings.TrimSpace(strings.Replace(strings.Replace(str," ","",-1),"\n","",-1))
}