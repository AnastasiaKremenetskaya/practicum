package visa_forms

import "fmt"

const FinalText = "Заявление принято. " +
	"Срок обработки заявления в страну %s составляет %d дней"

type Base struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func (f *Base) ScanBaseData() {
	fmt.Println("Введите имя:")
	fmt.Scan(&f.Name)
	fmt.Println("Введите фамилию:")
	fmt.Scan(&f.Surname)
	fmt.Println("Введите отчество:")
	fmt.Scan(&f.Patronymic)
}
