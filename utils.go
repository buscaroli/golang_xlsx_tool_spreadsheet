package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/tealeg/xlsx"
)

// populateSheet extracts the data from the xlsx file and returns it
func populateSheet(filename string) []string {
	names := make([]string, 0)

	excelFileName := filename
	xlFile, err := xlsx.OpenFile(excelFileName)

	if err != nil {
		fmt.Printf("Error reading file.")
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for index, cell := range row.Cells {
				if index == 1 && strings.TrimSpace(cell.String()) != "" && strings.TrimSpace(cell.String()) != "Store Name" {
					names = append(names, strings.TrimSpace(cell.String()))
				}
			}
		}
	}

	return names
}

// printUniqueStoreNames prints a list of all the stores visited by the user
func printUniqueStoreNames(sliceOfNames []string) {
	// I have decided to keep a map of the stores in case in a future version I decide to return that value instaead of printing it
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range sliceOfNames {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	fmt.Print("\n~~ Unique Store Names ~~\n")
	for i, store := range list {
		fmt.Printf("%d - %s\n", i, store)
	}
	fmt.Print("\n~~  ~~  ~~  ~~  ~~\n\n")
}

// Calculates the total mileage of the shifts
func getTotalMileage(storeNames []string) float64 {
	total := 0.0

	for _, store := range storeNames {
		total += StoreDistance[store]
	}

	// return travel
	return total * 2
}

// Prints the total mileage of the shifts
func printTotalMileage(stores []string) {
	fmt.Printf("\n\n~~ Total Mileage ~~\n %.2f\n\n~~  ~~  ~~  ~~  ~~\n\n", getTotalMileage(stores))
}

// Prints some statistics about the shifts
func printStatistics(stores []string) {
	total := 0.0

	for _, store := range stores {
		StoreCount[store] += 1
		total += 1
	}

	fmt.Print("\n\n~~ Times worked at each store and %\n\n")

	for name, times := range StoreCount {
		var percent = float64(times) / total * 100
		fmt.Printf("%s\t\t%d\t\t%.2f percent\n", name, times, percent)
	}

}

// Prints statistical data with a nicer formatting due to go templates
func printStatisticsWithTemplates(stores []string) {
	total := 0.0

	for _, store := range stores {
		StoreCount[store] += 1
		total += 1
	}

	fmt.Print("\n\n~~ Times worked at each store and %\n\n")

	templateString := "Name: {{.Name}}\n   Worked: {{.Times}} times\n   Percent: {{.Percent}}%\n\n"
	tmpl, err := template.New("statistics").Parse(templateString)
	if err != nil {
		log.Fatal("Error parsing file: ", err)
	}

	for name, times := range StoreCount {
		var percent = float64(times) / total * 100
		// fmt.Printf("%s\t\t%d\t\t%.2f percent\n", name, times, percent)
		stringTimes := fmt.Sprint(times)
		stringPercent := fmt.Sprintf("%.2f", percent)

		current := Shift{name, stringTimes, stringPercent}
		err = tmpl.Execute(os.Stdout, current)
		if err != nil {
			log.Fatal("Error executing template: ", err)
		}
	}

}

type Shift struct {
	Name    string
	Times   string
	Percent string
}
