package australia

import "fmt"

func (f *Form) ScanCountryData() {
	f.ScanBaseData()
	fmt.Println("Аллергические реакции:")
	fmt.Scan(&f.AllergicReactions)
}
