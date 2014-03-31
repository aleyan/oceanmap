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
* The maximum number of pages to scan should be a parameter.
* Add an option to capture links to resources in addition to other pages.
* Instead of manually translating relative links from current page, use go URL library.
* Use a string buffer when outputting the final graphml diagram.
