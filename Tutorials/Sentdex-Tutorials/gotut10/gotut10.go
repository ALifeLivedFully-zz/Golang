package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
)

func main()  {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringbody := string(bytes)
	fmt.Println(stringbody)
	resp.Body.Close()
}

/*
this was a goddamn nightmare, dont forget to use the full url eg. include https://www.
*/