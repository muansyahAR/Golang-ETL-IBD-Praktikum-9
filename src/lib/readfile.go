package lib

import (
	"bufio"
	"container/list"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"

	"github.com/dineshsonachalam/CSV-and-Excel-data-to-JSON/parser"
)

// ReadCsvFile baca file csv
func ReadCsvFile(filePath string) []map[string]interface{} {
	// Load a csv file.
	f, _ := os.Open(filePath)
	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	result, _ := r.ReadAll()
	parsedData := make([]map[string]interface{}, 0, 0)
	header_name := result[0]

	for row_counter, row := range result {

		if row_counter != 0 {
			var singleMap = make(map[string]interface{})
			for col_counter, col := range row {
				singleMap[header_name[col_counter]] = col
			}
			if len(singleMap) > 0 {

				parsedData = append(parsedData, singleMap)
			}
		}
	}
	fmt.Println("Length of parsedData:", len(parsedData))
	return parsedData

}

// ReadXlsxFile untuk baca Xlsx
func ReadXlsxFile(filePath string) []map[string]interface{} {
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		fmt.Println("Error reading the file")
	}

	parsedData := make([]map[string]interface{}, 0, 0)
	header_name := list.New()
	// sheet
	for _, sheet := range xlFile.Sheets {
		// rows
		for row_counter, row := range sheet.Rows {

			// column
			header_iterator := header_name.Front()
			var singleMap = make(map[string]interface{})

			for _, cell := range row.Cells {
				if row_counter == 0 {
					text := cell.String()
					header_name.PushBack(text)
				} else {
					text := cell.String()
					singleMap[header_iterator.Value.(string)] = text
					header_iterator = header_iterator.Next()
				}
			}
			if row_counter != 0 && len(singleMap) > 0 {

				parsedData = append(parsedData, singleMap)
			}

		}
	}
	fmt.Println("Length of parsedData:", len(parsedData))
	return parsedData
}

// ReadXlsFile untuk baca Xls
func ReadXlsFile(filePath string) []map[string]interface{} {
	parsedData := make([]map[string]interface{}, 0, 0)
	println(filePath)
	if xlFile, err := xls.Open(filePath, "utf-8"); err == nil {
		total_sheets := xlFile.NumSheets()
		for sheetCounter := 0; sheetCounter < total_sheets; sheetCounter++ {
			if sheet := xlFile.GetSheet(sheetCounter); sheet != nil {
				header_name := list.New()
				for rowCounter := 0; rowCounter <= (int(sheet.MaxRow)); rowCounter++ {
					row := sheet.Row(rowCounter)
					header_iterator := header_name.Front()
					var singleMap = make(map[string]interface{})
					for colCounter := 0; colCounter < (int(row.LastCol())); colCounter++ {
						if rowCounter == 0 {
							text := row.Col(colCounter)
							header_name.PushBack(text)
						} else {
							text := row.Col(colCounter)
							singleMap[header_iterator.Value.(string)] = text
							header_iterator = header_iterator.Next()
						}
					}
					if rowCounter != 0 && len(singleMap) > 0 {
						parsedData = append(parsedData, singleMap)
					}
				}
			}
		}
	}
	fmt.Println("Length of parsedData:", len(parsedData))
	return parsedData
}

// ExcelCsvParser untuk parsing
func ExcelCsvParser(blobPath string, blobExtension string) (parsedData []map[string]interface{}) {
	if blobExtension == ".csv" {
		fmt.Println("-------We are parsing an csv file.-------------")
		parsedData := parser.ReadCsvFile(blobPath)
		fmt.Printf("Type:%T\n", parsedData)
		return parsedData

	} else if blobExtension == ".xlsx" {
		fmt.Println("----------------We are parsing an xlsx file.---------------")
		parsedData := parser.ReadXlsxFile(blobPath)
		return parsedData
	} else if blobExtension == ".xls" {
		fmt.Println("----------------We are parsing an xls file.---------------")
		parsedData := parser.ReadXlsFile(blobPath)
		return parsedData
	}
	return parsedData
}
