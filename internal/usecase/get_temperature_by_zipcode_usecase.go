package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GetTemperatureByZipcodeResponse struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type WeatherApiResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperatureByZipcode(zipcode string) (*GetTemperatureByZipcodeResponse, error) {
	zipcodeData, err := getZipcodeData(zipcode)

	if err != nil {
		return nil, err
	}

	weatherApiResponse, err := getWeatherData(zipcodeData.Localidade)

	if err != nil {
		return nil, err
	}

	celsiusTemp := weatherApiResponse.Current.TempC
	fahrenheitTemp := (celsiusTemp * 1.8) + 32
	kelvinTemp := celsiusTemp + 273

	getTemperatureByZipcodeResponse := GetTemperatureByZipcodeResponse{
		TempC: celsiusTemp,
		TempF: fahrenheitTemp,
		TempK: kelvinTemp,
	}

	return &getTemperatureByZipcodeResponse, nil
}

func getZipcodeData(zipcode string) (*ViaCepResponse, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + zipcode + "/json")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 404 {
		return nil, errors.New("can not find zipcode")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var viaCepResponse ViaCepResponse
	err = json.Unmarshal(body, &viaCepResponse)

	if err != nil {
		return nil, err
	}

	return &viaCepResponse, nil
}

func getWeatherData(city string) (*WeatherApiResponse, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=b5dde200e1bc40e5a61213318242805&q=%s&aqi=no", url.QueryEscape(city))
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 400 {
		return nil, errors.New("can not find temperature")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var weatherApiResponse WeatherApiResponse
	err = json.Unmarshal(body, &weatherApiResponse)

	if err != nil {
		return nil, err
	}

	return &weatherApiResponse, nil
}
