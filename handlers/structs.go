package handler


type Country struct {
    Name        string            `json:"name"`
    Continents  []string          `json:"continents"`
    Population  int               `json:"population"`
    Languages   map[string]string `json:"languages"`
    Borders     []string          `json:"borders"`
    Flag        string            `json:"flag"`
    Capital     string          `json:"capital"`
    Cities      []string          `json:"cities"`
}

type CountryResponse struct {
    Name        struct {Common string} `json:"name"`
    Continents  []string               `json:"continents"`
    Population  int                    `json:"population"`
    Languages   map[string]string      `json:"languages"`
    Borders     []string               `json:"borders"`
    Flag        string                 `json:"flag"`
    Capital     []string               `json:"capital"`
    Cities      []string               `json:"cities"`
}

type cityResponse struct {
    Error bool             `json:"error"`
    Msg   string           `json:"msg"`
    Data  [] struct {
        Country string     `json:"country"`
        Cities []string    `json:"cities"`
    }                      `json:"data"`
}


