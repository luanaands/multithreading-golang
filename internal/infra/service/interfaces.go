package service

import "github.com/luanaands/multithreading-golang/internal/dto"

type CepInterface interface {
	GetBrasilApi(cep string, url string) (*dto.CepResponse, error)
	GetViaCep(cep string, url string) (*dto.CepResponse, error)
}
