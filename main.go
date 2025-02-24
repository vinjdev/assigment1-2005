package main

import (
    "example.com/assignment1/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
    "time"
    
)



func main()  {

    fmt.Println("Starting server...") 

    var startTime time.Time

    startTime = time.Now()

    // Setting a port
    port := os.Getenv("Port")

    if port == "" {
        log.Println("$PORT has not been set. Default: 8080")
        port = "8080"
    }

    // init the handler for endpoints
    http.HandleFunc(handler.DEFAULT_PATH,handler.EmptyHandler)
    http.HandleFunc(handler.INFO_PATH, handler.InfoHandler)
    http.HandleFunc(handler.POPULATION_PATH, handler.PopulationHandler)
    http.HandleFunc(handler.STATUS_PATH, func(w http.ResponseWriter, r *http.Request) {
        handler.StatusHandler(w,r,startTime)
    })
 
    // start server
    fmt.Println("Starting server on " + port)
    fmt.Println("http://localhost:8080/countryinfo/v1")
    log.Fatal(http.ListenAndServe(":"+port,nil))

}
