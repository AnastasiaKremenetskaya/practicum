package main

import (
	"fmt"
	"practicum/visa_forms"
	"practicum/visa_forms/australia"
	"practicum/visa_forms/egypt"
	"practicum/visa_forms/germany"
)

type fillableForm interface {
	ScanCountryData()
	ScanBaseData()
}

type processableForm interface {
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

	var form fillableForm
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
	if v, ok := form.(processableForm); ok {
		PrintResult(v)
	}
}

func PrintResult(form processableForm) {
	processingText := fmt.Sprintf(visa_forms.FinalText, form.GetCountryName(), form.GetProcessingDays())
	fmt.Println(processingText)
	form.PrintForm()
}
