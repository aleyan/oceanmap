#OceanMap

This program scans 1000 internaly linked pages on digitalocean.com in a breadth first search and stores them in digitalocean.graphml. This file can be viewed with open source graph software such as Gephi.

Additianlly a URL can be passed as an argument to the program. In this case the program scans just that page and prints the absolute paths of links to standard out.

##Requirements
This software was written with go 1.2.1 in mind. It needs Mercurial SCM installed to pull in dependencies from code.google.com

##How to build
Run the following command:

    go build oceanmap.go scrape go
Alternativey you can run this program directly:

    go run oceanmap.go scrape go

##To Do
* Add an option to track static assets used by a page.
* The maximum number of pages to scan should be a parameter.
* Instead of manually translating relative links from current page, use Go URL library.
* Use a string buffer or an XML builder when outputting the final graphml diagram.
* Have an argument for ignoring links in footers.
