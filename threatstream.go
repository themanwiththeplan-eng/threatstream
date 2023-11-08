package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
) // import "net/http"

type Threat struct{
	
}
func rua() string{
	rua := [10] string{
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36", 
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko", 
	"Mozilla/5.0 (Windows NT 5.1; rv:40.0) Gecko/20100101 Firefox/40.0", 
	"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36", 
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; rv:11.0) like Gecko", 
	"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:31.0) Gecko/20100101 Firefox/31.0", 
	"Mozilla/5.0 (Linux; U; Android 4.0.4; en-us; KFJWI Build/IMM76D) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.68 like Chrome/39.0.2171.93 Safari/537.36", 
	"Mozilla/5.0 (iPad; CPU OS 7_1 like Mac OS X) AppleWebKit/537.51.2 (KHTML, like Gecko) Version/7.0 Mobile/11D167 Safari/9537.53", 
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/600.1.25 (KHTML, like Gecko) Version/8.0 Safari/600.1.25", 
	"Mozilla/5.0 (Windows NT 6.3; Trident/7.0; Touch; rv:11.0) like Gecko",
	}
	return rua[rand.Intn(len(rua))]
}
func scrape(){
	fmt.Println(rua())
	userAgent := rua()
	c := colly.NewCollector(
	colly.UserAgent(userAgent),
	colly.AllowURLRevisit(),
	colly.AllowedDomains("breachforums.is"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement){
		link := e.Attr("href")
		if !strings.HasPrefix(link, "breachforums.is") {
			return
		}
		e.Request.Visit(e.Request.AbsoluteURL(link))
	})
	c.OnHTML("div[class=thread__icon ficon_3]", func(e *colly.HTMLElement){
		link := e.DOM.Find("a[href]").Text()
		if len(link) == 0{
			return
		}else{
			e.Request.Visit(e.Request.AbsoluteURL(link))
		}
})
	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting", r.URL.String())
		c.SetRequestTimeout(10 * time.Second)
	})
	c.OnError(func(r *colly.Response, err error){
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response){
		fmt.Println(r.StatusCode)
	})
	c.Visit("https://breachforums.is/")
	log.Println(c)
}
func sendEmail(){
	// TODO: Send an email

}
func main(){
	scrape()

	fName := "threatstream.csv"
	f, err := os.Create(fName)
	if err != nil {
		log.Fatal(err)
		return	
	}
	defer f.Close()
	writer := csv.NewWriter(f)
	defer writer.Flush()

	writer.Write([]string{"URL", "Threat"})
}