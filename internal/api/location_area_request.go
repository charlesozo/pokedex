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
	locationAreaResp := LocationArea{}
	err := json.Unmarshal(data, &locationAreaResp)
	if err != nil{
		return LocationArea{}, err
	}
	return locationAreaResp, nil
}

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

func (c *Client) Xplore(input string) (XploreArea, error){
	endpoint := "/location-area/" + input
	fullURL := baseUrl + endpoint

	req, err := http.NewRequest("GET", fullURL, nil)
	if err!=nil{
        return XploreArea{}, nil
	}
    resp, err := c.httpClient.Do(req)
	if err != nil{
		return XploreArea{}, nil
	}
	defer resp.Body.Close()
    if resp.StatusCode > 399 {
		return XploreArea{} , fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err!=nil{
		return XploreArea{}, nil
	}
	xploreAreaResp := XploreArea{}
	err = json.Unmarshal(data, &xploreAreaResp)
	if err!=nil{
		return XploreArea{}, err
	}
	
	return xploreAreaResp, nil

}
func (c * Client) PokeData(input string) (PokemonData, error) {
	endpoint := "/pokemon/" + input
	fullUrl := baseUrl + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil{
		return PokemonData{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err!=nil{
		return PokemonData{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 399 {
		return  PokemonData{} , fmt.Errorf("bad status code: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil{
		return PokemonData{}, err
	}
	pokemonDataResp := PokemonData{}
	err = json.Unmarshal(data, &pokemonDataResp)
	if err != nil{
		return PokemonData{}, err
	}
	return pokemonDataResp, nil
}