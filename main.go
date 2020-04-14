package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	fmt.Println("Starting the application...")

	get_apiRequest := apiRequestDetail{
		httpMethod:    "GET",
		apiScheme:     "https",
		apiBaseUri:    "example.com",
		apiPath:       "path/{pathParam1}/to/{pathParam2}/get/api",
		requestHeader: map[string]string{"auth": "token"},
		pathParam:     map[string]string{"pathParam1": "pathParam1Value", "pathParam2": "pathParam2Value"},
		queryParam:    map[string]string{"contact": "918152623411"},
		requestBody:   map[string]string{"id": "123123", "name": "Widget Adapter11", "manufacturer": "d290f1ee-6c54-4b01-90e6-d701748f0851", "releaseDate": "2016-08-29T09:12:33.001Z"},
	}

	apiRequest := apiRequestDetail{
		httpMethod:    "POST",
		apiScheme:     "https",
		apiBaseUri:    "example.com",
		apiPath:       "path/{pathParam1}/to/{pathParam2}/post/api",
		requestHeader: map[string]string{"auth": "token"},
		pathParam:     map[string]string{"pathParam1": "pathParam1Value", "pathParam2": "pathParam2Value"},
		queryParam:    map[string]string{"contact": "918152623411"},
		requestBody:   map[string]string{"id": "123123", "name": "Widget Adapter11", "manufacturer": "d290f1ee-6c54-4b01-90e6-d701748f0851", "releaseDate": "2016-08-29T09:12:33.001Z"},
	}
	fmt.Println(get_apiRequest)

	urlResponse, err := request_url_builder(apiRequest.apiScheme, apiRequest.apiBaseUri, apiRequest.apiPath, apiRequest.pathParam, apiRequest.queryParam)
	if err != nil {
		fmt.Printf("url generation failed with error %s\n", err)
	}
	client := &http.Client{}
	request, err := http.NewRequest(apiRequest.httpMethod, urlResponse, request_data_builder(apiRequest.requestBody))
	if err != nil {
		fmt.Printf("url generation failed with error %s\n", err)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("url generation failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)

	result := map[string]string{"statusCode": string(response.StatusCode), "responseData": string(data)}

	fmt.Println(result)
}

func request_data_builder(requestBody map[string]string) *bytes.Buffer {
	if requestBody != nil {
		jsonValue, _ := json.Marshal(requestBody)
		return bytes.NewBuffer(jsonValue)
	} else {
		return nil
	}
}

func request_url_builder(apiScheme string, apiBaseUri string, apiPath string, pathParam map[string]string, queryParam map[string]string) (string, error) {
	apiUrl := url.URL{
		Scheme: apiScheme,
		Host:   apiBaseUri,
		Path:   apiPath,
	}

	if queryParam != nil {
		params := url.Values{}
		for key, value := range queryParam {
			params.Add(key, value)
		}
		apiUrl.RawQuery = params.Encode()
	}

	updatedUrl, err := url.QueryUnescape(apiUrl.String())
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if pathParam != nil {
		for key, value := range pathParam {
			updatedUrl = strings.ReplaceAll(updatedUrl, "{"+key+"}", value)
		}
	}
	return updatedUrl, nil
}
