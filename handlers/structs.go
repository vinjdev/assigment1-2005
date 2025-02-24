package handler

// ---------------------- Info ------------------

/*
    Used for correct json format, before marshalindent
*/
type CountryResponse struct {
    Name        string            `json:"name"`
    Continents  []string          `json:"continents"`
    Population  int               `json:"population"`
    Languages   map[string]string `json:"languages"`
    Borders     []string          `json:"borders"`
    Flag        string            `json:"flag"`
    Capital     string          `json:"capital"`
    Cities      []string          `json:"cities"`
}

/*
    The format used by the api for country data
*/
type CountryRequest struct {
    Name        struct {Common string} `json:"name"`
    Continents  []string               `json:"continents"`
    Population  int                    `json:"population"`
    Languages   map[string]string      `json:"languages"`
    Borders     []string               `json:"borders"`
    Flag        string                 `json:"flag"`
    Capital     []string               `json:"capital"`
    Cities      []string               `json:"cities"`
}

/*
    Format used by the api. Returns as string as this string is used in the cities array
*/
type cityRequest struct {
    Error bool             `json:"error"`
    Msg   string           `json:"msg"`
    Data  [] struct {
        Country string     `json:"country"`
        Cities []string    `json:"cities"`
    }                      `json:"data"`
}

// -------------- population ------------------------

/*
    The correct format for printing out the values in population
*/
type populationResponse struct {
    Mean   int         `json:"mean"`
    Values [] struct {
        Year int       `json:"year"`
        Value int      `json:"value"`
    }                  `json:"values"`
}

/*
    The json format for reading the data from api
*/
type populationAPIRequest struct {
    Error bool          `json:"error"`
    Msg string          `json:"msg"`
    Data [] struct {
        Country string  `json:"country"`
        Iso3 string    ` json:"iso3"`
        PopulationCounts [] struct {
            Year int    `json:"year"`
            Value int   `json:"value"`
        }               `json:"populationCounts"`        
    }                   `json:"data"`
}

type populationAPIResponse struct {
    Country string   `json:"country"`
    PopulationCounts [] struct {
            Year int    `json:"year"`
            Value int   `json:"value"`
        }               `json:"populationCounts"`
}

// -------------------------- STATUS --------------

type statusResponse struct {
    Countriesnowapi  string   `json:"countriesnowapi"`
    Restcountriesapi string   `json:"restcountriesapi"`
    Version          string   `json:"version"`
    Uptime           string   `json:"uptime"`
}






