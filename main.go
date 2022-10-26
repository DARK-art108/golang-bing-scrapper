package main

import(
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"math/rand"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

bingDomains = map[string]string{
	"com": "https://www.bing.com/search?q=",

}

type searchResult struct {
	ResultRank int
	ResultURL  string
	ResultTitle string
	ResultDescription string
}

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36"
}

func randomAgent() string{

}

func buildBingURLs{

}

func scrapeClientRequest{

}


func bingScrape{

}

func bingResultParser{
	
}

func main() {
}
