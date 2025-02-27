package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*
   Ensure the the REST method is get
*/
func InfoHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet: handleGetRequest(w,r)
    default:
        http.Error(w, "REST Method'"+r.Method+"' not supported. Currently only'"+http.MethodGet+"' are supported.",http.StatusNotImplemented)
        return
    } 
}

/*
    INFO Get function
    request the api and returns the data in JSON fromat
*/
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // --------------------- READING URL -------------------
    
    cc := r.PathValue("val")
    if len(cc) != 2 {
        http.Error(w, "Error reading the country code", http.StatusBadRequest)
        return
    }
    apiUrl := RESTCOUNTRY_API + cc

    limit := r.URL.Query().Get("limit")
    limitInt := 0
    
    if limit == "" {
        limitInt = 0
    } else {
        var err error
        limitInt,err = strconv.Atoi(limit)
        if err != nil {
            http.Error(w,"Error could not read the limit, use a number", http.StatusBadRequest)
            return
        }
    }
    
    fmt.Println("Limit:",limitInt)

    // ----------------- FETCHING API -----------------

    client := &http.Client{}
    defer client.CloseIdleConnections()

    // ----------------------- HANDLE THE COUNTRY ---------------------
    // request country
    reqCountry, err := http.NewRequest(http.MethodGet,apiUrl,nil) 
    if err != nil {
        http.Error(w, "Error creating request for countries",http.StatusInternalServerError)
        return
    }
    // response country 
    resCountry, err := client.Do(reqCountry)
    if err != nil {
        http.Error(w, "Error fetching data for countries",http.StatusInternalServerError)
        return
    }
    defer resCountry.Body.Close()
    
    // fetch the country and return a json in struct
    dataCountry, err := decodeApiCountriesBody(resCountry)
    if err != nil {
        http.Error(w, "Error decoding json",http.StatusInternalServerError)
        return
    }
    // ----------------- INFO LOG -----------------

    fmt.Printf("\n-----------\n")
    fmt.Printf("\nFETCHING DATA FROM RESTCOUNTRY\n")
	fmt.Printf("Status: %s\n", resCountry.Status)
	fmt.Printf("Status Code: %d\n", resCountry.StatusCode)
	fmt.Printf("Content Type: %s\n", resCountry.Header.Get("content-type"))
	fmt.Printf("Protocol: %s\n", resCountry.Proto)
	fmt.Printf("-----------\n")
    
    // ----------------  HANDLE THE CITIES --------------------------
    requestBody, err := json.Marshal(map[string]string{"country": dataCountry.Name})
    if err != nil {
        http.Error(w, "Error encoding JSON payload", http.StatusInternalServerError)
        return
    }

    reqCities, err := http.NewRequest(http.MethodPost,COUNTRIESNOW_API_CITIES,bytes.NewBuffer(requestBody))
    if err != nil {
        http.Error(w,"Error creating request for cities",http.StatusInternalServerError)
        return
    }
    reqCities.Header.Set("Content-Type","application/json")

    resCities, err := client.Do(reqCities)
    if err != nil {
        http.Error(w,"Error fetcing data for cities",http.StatusInternalServerError)
        return
    }
    defer resCities.Body.Close()

    dataCities, err := decodeApiCitiesBody(resCities,limitInt)
    if err != nil {
        http.Error(w,"Error decoding json",http.StatusInternalServerError)
        return
    }


    // COUNTRYNOW API
    fmt.Printf("\n-----------\n")
    fmt.Printf("\nFETCHING DATA FROM COUNTRIESNOW\n")
	fmt.Printf("Status: %s\n", resCities.Status)
	fmt.Printf("Status Code: %d\n", resCities.StatusCode)
	fmt.Printf("Content Type: %s\n", resCities.Header.Get("content-type"))
	fmt.Printf("Protocol: %s\n", resCities.Proto)
	fmt.Printf("-----------\n")
     
    // --------- RESPOSE JSON --------------

    // convert the go lang struct format to json data
    dataCountry.Cities = dataCities
    resJson, err := json.MarshalIndent(dataCountry, "", " ")
    if err != nil {
        http.Error(w, "Error encoding JSON",http.StatusInternalServerError)
        return
    }
    w.Write(resJson)    
}

/*
    Decodes the api from RESTCOUNTRIES api and returns the correct format
*/
func decodeApiCountriesBody(r *http.Response) (CountryResponse, error) {
    var data CountryResponse 
    var apiData []CountryRequest
    err := json.NewDecoder(r.Body).Decode(&apiData)
    if err != nil {
        return data, err
    }

    // check if API returned data
    if len(apiData) == 0 {
        return data, fmt.Errorf("no data found in api")
    }

    data = CountryResponse {
        Name:        apiData[0].Name.Common,
        Continents:  apiData[0].Continents,
        Population:  apiData[0].Population,
        Languages:   apiData[0].Languages,
        Borders:     apiData[0].Borders,
        Capital:     apiData[0].Capital[0],
        Cities:      []string{},           // read this later in the other api
    }
    
    return data, nil
}
/*
    Decodes the cities from countriesnow api and returns the cities in a list of strings
*/
func decodeApiCitiesBody(r *http.Response,limit int) ([]string, error) {
    var data cityRequest
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        return nil, err
    }

    if len(data.Data) == 0 {
        return nil, fmt.Errorf("no cities found")
    }
    cities := data.Data

    if limit > 0 && limit < len(cities) {
        cities = cities[:limit]
    }
    
    return cities, nil
}


