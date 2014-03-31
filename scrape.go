package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

//Depending on you install, you may have to get
//these dependencies manually.
//go get code.google.com/p/go-html-transform
//go get code.google.com/p/go.net
import (
	sel "code.google.com/p/go-html-transform/css/selector"
	h5 "code.google.com/p/go-html-transform/h5"
)

type Page struct {
	Title   string
	Url     string
	Visited bool
	Links   []*Page
}

var pages map[string]*Page
var unvisited []*Page

func init() {
	pages = make(map[string]*Page)
}

func CanonizeUrl(link, page string) string {
	//Constructs a fully qualified path with protocol out of a possible
	//relative path and the current page
	url := ""
	if m, _ := regexp.MatchString("^http(s)?://.*", link); m {
		//fully qualified, just need to drop the #part
		url = link
	} else if m, _ := regexp.MatchString("^[a-zA-Z0-9]+\\.[a-zA-Z0-9\\.]+", link); m {
		//has a domain, just need to add the protocol.
		if strings.HasPrefix(page, "https") {
			url = "https://" + link
		} else {
			url = "http://" + link
		}
	} else if strings.HasPrefix(link, "#") {
		//internal link.
		url = page
	} else {
		//relative path. use everything before the last slash from the page
		if strings.HasPrefix(link, "/") {
			re := regexp.MustCompile("/[a-zA-Z0-9#\\?/&=\\-_]*$")
			url = re.ReplaceAllString(page, "") + link
		} else {
			re := regexp.MustCompile("/[a-zA-Z0-9#\\?&=\\-_]*$")
			url = re.ReplaceAllString(page, "") + "/" + link
		}
	}
	re := regexp.MustCompile("#[a-zA-Z0-9]*$")
	url = re.ReplaceAllString(url, "")
	return url
}

func IsLegalUrl(url string) bool { //Works only on canonized URLs
	if m, _ := regexp.MatchString("^http(s)?://(www\\.)?digitalocean\\.com(/.*)?$", url); m {
		return true
	} else {
		return false
	}
}

func ExtractTitle(html string) string {
	title := ""

	// Get a parse tree for this HTML
	h5tree, err := h5.NewFromString(html)
	if err != nil {
		return title
	}
	n := h5tree.Top()

	// Create a Chain object from a CSS selector statement
	chn, err := sel.Selector("title")
	if err != nil {
		return title
	}

	// Find the Title element and read it's text
	if len(chn.Find(n)) > 0 {
		title = chn.Find(n)[0].FirstChild.Data
	}

	return title
}

func ExtractUrls(html string) []string {
	urls := []string{}

	// Get a parse tree for this HTML
	h5tree, err := h5.NewFromString(html)
	if err != nil {
		return urls
	}
	n := h5tree.Top()

	// Create a Chain object from a CSS selector statement
	chn, err := sel.Selector("a")
	if err != nil {
		return urls
	}

	// Find the anchor and read it's href attribute
	as := chn.Find(n)
	for _, a := range as {
		for _, attr := range a.Attr {
			if attr.Key == "href" {
				urls = append(urls, attr.Val)
			}
		}
	}

	return urls
}

func GetPage(page Page) string {
	//Goes to the interwebz to pull down the HTML of a given page object
	resp, err := http.Get(page.Url)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(body)
}

func appendIfMissing(slice []string, s string) []string {
	for _, ele := range slice {
		if ele == s {
			return slice
		}
	}
	return append(slice, s)
}

func TraverseGraph() {
	for i := 0; i < 1000 && len(unvisited) > 0; i++ {
		TraversePage(unvisited[0])
		unvisited = unvisited[1:len(unvisited)]
		if (i)%25 <= 1 {
			fmt.Print(".")
		}
	}
}

func TraversePage(page *Page) {
	//Visits a page and adds any links found to it's Links attribute
	//and adds any previously unknown pages to unvisited slice.
	if page.Visited {
		return
	}

	html := GetPage(*page)
	page.Title = ExtractTitle(html)
	links := ExtractUrls(html)
	urls := []string{}

	for _, link := range links {
		url := CanonizeUrl(link, page.Url)
		if IsLegalUrl(url) && url != page.Url {
			urls = appendIfMissing(urls, url)
		}
	}

	for _, url := range urls {
		linked, ok := pages[url]
		if !ok {
			linked = new(Page)
			linked.Url = url
			unvisited = append(unvisited, linked)
			pages[url] = linked
		}
		page.Links = append(page.Links, linked)
	}
	page.Visited = true
}
func PrintPageLinks(page *Page) {
	for _, link := range page.Links {
		fmt.Println(link.Url)
	}
}

func PrintGraph() {
	re := regexp.MustCompile(".com/[a-zA-Z0-9\\-_]+")
	contents := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>" +
		"<graphml xmlns=\"http://graphml.graphdrawing.org/xmlns\">" +
		"<key id=\"title\" for=\"node\" attr.name=\"title\" attr.type=\"string\"/>" +
		"<key id=\"category\" for=\"node\" attr.name=\"category\" attr.type=\"string\"/>" +
		"<graph id=\"G\" edgedefault=\"directed\">\n"
	for _, page := range pages {
		if page.Visited {
			categories := re.FindStringSubmatch(page.Url)
			category := ""
			if len(categories) > 0 && len(categories[0]) > 5 {
				category = categories[0][5:]
			}

			contents = contents + fmt.Sprintf("<node id=\"%s\"><data key=\"title\">%s</data><data key=\"category\">%s</data></node>\n",
				page.Url, page.Title, category)
		}
	}
	edge := 0
	for _, page := range pages {
		if !page.Visited {
			continue
		}
		for _, link := range page.Links {
			if !link.Visited {
				continue
			}
			edge++
			contents = contents + fmt.Sprintf("<edge id=\"e%v\" source=\"%s\" target=\"%s\"/>\n",
				edge,
				page.Url,
				link.Url)
		}
	}
	contents = contents + "</graph></graphml>"
	ioutil.WriteFile("digitalocean.graphml", []byte(contents), 0x777)
}
