package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ricardolindner/go-expert-cloud-run/internal/services"
	"github.com/ricardolindner/go-expert-cloud-run/internal/util"
)

type WeatherResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	if !util.IsValidCEP(cep) {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	cepInfo, err := services.GetCEPInfo(cep)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}

	weather, err := services.GetWeather(cepInfo.Localidade)
	if err != nil {
		http.Error(w, "can not find weather for this location", http.StatusNotFound)
		return
	}

	tempC := weather.Current.TempC
	tempF := tempC*1.8 + 32
	tempK := tempC + 273.0

	response := WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	jsonString := strings.Replace(string(jsonData), ":", ": ", -1)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonString))
}
