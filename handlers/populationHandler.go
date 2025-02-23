package handler

import (
    "fmt"
    "net/http"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet: getRequestPopulation(w,r)
    default:
        http.Error(w, "REST Method'"+r.Method+"' not supported. Currently only'"+http.MethodGet+"' are supported.",http.StatusNotImplemented)
        return
    }
}

func getRequestPopulation(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    output := "Welcome to population page"

    _, err := fmt.Fprintf(w,"%v",output)
    if err != nil {
        http.Error(w, "Error when output", http.StatusInternalServerError)
    }

}
