package dto

import "github.com/luanaands/multithreading-golang/internal/entity"

func FromBrasilApi(resp *entity.CepBrasilApiResponse) *CepResponse {
	return &CepResponse{
		Cep:          resp.Cep,
		Street:       resp.Street,
		Neighborhood: resp.Neighborhood,
		City:         resp.City,
		State:        resp.State,
	}
}

func FromViaCep(resp *entity.CepViaCepResponse) *CepResponse {
	return &CepResponse{
		Cep:          resp.Cep,
		Street:       resp.Logradouro,
		Neighborhood: resp.Bairro,
		City:         resp.Localidade,
		State:        resp.Uf,
	}
}
