package main

type apiRequestDetail struct {
	httpMethod    string
	apiScheme     string
	apiBaseUri    string
	apiPath       string
	requestHeader map[string]string
	pathParam     map[string]string
	queryParam    map[string]string
	requestBody   map[string]string
}
