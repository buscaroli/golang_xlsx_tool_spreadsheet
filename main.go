package main

func main() {
	stores := populateSheet("./sheet.xlsx")

	printUniqueStoreNames(stores)

	printTotalMileage(stores)

	printStatistics(stores)

	printStatisticsWithTemplates(stores)
}
