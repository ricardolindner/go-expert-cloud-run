package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CEPInfo struct {
	Localidade string `json:"localidade"`
	Erro       string `json:"erro,omitempty"`
}

func GetCEPInfo(cep string) (*CEPInfo, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not find cep: %s", cep)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println("ViaCEP complete response: ", string(bodyBytes))

	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var cepInfo CEPInfo
	if err := json.NewDecoder(resp.Body).Decode(&cepInfo); err != nil {
		return nil, err
	}

	if cepInfo.Erro == "true" {
		return nil, fmt.Errorf("cep not found")
	}

	return &cepInfo, nil
}
