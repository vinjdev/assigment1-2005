package handler

import (
    "fmt"
    "net/http"
)



func StatusHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet: getRequestStatus(w,r)
    default:
        http.Error(w, "REST Method'"+r.Method+"' not supported. Currently only'"+http.MethodGet+"' are supported.",http.StatusNotImplemented)
        return
    }

}

func getRequestStatus(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    client := &http.Client{}
    defer client.CloseIdleConnections()

    // ---------------- HANDLE COUNTRIES NOW API ----------------------
    reqCountriesNowApi, err := http.NewRequest(http.MethodGet,COUNTRIESNOW_API, nil)
    if err != nil {
        http.Error(w, "Error making request for name api:"+err.Error(),http.StatusInternalServerError)
        return
    }
    
    resCountriesNowApi, err := client.Do(reqCountriesNowApi) 
    if err != nil {
        http.Error(w,"Error fetching request for COUNTRIES NOW API"+err.Error(),http.StatusBadRequest)
        return
    }
    defer resCountriesNowApi.Body.Close()

    // -------------------- HANDLE RESTCOUNTRIES API ------------------

    fmt.Printf("\nFETCHING DATA COUNTRIES USED FOR NAME\n")
	fmt.Printf("Status: %s\n", resCountriesNowApi.Status)
	fmt.Printf("Status Code: %d\n", resCountriesNowApi.StatusCode)
	fmt.Printf("Content Type: %s\n", resCountriesNowApi.Header.Get("content-type"))
	fmt.Printf("Protocol: %s\n", resCountriesNowApi.Proto)
	fmt.Printf("-----------\n")

}

