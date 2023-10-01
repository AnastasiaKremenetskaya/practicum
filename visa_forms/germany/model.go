package germany

import (
	"practicum/visa_forms"
)

const (
	CountryName    = "Германия"
	processingDays = 14
)

type Form struct {
	visa_forms.Base
	GermanLanguageLevel string `json:"Уровень немецкого языка"`
}
