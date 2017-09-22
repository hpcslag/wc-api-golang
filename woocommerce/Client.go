package woocommerce

import (
	"net/url"
	"net/http"
	"fmt"
	"woocommerce/lib"
	"strings"
	"bytes"
	"encoding/json"
	"io"
)

type Client struct{
	domain *url.URL
	ck	string
	cs  string
	client *http.Client
}

func NewClient(domainURL, ck,cs string) *Client{
	
	domain, err := url.Parse(domainURL)
	if err != nil{
		panic(err)
	}

	domain.Path = "/wp-json/wc/v2/"
	
	c := &Client{
		domain,
		ck,
		cs,
		&http.Client{},
	}
	return c
}

func (C *Client)request(method, path string, params *url.Values,data interface{}) io.ReadCloser{
	
	util := lib.Utils{}
	signature, signature_params := util.Oauth_generator(C.ck, C.cs, strings.ToUpper(method),(C.domain.String() + path))
	signature_params.Add("oauth_signature", signature)

	//Arrangement of url parameter(specUrl) and Oauth(oauth_generator) parameter will affect result of Signature
	specUrl := C.domain.String() + path + "?" + "oauth_consumer_key=" + signature_params.Get("oauth_consumer_key") + "&oauth_nonce=" + signature_params.Get("oauth_nonce") + "&oauth_signature=" + url.QueryEscape(signature) + "&oauth_signature_method=" + signature_params.Get("oauth_signature_method") + "&oauth_timestamp=" + signature_params.Get("oauth_timestamp")
	
	fmt.Println(specUrl)


	body := new(bytes.Buffer)
	encoder := json.NewEncoder(body)
	if err := encoder.Encode(data); err != nil {
		return nil
	}

	req, _ := http.NewRequest(strings.ToUpper(method), specUrl, body)
	req.Header.Set("Content-Type", "application/json")
	
	res, err := C.client.Do(req)
	if err != nil{
		return nil
	}
	
	return res.Body
}

func (c *Client)Get(path string, params *url.Values) io.ReadCloser{
	if params != nil{
		return c.request("GET", path, params,nil)
	}
	return c.request("GET", path, nil,nil)
}

func (c *Client)Post(path string, data interface{}) io.ReadCloser{
	return c.request("POST", path, nil, data)
}

func (c *Client)Put(path string, data interface{}) io.ReadCloser{
	return c.request("POST", path, nil, data)
}

func (c *Client)Delete(path string, params *url.Values) io.ReadCloser{
	if params != nil{
		return c.request("DELETE", path, params,nil)
	}
	return c.request("DELETE", path, nil,nil)
}

func (c *Client)Patch(path string, params *url.Values) io.ReadCloser{
	if params != nil{
		return c.request("PATCH", path, params,nil)
	}
	return c.request("PATCH", path, nil,nil)
}

func (c *Client)Option(path string, params *url.Values) io.ReadCloser{
	if params != nil{
		return c.request("OPTIONS", path, params,nil)
	}
	return c.request("OPTIONS", path, nil,nil)
}
