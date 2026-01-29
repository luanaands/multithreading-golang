package service

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/luanaands/multithreading-golang/internal/dto"
	"github.com/luanaands/multithreading-golang/internal/entity"
)

type CepService struct {
	client *http.Client
}

func NewCepService() *CepService {
	return &CepService{
		client: &http.Client{},
	}
}

func (s *CepService) GetBrasilApi(cep string, url string) (*dto.CepResponse, error) {
	req, err := http.NewRequest("GET", url+"/"+cep, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *entity.CepBrasilApiResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	var dtoResponse *dto.CepResponse
	dtoResponse = dto.FromBrasilApi(response)
	return dtoResponse, nil
}

func (s *CepService) GetViaCep(cep string, url string) (*dto.CepResponse, error) {
	req, err := http.NewRequest("GET", url+"/"+cep+"/json", nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *entity.CepViaCepResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	var dtoResponse *dto.CepResponse
	dtoResponse = dto.FromViaCep(response)
	return dtoResponse, nil
}
