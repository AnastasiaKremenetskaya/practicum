package egypt

import "fmt"

func (f *Form) ScanCountryData() {
	f.ScanBaseData()
	fmt.Println("Место работы:")
	fmt.Scan(&f.Job)
	fmt.Println("Название отеля, где планируется проживание:")
	fmt.Scan(&f.Hotel)
}
