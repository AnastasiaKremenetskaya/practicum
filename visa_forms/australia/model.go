package australia

import (
	"practicum/visa_forms"
)

const (
	CountryName    = "Австралия"
	processingDays = 30
)

type Form struct {
	visa_forms.Base
	AllergicReactions string `json:"Аллергические реакции"`
}
