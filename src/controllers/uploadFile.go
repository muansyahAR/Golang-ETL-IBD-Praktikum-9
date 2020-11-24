package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	f "src/lib"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// unique cari unik
func unique(dtSlice []interface{}) []interface{} {
	keys := make(map[interface{}]bool)
	var list []interface{}
	for _, entry := range dtSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// UploadFile untuk unggah dokumen
func UploadFile(c *fiber.Ctx) error {
	//s := []string{}
	file, err := c.FormFile("document")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var extension = filepath.Ext(file.Filename)
	log.Print(file.Filename)
	log.Println(extension)

	// Save file to root directory:
	c.SaveFile(file, fmt.Sprintf("./temp/"+file.Filename))
	parsedData := f.ExcelCsvParser("./temp/"+file.Filename, extension)
	for i := 0; i < len(parsedData); i++ {
		delete(parsedData[i], "")
	}
	parsedJSON, _ := json.Marshal(parsedData)

	fmt.Println(parsedData[0])
	println()
	var s []interface{}
	for i := 0; i < len(parsedData); i++ {
		var dt = parsedData[i]["Date & Time"]
		var str = dt.(string)
		res1 := strings.Split(str, " ")
		//fmt.Println(reflect.TypeOf(res1))
		s = append(s, res1[0])
	}
	uniqueSlice := unique(s)
	fmt.Println(uniqueSlice)
	fmt.Print(len(uniqueSlice))
	var arrGroup []interface{}
	var cr string
	for x := 0; x < len(uniqueSlice); x++ {
		if uniqueSlice[x] != "" {
			//var arrTemp []float64
			var nTemp = 0.0
			for i := 0; i < len(parsedData); i++ {
				var dt = parsedData[i]["Date & Time"]
				var str = dt.(string)
				res1 := strings.Split(str, " ")
				if uniqueSlice[x] == res1[0] {
					cr = parsedData[i]["Credit"].(string)
					cr = strings.Replace(cr, ",", "", -1)
					cr, err := strconv.ParseFloat(cr, 64)
					if err != nil {
						fmt.Println(cr)
					}
					nTemp += cr
				} else {
					continue
				}

			}
			arrGroup = append(arrGroup, nTemp)
		}
	}
	fmt.Printf("%.4f\n", arrGroup)

	return c.SendString(string(parsedJSON))
}
