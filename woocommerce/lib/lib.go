package lib

import (
	"net/url"
)


type Utils struct{
}

func (l Utils)Oauth_generator(_KEY, _SECRET, HttpMethod, RequestUrl string) (string, *url.Values){
	return oauth_generator(_KEY, _SECRET, HttpMethod, RequestUrl)
}

func (l Utils)Cool()string{
	return getSha1Nonce()
}