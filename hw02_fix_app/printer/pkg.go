package printer

import (
	"fmt"

	"github.com/DmitryFPS/go/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) {
	var str string
	for i := 0; i < len(staff); i++ {
		id := staff[i].UserID
		age := staff[i].Age
		name := staff[i].Name
		departmentID := staff[i].DepartmentID
		str = fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ", id, age, name, departmentID)
		fmt.Println(str)
	}

	fmt.Println(str)
}
