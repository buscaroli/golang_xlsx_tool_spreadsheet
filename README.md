# GoXLSX

A little cli program I have made in order to get some data out of a personal spreadsheet.
I have been asked to send the total mileage run during the previous financial year and I have decided to estrapolate the data programmatically using Go.

## How to use

```
go run main.go
```

The .xlxs file needs to be named sheet.xlsx and needs to have the store name inside the B column.
Add the store names and miles from home to your StoreCount and StoreDistance maps.
