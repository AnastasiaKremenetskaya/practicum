package egypt

import "practicum/visa_forms"

const (
	CountryName    = "Египет"
	processingDays = 7
)

type Form struct {
	visa_forms.Base
	Job   string `json:"job"`
	Hotel string `json:"hotel"`
}
