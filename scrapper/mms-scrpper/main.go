package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

var baseURL = "https://www.mymusicsheet.com/search?page="
var mmsPage = 1000

func main() {
	getPage(1)
}

func getPage(page int) {
	pageURL := baseURL + strconv.Itoa(page)
	fmt.Println("Requesting" + pageURL)
	resp, err := http.Get(pageURL)
	checkErr(err)
	checkCode(resp)

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	checkErr(err)
	fmt.Println(doc)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
}
