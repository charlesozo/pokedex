package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationArea, error) {
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint
	if pageUrl != nil{
		fullURL = *pageUrl
	}
	data, ok := c.cache.Get(fullURL)
	if ok{
	// cache hit
	fmt.Println("cache hit")	
	locationAreaResp := LocationArea{}
	err := json.Unmarshal(data, &locationAreaResp)
	if err != nil{
		return LocationArea{}, err
	}
	return locationAreaResp, nil
}
fmt.Println("cache miss")
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{} , err
	}

  
	resp, err := c.httpClient.Do(req)
    if err != nil {
		return LocationArea{} , err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return LocationArea{} , fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err = io.ReadAll(resp.Body)
	if err !=nil{
		return LocationArea{}, err
	}
	
	locationAreaResp := LocationArea{}
	err = json.Unmarshal(data, &locationAreaResp)
	if err != nil{
		return LocationArea{}, err
	}
	c.cache.Add(fullURL, data)
	return locationAreaResp, nil
}