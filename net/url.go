package net

import (
	"fmt"
	"strings"
	"net/url"
)

func RequestURI(_url string) (string, error) {

	// if not start with http:// or https://, 
	if !strings.HasPrefix(_url,"http://")&&!strings.HasPrefix(_url,"https://"){
		return _url, fmt.Errorf("not a valid url")
	}
	
	parsed, err := url.Parse(_url)
	
	if err!=nil {
		return _url, err
	}else
	{
		avatarURL := parsed.RequestURI() 
		return avatarURL, nil
		//return avatarURL[1:len(avatarURL)], nil
	}

}