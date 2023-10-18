package main

import (
	"strings"
	"testing"
	"time"

	"github.com/gocolly/colly"
)



	
	func TestRua(t *testing.T) {
		rua := rua()
		if rua == "" {
			t.Error("Expected non-empty string, got empty string")
		}	
	
		expectedUserAgents := []string{
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
		found := false
		for _, userAgent := range expectedUserAgents {
			if strings.HasPrefix(rua, userAgent) {
				found = true
				break
			}
		}
	
		if !found {
			t.Errorf("Unexpected user agent: %s", rua)
		}
	}

	func testScrape(t *testing.T) {
		// Test case: Check if the scrape function visits the correct URL
		t.Run("Visit Correct URL", func(t *testing.T) {
			// Mock the colly collector
			mockCollector := &colly.Collector{}
	
			// Create a mock request
			mockRequest := &colly.Request{}
			mockCollector.OnResponse(func(r *colly.Response) {
				// Assert that the correct URL is visited
				expectedURL := "https://breachforums.is/"
				if r.Request.URL.String() != expectedURL {
					t.Errorf("Expected to visit URL %s, got: %s", expectedURL, r.Request.URL.String())
				}
			})
			mockCollector.OnRequest(func(r *colly.Request) {
				mockRequest.Visit(r.URL.String())
			})
	
			// Set the mock collector
			mockCollector.SetRequestTimeout(10 * time.Second)
			mockCollector.OnError(func(r *colly.Response, err error) {
				// Assert that the error is not nil
				if err != nil {
					t.Errorf("Unexpected error: %s", err)
				}
			})
	
			// Call the scrape function
			scrape()
		})
	}
