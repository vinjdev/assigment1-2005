package handler

const (
    DEFAULT_PATH = "/countryinfo/v1/"
    INFO_PATH =  DEFAULT_PATH + "info/{val}" // {?limit=10}
    POPULATION_PATH = DEFAULT_PATH + "population/"
    STATUS_PATH = DEFAULT_PATH + "status/"
    COUNTRY_API = "http://129.241.150.113:8080/v3.1/alpha/"
    CITIES_API  = "https://countriesnow.space/api/v0.1/countries"
)
