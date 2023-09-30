package germany

import "fmt"

func (f *Form) ScanCountryData() {
	f.ScanBaseData()
	fmt.Println("Уровень немецкого языка:")
	fmt.Scan(&f.GermanLanguageLevel)
}
