package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nomada-sh/levita-stp/models"
)

const urlSTPDeposito = "/speiws/rest/ordenPago/registra"

func (ucase *usecase) dispersionRequest(dispersion models.DispersionInput) (*models.Dispersion, error) {
	var result models.Dispersion
	client := &http.Client{}
	url := fmt.Sprintf("%v%v", ucase.STPAPIURL, urlSTPDeposito)
	jsonDispersion, err := json.Marshal(dispersion)
	if err != nil {
		return nil, err
	}

	body := bytes.NewBuffer(jsonDispersion)
	request, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
