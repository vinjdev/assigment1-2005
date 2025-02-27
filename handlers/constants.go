package handler

const (
    DEFAULT_PATH =                    "/countryinfo/v1/"

    INFO_PATH =  DEFAULT_PATH +       "info/{val}"         // {?limit=10}
    POPULATION_PATH = DEFAULT_PATH +  "population/{val}"   // {?limit=10}
    STATUS_PATH = DEFAULT_PATH +      "status/"

    RESTCOUNTRY_API =                 "http://129.241.150.113:8080/v3.1/alpha/"
    COUNTRIESNOW_API  =               "https://countriesnow.space/api/v0.1/countries"
    COUNTRIESNOW_API_CITIES =         "https://countriesnow.space/api/v0.1/countries/cities"
    POPULATION_API =                  "https://countriesnow.space/api/v0.1/countries/population"
)
