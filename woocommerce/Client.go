package main

import (
	"fmt"
	"./lib"
)

func main(){

	b := lib.Utils{}
	
	s,_ := b.Oauth_generator("x","x", "GET","http://192.168.1.140/wp-json/wc/v2/products")

	fmt.Println(s)
}

