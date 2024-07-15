package main

import (
	"encoding/json"
	"github.com/enermax626/go-postalcode-temperature/internal/dao"
	"github.com/enermax626/go-postalcode-temperature/internal/model"
	"github.com/enermax626/go-postalcode-temperature/internal/service"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/weather/{postalCode}", findWeatherByPostalCode())
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Something went wrong when running the HTTP server...", err)
	}
}

func findWeatherByPostalCode() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var addressService = service.NewAddressService(dao.NewAddressDao())
		weatherService := service.NewWeatherService(addressService, dao.NewWeatherDao())

		postalCode := r.PathValue("postalCode")
		weatherResponse, err := weatherService.FindWeatherByPostalCode(postalCode)
		if err != nil {
			switch err {
			case model.ErrPostalCodeNotFound:
				w.WriteHeader(http.StatusNotFound)
				w.Write(MarshalResponse(ErrorResponse{
					Message: err.Error(),
				}))
			case model.ErrInvalidPostalCode:
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write(MarshalResponse(ErrorResponse{
					Message: err.Error(),
				}))
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(MarshalResponse(ErrorResponse{
					Message: "Internal server error",
				}))
			}
			return
		}
		_, _ = w.Write(MarshalResponse(weatherResponse))
	}
}

func MarshalResponse(res interface{}) []byte {
	marshalledResponse, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		return []byte{}
	}
	return marshalledResponse
}
