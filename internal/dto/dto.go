package dto

type CepResponse struct {
	Cep          string `json:"cep"`
	Street       string `json:"street"`       // logradouro / street
	Neighborhood string `json:"neighborhood"` // bairro / neighborhood
	City         string `json:"city"`         // localidade / city
	State        string `json:"state"`        // uf / state
}

type Response struct {
	Host     string      `json:"host"`
	Response CepResponse `json:"response"`
}
