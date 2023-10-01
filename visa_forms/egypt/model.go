package egypt

import "practicum/visa_forms"

const (
	CountryName    = "Египет"
	processingDays = 7
)

type Form struct {
	visa_forms.Base
	Job   string `json:"Место работы"`
	Hotel string `json:"Название отеля"`
}
