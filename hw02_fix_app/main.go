package main

import (
	"fmt"

	"github.com/DmitryFPS/go/hw02_fix_app/printer"
	"github.com/DmitryFPS/go/hw02_fix_app/reader"
	"github.com/DmitryFPS/go/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")
	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Println("Error", &path)
	}

	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
