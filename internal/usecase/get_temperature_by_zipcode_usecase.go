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

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func GetTemperatureByZipcode(zipcode string) (*GetTemperatureByZipcodeResponse, *ErrorResponse) {
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

func getZipcodeData(zipcode string) (*ViaCepResponse, *ErrorResponse) {
	url := "http://viacep.com.br/ws/" + url.PathEscape(zipcode) + "/json"
	resp, err := http.Get(url)

	if err != nil {
		return nil, buildErrorResponse(http.StatusInternalServerError, err)
	}

	if resp.StatusCode == 404 {
		return nil, buildErrorResponse(http.StatusUnprocessableEntity, errors.New("can not find zipcode"))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, buildErrorResponse(http.StatusInternalServerError, err)
	}

	var viaCepResponse ViaCepResponse
	err = json.Unmarshal(body, &viaCepResponse)

	if err != nil {
		return nil, buildErrorResponse(http.StatusInternalServerError, err)

	}

	fmt.Printf("viaCepResponse: %+v\n", viaCepResponse)

	return &viaCepResponse, nil
}

func getWeatherData(city string) (*WeatherApiResponse, *ErrorResponse) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=b5dde200e1bc40e5a61213318242805&q=%s&aqi=no", url.QueryEscape(city))
	resp, err := http.Get(url)

	if err != nil {
		return nil, buildErrorResponse(http.StatusInternalServerError, err)
	}

	if resp.StatusCode == 400 {
		errorMessage := fmt.Sprintf("can not find weather for city %s", city)
		return nil, buildErrorResponse(http.StatusNotFound, errors.New(errorMessage))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, buildErrorResponse(http.StatusInternalServerError, err)
	}

	var weatherApiResponse WeatherApiResponse
	err = json.Unmarshal(body, &weatherApiResponse)

	if err != nil {
		return nil, buildErrorResponse(http.StatusInternalServerError, err)
	}

	return &weatherApiResponse, nil
}

func buildErrorResponse(statusCode int, err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
	}
}
