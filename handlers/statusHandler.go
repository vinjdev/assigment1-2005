package handler

import (
	"fmt"
	"net/http"
	"time"
	"encoding/json"
)


/*
    Ensure the the REST method is get
*/
func StatusHandler(w http.ResponseWriter, r *http.Request, sTime time.Time) {
    switch r.Method {
    case http.MethodGet: getRequestStatus(w,r,sTime)
    default:
        http.Error(w, "REST Method'"+r.Method+"' not supported. Currently only'"+http.MethodGet+"' are supported.",http.StatusNotImplemented)
        return
    }

}
/*
    STATUS Get function
    request the api and returns the data in JSON fromat
*/
func getRequestStatus(w http.ResponseWriter, r *http.Request, sTime time.Time) {
    w.Header().Set("Content-Type", "application/json")

    client := &http.Client{}
    defer client.CloseIdleConnections()

    // ---------------- HANDLE COUNTRIES NOW API ----------------------
    reqCountriesNowApi, err := http.NewRequest(http.MethodGet,COUNTRIESNOW_API, nil)
    if err != nil {
        http.Error(w, "Error making request for COUNTRIESNOW api",http.StatusInternalServerError)
        return
    }
    
    resCountriesNowApi, err := client.Do(reqCountriesNowApi) 
    if err != nil {
        http.Error(w,"Error fetching request for COUNTRIESNOW API",http.StatusBadRequest)
        return
    }
    defer resCountriesNowApi.Body.Close()

    // OUTPUT FOR CONSOLE
    fmt.Printf("\nFETCHING DATA FROM COUNTRIESNOW\n")
	fmt.Printf("Status: %s\n", resCountriesNowApi.Status)
	fmt.Printf("Status Code: %d\n", resCountriesNowApi.StatusCode)
	fmt.Printf("Content Type: %s\n", resCountriesNowApi.Header.Get("content-type"))
	fmt.Printf("Protocol: %s\n", resCountriesNowApi.Proto)
	fmt.Printf("-----------\n")

    // -------------------- HANDLE RESTCOUNTRIES API ------------------
    reqRESTCountriesApi, err:= http.NewRequest(http.MethodGet,RESTCOUNTRY_API+"no",nil) // just add a working country code to check if it works
    if err != nil {
        http.Error(w, "Error making request for restcountries api",http.StatusInternalServerError)
        return 

    }
    resRESTCountriesApi, err := client.Do(reqRESTCountriesApi)
    if err != nil {
        http.Error(w, "Error making request for RESTCOUNTRY api",http.StatusInternalServerError)
        return

    }
    defer resRESTCountriesApi.Body.Close()

    // OUTPUT FOR CONSOLE
    fmt.Printf("\nFETCHING DATA FROM RESTCOUNTRY\n")
	fmt.Printf("Status: %s\n", resRESTCountriesApi.Status)
	fmt.Printf("Status Code: %d\n", resRESTCountriesApi.StatusCode)
	fmt.Printf("Content Type: %s\n", resRESTCountriesApi.Header.Get("content-type"))
	fmt.Printf("Protocol: %s\n", resCountriesNowApi.Proto)
	fmt.Printf("-----------\n")

    dataStatus, err := formatJSON(resCountriesNowApi, resRESTCountriesApi,sTime)
    if err != nil {
        http.Error(w,"Error formating the json",http.StatusInternalServerError)
        return

    }
    resJson, err := json.MarshalIndent(dataStatus, ""," ")
    if err != nil {
        http.Error(w,"Error formating the json",http.StatusInternalServerError)
        return
    }

    w.Write(resJson)

}

/*
    Returns the correct format used for status. caclualte the uptime for the service
*/
func formatJSON(resCountry *http.Response, resREST *http.Response,sTime time.Time) (statusResponse, error) {
    var data statusResponse
    
    uptime := time.Since(sTime).Seconds()
    
    data.Countriesnowapi = resCountry.Status
    data.Restcountriesapi = resREST.Status
    data.Version = "v1"
    data.Uptime = uptime
    return data, nil

}

