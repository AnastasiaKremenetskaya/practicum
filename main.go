package main

import (
	"fmt"
	"practicum/visa_forms"
	"practicum/visa_forms/australia"
	"practicum/visa_forms/egypt"
	"practicum/visa_forms/germany"
)

type visaForm interface {
	ScanCountryData()
	GetProcessingDays() int
	GetCountryName() string
	PrintForm()
}

func main() {
	var (
		destinationCountry string
	)

	fmt.Println("Введите страну назначения:")
	fmt.Scan(&destinationCountry)

	var form visaForm
	switch destinationCountry {
	case germany.CountryName:
		form = &germany.Form{}
	case australia.CountryName:
		form = &australia.Form{}
	case egypt.CountryName:
		form = &egypt.Form{}
	default:
		fmt.Println("Ошибка: неизвестная страна назначения")
		return
	}
	form.ScanCountryData()
	form.PrintForm()
	fmt.Println()
	fmt.Printf(visa_forms.FinalText, form.GetCountryName(), form.GetProcessingDays())
}
