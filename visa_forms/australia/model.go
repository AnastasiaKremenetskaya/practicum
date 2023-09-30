package australia

import (
	"practicum/visa_forms"
)

type hasDengueVaccination string

const (
	CountryName    = "Австралия"
	processingDays = 30
)

const (
	yes hasDengueVaccination = "Да"
	no                       = "Нет"
)

type Form struct {
	visa_forms.Base
	HasDengueVaccination bool `json:"has_dengue_vaccination"`
}
