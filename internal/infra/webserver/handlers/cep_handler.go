package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luanaands/multithreading-golang/internal/dto"
	"github.com/luanaands/multithreading-golang/internal/infra/service"
)

type CepHandler struct {
	Service service.CepInterface
}

func NewCepHandler(service service.CepInterface) *CepHandler {
	return &CepHandler{
		Service: service,
	}
}

// GetCep busca informações de CEP em paralelo (BrasilAPI e ViaCEP)
// @Summary Buscar CEP
// @Description Retorna dados do CEP consultando BrasilAPI e ViaCEP em paralelo
// @Tags CEP
// @Accept json
// @Produce json
// @Param cep query string true "CEP sem formatação (ex: 01001000)"
// @Success 200 {object} dto.CepResponse "CEP encontrado"
// @Failure 400 {object} map[string]string "CEP é obrigatório"
// @Failure 404 {object} map[string]string "CEP não encontrado"
// @Failure 500 {object} map[string]string "Erro interno"
// @Router /cep [get]
func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	brasilApiUrl := r.Context().Value("BrasilApHost").(string)
	viaCepUrl := r.Context().Value("ViaCepHost").(string)
	cep := r.URL.Query().Get("cep")

	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "CEP é obrigatório"})
		return
	}
	// FAZER UM CANAL
	responseChannel := make(chan *dto.CepResponse)
	// OUTRO CANAL
	responseChannel2 := make(chan *dto.CepResponse)

	go func() {
		//time.Sleep(1 * time.Second) // Simula demora na resposta

		response, err := h.Service.GetBrasilApi(cep, brasilApiUrl)
		if err != nil {
			responseChannel <- nil
			return
		}
		responseChannel <- response
	}()

	go func() {
		response, err := h.Service.GetViaCep(cep, viaCepUrl)
		if err != nil {
			responseChannel2 <- nil
			return
		}
		responseChannel2 <- response
	}()

	var result *dto.Response
	select {
	case res := <-responseChannel:
		if res != nil {
			result = &dto.Response{
				Host:     brasilApiUrl,
				Response: *res,
			}
		}
	case res2 := <-responseChannel2:
		if res2 != nil {
			result = &dto.Response{
				Host:     viaCepUrl,
				Response: *res2,
			}
		}
	}

	if result == nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Erro interno"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
