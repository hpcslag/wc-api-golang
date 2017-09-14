package main

import (
	"net/url"
	"net/http"
	"fmt"
	"./lib"
)

type Client struct{
	domain string
	ck	string
	cs  string
	options *Options
	client *Client
}

func NewClient(domain, ck,cs string, options interface{}) *Client{
	c := &Client{
		domain,
		ck,
		cs,
		options,
		&http.Client,
	}
	return &c
}

func request(path string, params *url.Values) *http.Client{
	if(url.indexOf("https://") != -1){
		//make ssl protocol
	}else{
		//use oauth single request.
		util := lib.Utils{}
		signature, signature_params := util.Oauth_generator("ck","cs", "GET","http://192.168.1.140/wp-json/wc/v2/products")
		signature_params.Add("oauth_signature", signature)

		specUrl := path + "?" + url.QueryEscape(signature_params.Encode())
		
		//req := &http.Client{}
		//req.Do(specUrl)
	}
}

func (c *Client)Get(path string, params *url.Values) *http.Client{
	if params != nil{
		request(path, &params)
	}
}

func main(){

	b := lib.Utils{}
	
	s,_ := b.Oauth_generator("x","x", "GET","http://192.168.1.140/wp-json/wc/v2/products")

	fmt.Println(s)



	//simulation:
	c:= &NewClient()

	c.Get()
}

