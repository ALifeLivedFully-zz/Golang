/*package main

import (
	"io/ioutil"
	"fmt"
	"net/http"
	"encoding/xml"
)

type SitemapIndex struct {
	Locations []Location ;xml"sitemap"
}

type Location struct {
	Loc string ;xml 
}


func main()  {
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var s SitemapIndex
	xml.Unmarshal(bytes, &s)

	fmt.Println(s.Locations)
}
*/
package main

import (
	
)



func main()  {
	
}