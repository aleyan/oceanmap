package main

import "testing"

func areSlicesDiff(left []string, right []string) bool {
    if len(left) != len(right) {
        return true
    }
    for i, v := range left {
        if v != right[i] {
            return true
        }
    }
    return false
}

func TestCanonizeUrl( t *testing.T){
	link := "www.digitalocean.com/hello"
	page := "http://www.digitalocean.com"
	url_out := "http://www.digitalocean.com/hello"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}	

	link = "http://www.digitalocean.com/hello"
	page = "http://www.digitalocean.com"
	url_out = "http://www.digitalocean.com/hello"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}		

	link = "https://www.digitalocean.com/hello"
	page = "http://www.digitalocean.com"
	url_out = "https://www.digitalocean.com/hello"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}

	link = "/hello"
	page = "http://www.digitalocean.com"
	url_out = "http://www.digitalocean.com/hello"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}

	link = "/hello"
	page = "http://www.digitalocean.com/"
	url_out = "http://www.digitalocean.com/hello"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}	

	link = "/hello/is/it/"
	page = "http://digitalocean.com/goodby/ciao"
	url_out = "http://digitalocean.com/hello/is/it/"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}

	link = "/company/blog/"
	page = "http://digitalocean.com/community/articles/digitalocean-community-article-suggestions-and-ideas"
	url_out = "http://digitalocean.com/company/blog/"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}	

	link = "hello/is/it/me/you/are/looking/for"
	page = "http://www.digitalocean.com/hello"
	url_out = "http://www.digitalocean.com/hello/is/it/me/you/are/looking/for"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}

	link = "hello/is/it/me/you/are/looking/for"
	page = "http://www.digitalocean.com/hello?goodbye"
	url_out = "http://www.digitalocean.com/hello/is/it/me/you/are/looking/for"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}

	link = "is?it"
	page = "http://www.digitalocean.com/hello"
	url_out = "http://www.digitalocean.com/is?it"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}

	link = "is?it=me&you=are"
	page = "http://www.digitalocean.com/hello"
	url_out = "http://www.digitalocean.com/is?it=me&you=are"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}	

	link = "www.digitalocean.com/hello#isit"
	page = "http://www.digitalocean.com/hello"
	url_out = "http://www.digitalocean.com/hello"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}	

	link = "#isit"
	page = "http://www.digitalocean.com/hello"
	url_out = "http://www.digitalocean.com/hello"
	if res :=CanonizeUrl(link, page); res != url_out {
		t.Errorf("%s (%s) expected to cannonize to %s, got %s", link, page, url_out, res)
	}	
}

func TestIsLegalUrl( t *testing.T){
	in := "www.digitalocean.com/hello"
	if IsLegalUrl(in) != false {
		t.Errorf("%s expected to be %v", in, false)
	}

	in = "https://www.digitalocean.com/hello"
	if IsLegalUrl(in) != true {
		t.Errorf("%s expected to be %v", in, true)
	}

	in = "http://www.digitalocean.com/hello"
	if IsLegalUrl(in) != true {
		t.Errorf("%s expected to be %v", in, true)
	}

	in = "https://digitalocean.com/hello"
	if IsLegalUrl(in) != true {
		t.Errorf("%s expected to be %v", in, true)
	}

	in = "http://digitalocean.com"
	if IsLegalUrl(in) != true {
		t.Errorf("%s expected to be %v", in, true)
	}

	in = "http://other.digitalocean.com/hello"
	if IsLegalUrl(in) != false {
		t.Errorf("%s expected to be %v", in, false)
	}	

	in = "https://other.digitalocean.com/hello"
	if IsLegalUrl(in) != false {
		t.Errorf("%s expected to be %v", in, false)
	}

	in = "http://www.digitalocean.com.sample.com/hello"
	if IsLegalUrl(in) != false {
		t.Errorf("%s expected to be %v", in, false)
	}			

	in = "other.digitalocean.com/hello"
	if IsLegalUrl(in) != false {
		t.Errorf("%s expected to be %v", in, false)
	}

	in = "www.digitalocean.other/hello"
	if IsLegalUrl(in) != false {
		t.Errorf("%s expected to be %v", in, false)
	}

	in = "http://www.digitalocean.other/hello"
	if IsLegalUrl(in) != false {
		t.Errorf("%s expected to be %v", in, false)
	}	

	in = "www.sample.com/hello"
	if IsLegalUrl(in) != false {
		t.Errorf("%s expected to be %v", in, false)
	}
}

func TestExtractTitle( t *testing.T){
	in := "<html><head></head><body></body></html>"
	out := ""
	if res := ExtractTitle(in); res != out {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}

	in = "<html><head><title>Example Domain</title></head><body></body></html>"
	out = "Example Domain"
	if res := ExtractTitle(in); res != out {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}	

}

func TestExtractUrls( t *testing.T){
	in := "<html><head></head><body></body></html>"
	out := []string{}
	if res :=ExtractUrls(in); areSlicesDiff(res, out) {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}

	in = "<html><head><link href='asset.css'/></head><body></body></html>"
	out = []string{}
	if res :=ExtractUrls(in); areSlicesDiff(res, out) {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}	

	in = "<html><head></head><body><a href='sample.com'></a></body></html>"
	out = []string{"sample.com"}
	if res :=ExtractUrls(in); areSlicesDiff(res, out) {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}

	in = "<html><head></head><body><A HREF='sample.com/capital'></A></body></html>"
	out = []string{"sample.com/capital"}
	if res :=ExtractUrls(in); areSlicesDiff(res, out) {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}	

	in = "<html><head></head><body><a href='/relative.html'></a></body></html>"
	out = []string{"/relative.html"}
	if res :=ExtractUrls(in); areSlicesDiff(res, out) {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}		

	in = "<html><head></head><body><div><p>text<a href='sample.com'></a>text</p></div></body></html>"
	out = []string{"sample.com"}
	if res :=ExtractUrls(in); areSlicesDiff(res, out) {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}

	in = "<html><head></head><body><div><p>text<a href='sample.com'>a</a>text<a href='resample.com'>b</a></p></div></body></html>"
	out = []string{"sample.com", "resample.com"}
	if res :=ExtractUrls(in); areSlicesDiff(res, out) {
		t.Errorf("%s expected to be %s, got %s", in, out, res)
	}		
}

func TestGetPage( t *testing.T){
	page := Page{Url: "http://example.com"}
	html := GetPage(page)
	if html == "" {
		t.Errorf("Expected to load %s, didn't get anything back.", page.Url)
	}

	if len(html) < 60 {
		t.Errorf("Expected to load atleast 60 characters, got.", len(html))
	}	
}