package germany

import (
	"encoding/json"
	"fmt"
)

func (f *Form) PrintForm() {
	res, _ := json.Marshal(f)
	fmt.Printf("Ваше заявление: " + string(res))
}
