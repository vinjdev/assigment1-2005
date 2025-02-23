package handler

import (
    "fmt"
    "net/http"
)


func EmptyHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-type","text/html")

    output := "Root level. See <a href=\"" + INFO_PATH + "\">Info</a>"
    output += " or <a href=\"" + POPULATION_PATH + "\">Population</a>"
    output += " or <a href=\"" + STATUS_PATH + "\">Status</a>"
    
    _,err := fmt.Fprintf(w,"%v",output)
    if err != nil {
        http.Error(w,"Error when returning output",http.StatusInternalServerError)
    }
}
