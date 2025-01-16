package main

import "fmt"

func main() {
	var uname string = "admin"
	var passwd string  = "1234"

	fmt.Println("Authorization: Basic", uname + ":" + passwd)
}