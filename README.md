# Description
A utility I use to generate a c language header ".h" from html. The utility strips html and javascript comments, replaces tabs, newlines and spaces with a single space.
# Installation
[install go lang here](https://golang.org/dl/)
set $GOPATH (eg */home/dave/go*) and add *$GOPATH/bin* to $PATH 
`go get -u github.com/centretown/htmltoc`
`cd $GOPATH/src/github.com/centretown/htmltoc`
`go install`
# Usage
 displays help
`htmltoc -h` 
 
reads tabs.html and creates/overwrites tabs.h
`htmltoc -i tabs`

reads tabs.html and creates/overwrites index.h
`htmltoc -i tabs -o index`

# Cautions
*"//"* comments should always be followed by a space eg // this is comment
javascript code should end statements with a "*;"*

