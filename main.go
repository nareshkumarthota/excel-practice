package main

import (
	"fmt"
	"sync"

	"github.com/tealeg/xlsx"
)

var totalBenValues map[string][]string
var totalValueSet map[string]string
var mtx sync.Mutex
var mdk []string
var sgd []string

type data struct {
	mandal     string
	totalHabs  int
	gpCount    int
	benCount   int
	tallyCount int
	diff       int
}

func main() {

	mdk = []string{"Shankarampet(A)", "HAVELI GHANPUR", "TEKMAL", "Medak", "Papannapet"}

	// excelFileName := "/home/naresh/Desktop/Docs/MissionBhagiratha_GP_Resolution_Report.xlsx"
	excelFileName := "./MB_Hab_Wise_Benficiaries_Report.xlsx"
	tempBenf := evaluate(excelFileName)
	excelFileName = "./MissionBhagiratha_GP_Resolution_Report.xlsx"
	tempGP := evaluate(excelFileName)

	totalArry := make([]data, 0)

	totalValueSet = make(map[string]string)
	for k, v := range tempGP {

		var talliedCount int

		for k1 := range tempBenf {
			if k == k1 {
				for _, val1 := range tempGP[k] {
					for _, val2 := range tempBenf[k1] {
						if val1 == val2 {
							talliedCount++
						}
					}
				}
			}
		}
		dataVal := data{}

		dataVal.mandal = k
		dataVal.gpCount = len(v)
		dataVal.benCount = len(tempBenf[k])
		dataVal.tallyCount = talliedCount
		dataVal.diff = dataVal.totalHabs - dataVal.tallyCount

		totalArry = append(totalArry, dataVal)

		for _, val := range mdk {
			if val == k {
				fmt.Printf("%20v", k)
				tgp := fmt.Sprintf("TOTAL-GP: %v", len(v))
				fmt.Printf("%20v", tgp)
				tbn := fmt.Sprintf("TOTAL-BN: %v", len(tempBenf[k]))
				fmt.Printf("%20v", tbn)
				tlc := fmt.Sprintf("TallyCount: %v", talliedCount)
				fmt.Printf("%20v", tlc)
				fmt.Printf("\n")
			}
		}

	}
}

func evaluate(filename string) map[string][]string {
	totalBenValues = nil

	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Println("ERROR", err)
	}

	totalBenValues = make(map[string][]string)
	temp := make([]string, 0)
	keyExist := false

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			var key, value string
			for c, cell := range row.Cells {

				if c == 0 {
					key = cell.String()
				} else if c == 1 {
					value = cell.String()
				}
			}

			for k1 := range totalBenValues {
				if k1 == key {
					keyExist = true
					break
				} else {
					keyExist = false
				}
			}
			if !keyExist {
				temp = nil
				temp = make([]string, 0)
			}
			temp = append(temp, value)
			totalBenValues[key] = temp

		}
	}

	for k1, v1 := range totalBenValues {
		totalBenValues[k1] = unique(v1)
	}

	return totalBenValues
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
