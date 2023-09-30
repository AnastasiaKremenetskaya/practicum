package australia

import "fmt"

func (f *Form) ScanCountryData() {
	f.ScanBaseData()
	fmt.Println("Наличие прививки от лихорадки Денге (Да/Нет):")
	var answer hasDengueVaccination
	fmt.Scan(&answer)
	switch answer {
	case yes:
		f.HasDengueVaccination = true
	}
}
