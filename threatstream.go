package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
) // import "net/http"

func scrape(){
	c := colly.NewCollector(
	colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.142.86 Safari/537.36"),
	colly.AllowURLRevisit(),
	colly.AllowedDomains("https://www.breachforums.is", "https://www.exposed.vc"),
	)
	err := c.Post("https://breachforums.is/member?action=login", map[string]string{
		"username": "username",
		"password": "password",
	})
	if err != nil {
		log.Fatal(err)
	}

	c.OnResponse(func(r *colly.Response){
		fmt.Println(r.StatusCode)
	})

	c.OnHTML("*", func(e *colly.HTMLElement){
		fmt.Println(e.Text)
		
	})
	c.OnRequest(func(r *colly.Request){
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnError(func(r *colly.Response, err error){
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	c.Visit("https://www.breachforums.is")
	c.Visit("https://www.exposed.vc")
}

func main(){
scrape()
}