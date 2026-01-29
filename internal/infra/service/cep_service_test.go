package service

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetViaCep(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/01001000/json" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"cep":"01001-000","logradouro":"Praça da Sé","bairro":"Sé","localidade":"São Paulo","uf":"SP","ddd":"11"}`)
	}))
	defer ts.Close()

	s := &CepService{client: ts.Client()}
	resp, err := s.GetViaCep("01001000", ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, resp.Cep, "01001-000")
	assert.Equal(t, resp.Street, "Praça da Sé")
	assert.Equal(t, resp.Neighborhood, "Sé")
	assert.Equal(t, resp.City, "São Paulo")
	assert.Equal(t, resp.State, "SP")
}

func TestBrasilApi(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/01001000" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{"cep":"01001000","state":"SP","city":"São Paulo","neighborhood":"Sé","street":"Praça da Sé","service":"open-cep"}`)
	}))
	defer ts.Close()

	s := &CepService{client: ts.Client()}
	resp, err := s.GetBrasilApi("01001000", ts.URL)
	assert.Nil(t, err)
	assert.Equal(t, resp.Cep, "01001000")
	assert.Equal(t, resp.Street, "Praça da Sé")
	assert.Equal(t, resp.Neighborhood, "Sé")
	assert.Equal(t, resp.City, "São Paulo")
	assert.Equal(t, resp.State, "SP")
}
