package main

import "crypto/md5"

var username = "jake"
var password = "DaPassword"

func main() {
	h := md5.New() // this will also be ignored
	println(h)
	println(username, password)
}
