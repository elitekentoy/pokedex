package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (RespShallowLocations, error) {
	url := Pokeapi_location_area_url
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}

func (c *Client) LocationDetails(pageUrl *string) (RespFullLocations, error) {
	url := *pageUrl

	// Check if Exists in Cache
	if val, ok := c.cache.Get(url); ok {
		locationResp := RespFullLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespFullLocations{}, err
		}

		return locationResp, nil
	}

	// Create Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespFullLocations{}, err
	}

	// Execute Request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespFullLocations{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespFullLocations{}, err
	}

	locationResp := RespFullLocations{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespFullLocations{}, err
	}

	c.cache.Add(url, data)

	return locationResp, nil
}
