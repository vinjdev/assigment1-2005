package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func InfoHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet: handleGetRequest(w,r)
    default:
        http.Error(w, "REST Method'"+r.Method+"' not supported. Currently only'"+http.MethodGet+"' are supported.",http.StatusNotImplemented)
        return
    } 
}

// bruker api_url appender country_code, og printer ut json format
func handleGetRequest(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    // ----------------- FETCHING API -----------------
    
    cc := r.PathValue("val")
    if cc == "" {
        http.Error(w, "Missing country code in URL", http.StatusBadRequest)
        return
    }
    apiUrl := COUNTRY_API + cc

    client := &http.Client{}
    defer client.CloseIdleConnections()

    // HANDLE THE COUNTRY
    // request country
    reqCountry, err := http.NewRequest(http.MethodGet,apiUrl,nil) 
    if err != nil {
        http.Error(w, "Error creating request for countries: "+err.Error(),http.StatusInternalServerError)
        return
    }
    // response country 
    resCountry, err := client.Do(reqCountry)
    if err != nil {
        http.Error(w, "Error fetching data for countries: "+err.Error(),http.StatusInternalServerError)
        return
    }
    defer resCountry.Body.Close()
    
    // fetch the country and return a json in struct
    dataCountry, err := fetchApiCountriesBody(resCountry)
    if err != nil {
        http.Error(w, "Error decoding json: "+err.Error(),http.StatusInternalServerError)
    }
    
    // HANDLE THE CITIES
    reqCities, err := http.NewRequest(http.MethodGet,CITIES_API,nil)
    if err != nil {
        http.Error(w,"Error creating request for cities: "+err.Error(),http.StatusInternalServerError)
    }

    resCities, err := client.Do(reqCities)
    if err != nil {
        http.Error(w,"Error fetcing data for cities: "+err.Error(),http.StatusInternalServerError)
    }
    defer resCities.Body.Close()

    dataCities, err := fetchApiCitiesBody(resCities,dataCountry.Name)
    if err != nil {
        http.Error(w,"Error decoding json: "+err.Error(),http.StatusInternalServerError)
    }


    // ----------------- INFO LOG -----------------

    // COUNTRY API
    fmt.Printf("\n-----------\n")
    fmt.Printf("\nFETCHING DATA COUNTRIES\n")
	fmt.Printf("Status: %s\n", resCountry.Status)
	fmt.Printf("Status Code: %d\n", resCountry.StatusCode)
	fmt.Printf("Content Type: %s\n", resCountry.Header.Get("content-type"))
	fmt.Printf("Protocol: %s\n", resCountry.Proto)
	fmt.Printf("-----------\n")

    // CITIES API
    fmt.Printf("\n-----------\n")
    fmt.Printf("\nFETCHING DATA CITIES\n")
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
        http.Error(w, "Error encoding JSON: "+err.Error(),http.StatusInternalServerError)
        return
    }
    w.Write(resJson)    
}

func fetchApiCountriesBody(r *http.Response) (Country, error) {
    var data Country 
    var apiData []CountryResponse
    err := json.NewDecoder(r.Body).Decode(&apiData)
    if err != nil {
        return data, err
    }

    // check if API returned data
    if len(apiData) == 0 {
        return data, fmt.Errorf("no data found in api")
    }

    data = Country {
        Name:        apiData[0].Name.Common,
        Continents:  apiData[0].Continents,
        Population:  apiData[0].Population,
        Languages:   apiData[0].Languages,
        Borders:     apiData[0].Borders,
        Capital:     apiData[0].Capital[0],
        Cities:      []string{},
    }
    
    return data, nil
}

func fetchApiCitiesBody(r *http.Response,countryName string) ([]string, error) {
    var data cityResponse
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        return nil, err
    }

    for _, val := range data.Data {
        if val.Country == countryName {
            return val.Cities, nil 
        }
    } 
    return nil, fmt.Errorf("no cities found for country: %s", countryName)
}


