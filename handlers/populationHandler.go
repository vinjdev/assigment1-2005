package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
    Ensure the the REST method is get
*/
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet: getRequestPopulation(w,r)
    default:
        http.Error(w, "REST Method'"+r.Method+"' not supported. Currently only'"+http.MethodGet+"' are supported.",http.StatusNotImplemented)
        return
    }
}
/*
    POPULATION Get function
    request the api and returns the data in JSON fromat
*/
func getRequestPopulation(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    
    // ----------------- READING URL ------------------
    cc := r.PathValue("val")
    if len(cc) != 2 {
        http.Error(w, "Error reading the country code",http.StatusBadRequest) // not a server error
        return                                                                // therefore reuturn 400
    }
    limit := r.URL.Query().Get("limit")
    startYear := 0
    endYear   := 0

    if limit != "" {
        years := strings.Split(limit,"-")
        if (len(years) == 2 && len(years[0]) == 4 && len(years[1]) == 4) {
            var err error

            startYear, err = strconv.Atoi(years[0])
            if err != nil {
                http.Error(w,"Error reading the limit, write a year between 1960 and 2018",http.StatusBadRequest)
                return
            }
            endYear, err = strconv.Atoi(years[1])
            if err != nil {
                http.Error(w,"Error reading the limit, write a year between 1960 and 2018",http.StatusBadRequest)
                return
            }
            if years[0] > years[1] {
                http.Error(w,"First year provided needs to be lower than the second year",http.StatusBadRequest)
                return
            }
            
        } else {
            http.Error(w,"Error reading the limit, write a year between 1960 and 2018",http.StatusBadRequest)
            return
        }
    }

    // ----------------- FETCHING API -----------------
    apiURL := POPULATION_API
    apiNAME := RESTCOUNTRY_API + cc

    client := &http.Client{}
    defer client.CloseIdleConnections()

    // ------------------ handle the NAME api (country api) ------------------------
    reqName, err := http.NewRequest(http.MethodGet, apiNAME, nil)
    if err != nil {
        http.Error(w,"Error making request for the RESTCOUNTRY api", http.StatusInternalServerError)
        return
    }

    resName, err := client.Do(reqName)
    if err != nil {
        http.Error(w, "Error fetching request for the RESTCOUNTRY api", http.StatusInternalServerError)
        return
    }
    defer resName.Body.Close()

    dataName, err := decodeAPIName(resName)
    if err != nil {
        http.Error(w, "Error decoding json", http.StatusInternalServerError)
        return
    }
    // ------------------------ OUTPUT IN CONSOLE ------------------------------
    fmt.Printf("\nFETCHING DATA COUNTRIES USED FOR NAME\n")
	fmt.Printf("Status: %s\n", resName.Status)
	fmt.Printf("Status Code: %d\n", resName.StatusCode)
	fmt.Printf("Content Type: %s\n", resName.Header.Get("content-type"))
	fmt.Printf("Protocol: %s\n", resName.Proto)
	fmt.Printf("-----------\n")

    fmt.Println("\nName for Country:",dataName)
    
    // ----------------------- HANDLE THE POPULATION API ---------------------
    // request population api
    reqPopulation,err := http.NewRequest(http.MethodGet,apiURL,nil)
    if err != nil {
        http.Error(w,"Error making request for the COUNTRYNOW POPULATION api",http.StatusInternalServerError)
        return
    }

    resPopulation,err := client.Do(reqPopulation)
    if err != nil {
        http.Error(w, "Error fetching request for the COUNTRYNOW POPULATION api",http.StatusInternalServerError)
        return
    }
    defer resPopulation.Body.Close()

    dataPopulation,err := decodeAPIPopulation(resPopulation, dataName)
    if err != nil {
        http.Error(w, "Error decoding json",http.StatusInternalServerError)
        return
    }

    // ------------------------ OUTPUT IN CONSOLE ------------------------------
    fmt.Printf("\n-----------")
    fmt.Printf("\nFETCHING DATA COUNTRIES\n")
	fmt.Printf("Status: %s\n", resPopulation.Status)
	fmt.Printf("Status Code: %d\n", resPopulation.StatusCode)
	fmt.Printf("Content Type: %s\n", resPopulation.Header.Get("content-type"))
	fmt.Printf("Protocol: %s\n", resPopulation.Proto)
	fmt.Printf("-----------\n")
    
    
    // --------------------------- FORMATING THE JSON --------------------------
    
    data, err := correctFormat(dataPopulation, startYear, endYear)
    if err !=  nil {
        http.Error(w, "Error formating the json",http.StatusInternalServerError)
    }

    // Marrshall 
    resJson, err := json.MarshalIndent(data, "", " ")
    if err != nil {
        http.Error(w, "Error encoding JSON",http.StatusInternalServerError)
        return	
    }
    w.Write(resJson)

}

/*
    Decodes the RESTCOUNTRY API, to extract the name in strnig format
*/
func decodeAPIName(r *http.Response) (string, error) {
    var country []CountryRequest      // api format
    err := json.NewDecoder(r.Body).Decode(&country)
    if err != nil {
        return "", err
    }
    
    name := country[0].Name.Common 

    return name, nil
}

/*
    decodes the Countrynow api in population to extract the population data, with the correct name
*/
func decodeAPIPopulation(r *http.Response, name string) (populationAPIResponse,error) {
    var requestJSON  populationAPIRequest
    var responseJSON populationAPIResponse
    err := json.NewDecoder(r.Body).Decode(&requestJSON)  
    if err != nil {
        return responseJSON, err
    }

    responseJSON.Country = name
    
    for _, val := range requestJSON.Data {
        if val.Country == name {
            responseJSON.PopulationCounts = val.PopulationCounts 
        }
    }

    return responseJSON, nil
}


/*
    Returns the correct format needed for populatin json. takes in start and end year, so it can extract the years if requested
*/
func correctFormat(populationData populationAPIResponse, start int, end int) (populationResponse, error) {
    var data populationResponse
    
    // for the population data
    if start == 0 && end == 0 {
        data.Values = populationData.PopulationCounts 
    } else {
        for _, val := range populationData.PopulationCounts {
            if val.Year >= start && val.Year <= end {
                data.Values = append(data.Values, val)
            }
        }
    }

    // calculating mean
    total := 0
    for _, val := range data.Values {
        total += val.Value
    }

    if total != 0 {
        data.Mean = total / len(data.Values)
    }
     
    
    return data, nil
}

