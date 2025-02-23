package handler

import (
    "fmt"
    "net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    output := "Welcome to status page"

    _, err := fmt.Fprintf(w,"%v",output)
    if err != nil {
        http.Error(w, "Error when output", http.StatusInternalServerError)
    }

}


