package main

import (
	"fmt"
	"sync"

	"github.com/tealeg/xlsx"
)

var totalBenValues map[string][]string
var totalValueSet map[string]string
var mtx sync.Mutex

type mndlHabs struct {
	mandalName string
	totalHabs  int
}

var sangareddyDist map[string][]mndlHabs
var mdkDist map[string][]mndlHabs
var totalDist map[string]map[string][]mndlHabs

type mandlCount struct {
	name       string
	totalhbs   int
	totalgp    int
	totalbn    int
	tallycount int
	diff       int
}

func main() {

	sangareddyDist = make(map[string][]mndlHabs)
	sangareddyDist["Andole SD"] = []mndlHabs{mndlHabs{mandalName: "Andole", totalHabs: 30}, mndlHabs{mandalName: "Pulkal", totalHabs: 45}, mndlHabs{mandalName: "Wattpally", totalHabs: 28}}
	sangareddyDist["Hathnoora SD"] = []mndlHabs{mndlHabs{mandalName: "Ameenpur", totalHabs: 5}, mndlHabs{mandalName: "Gumadidala", totalHabs: 21}, mndlHabs{mandalName: "Hathnoora", totalHabs: 50}, mndlHabs{mandalName: "Jinnaram", totalHabs: 30}, mndlHabs{mandalName: "Patancheru", totalHabs: 24}, mndlHabs{mandalName: "RC Puram", totalHabs: 7}}
	sangareddyDist["Sangareddy SD"] = []mndlHabs{mndlHabs{mandalName: "Kondapur", totalHabs: 38}, mndlHabs{mandalName: "Kandi", totalHabs: 31}, mndlHabs{mandalName: "Sadasivapet", totalHabs: 38}, mndlHabs{mandalName: "Munipally", totalHabs: 31}, mndlHabs{mandalName: "Sangareddy", totalHabs: 23}}
	sangareddyDist["Narayankhed SD"] = []mndlHabs{mndlHabs{mandalName: "Sirgapur", totalHabs: 49}, mndlHabs{mandalName: "Nagalgidda", totalHabs: 62}, mndlHabs{mandalName: "Manoor", totalHabs: 33}, mndlHabs{mandalName: "Kalher", totalHabs: 49}, mndlHabs{mandalName: "Kangti", totalHabs: 52}, mndlHabs{mandalName: "Nagalagidda", totalHabs: 62}, mndlHabs{mandalName: "Narayankhed", totalHabs: 88}}
	sangareddyDist["Zaheerabad SD"] = []mndlHabs{mndlHabs{mandalName: "Nyalkal", totalHabs: 46}, mndlHabs{mandalName: "Jarasangam", totalHabs: 39}, mndlHabs{mandalName: "MOGUDAMPALLY", totalHabs: 29}, mndlHabs{mandalName: "Raikode", totalHabs: 35}, mndlHabs{mandalName: "Koheer", totalHabs: 30}, mndlHabs{mandalName: "Zaheerabad", totalHabs: 37}}

	mdkDist = make(map[string][]mndlHabs)
	mdkDist["Medak SD"] = []mndlHabs{mndlHabs{mandalName: "Alladurg", totalHabs: 28}, mndlHabs{mandalName: "REGODE", totalHabs: 26}, mndlHabs{mandalName: "Shankarampet(A)", totalHabs: 41}, mndlHabs{mandalName: "HAVELI GHANPUR", totalHabs: 63}, mndlHabs{mandalName: "TEKMAL", totalHabs: 64}, mndlHabs{mandalName: "Medak", totalHabs: 44}, mndlHabs{mandalName: "Papannapet", totalHabs: 58}}
	mdkDist["Toopran SD"] = []mndlHabs{mndlHabs{mandalName: "Narsingi", totalHabs: 13}, mndlHabs{mandalName: "NIZAMPET", totalHabs: 23}, mndlHabs{mandalName: "Ramayampet", totalHabs: 38}, mndlHabs{mandalName: "Shankarampet-R", totalHabs: 49}, mndlHabs{mandalName: "Toopran", totalHabs: 35}, mndlHabs{mandalName: "Chegunta", totalHabs: 48}, mndlHabs{mandalName: "Manoharabad", totalHabs: 19}}
	mdkDist["Narsapur SD"] = []mndlHabs{mndlHabs{mandalName: "CHILIPICHEDU", totalHabs: 41}, mndlHabs{mandalName: "Shivampet", totalHabs: 79}, mndlHabs{mandalName: "Yeldurthy", totalHabs: 52}, mndlHabs{mandalName: "Kowdipally", totalHabs: 97}, mndlHabs{mandalName: "Kulcharam", totalHabs: 43}, mndlHabs{mandalName: "Narsapur", totalHabs: 98}}

	totalDist = make(map[string]map[string][]mndlHabs)
	totalDist["Medak"] = mdkDist
	totalDist["Sangareddy"] = sangareddyDist

	// excelFileName := "/home/naresh/Desktop/Docs/MissionBhagiratha_GP_Resolution_Report.xlsx"
	excelFileName := "./MB_Hab_Wise_Benficiaries_Report.xlsx"
	tempBenf := evaluate(excelFileName)
	excelFileName = "./MissionBhagiratha_GP_Resolution_Report.xlsx"
	tempGP := evaluate(excelFileName)

	totalValueSet = make(map[string]string)

	totalMndls := make([]mandlCount, 0)

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

		tempMndlCount := mandlCount{}

		tempMndlCount.name = k
		tempMndlCount.totalgp = len(v)
		tempMndlCount.totalbn = len(tempBenf[k])
		tempMndlCount.tallycount = talliedCount

		totalMndls = append(totalMndls, tempMndlCount)

		// fmt.Printf("%20v", k)
		// tgp := fmt.Sprintf("TOTAL-GP: %v", len(v))
		// fmt.Printf("%20v", tgp)
		// tbn := fmt.Sprintf("TOTAL-BN: %v", len(tempBenf[k]))
		// fmt.Printf("%20v", tbn)
		// tlc := fmt.Sprintf("TallyCount: %v", talliedCount)
		// fmt.Printf("%20v", tlc)
		// fmt.Printf("\n")
	}

	for dist, sdval := range totalDist {

		var hbs, gps, bns, tallys, diffs int
		for k, v := range sdval {
			var totalhbs, totalgps, totalbns, totaltally, totaldiff int
			for _, v1 := range v {

				for _, vv1 := range totalMndls {
					if v1.mandalName == vv1.name {
						vv1.totalhbs = v1.totalHabs
						vv1.diff = vv1.totalhbs - vv1.tallycount
						totalhbs = totalhbs + vv1.totalhbs
						totalgps = totalgps + vv1.totalgp
						totalbns = totalbns + vv1.totalbn
						totaltally = totaltally + vv1.tallycount
						totaldiff = totaldiff + vv1.diff

						fmt.Printf("%20v", vv1.name)
						fmt.Printf("%5v", vv1.totalhbs)
						fmt.Printf("%5v", vv1.totalgp)
						fmt.Printf("%5v", vv1.totalbn)
						fmt.Printf("%5v", vv1.tallycount)
						fmt.Printf("%5v", vv1.diff)
						fmt.Println()
					}
				}

			}
			fmt.Printf("%20v", k)
			fmt.Printf("%5v", totalhbs)
			fmt.Printf("%5v", totalgps)
			fmt.Printf("%5v", totalbns)
			fmt.Printf("%5v", totaltally)
			fmt.Printf("%5v", totaldiff)
			fmt.Println()
			fmt.Println()

			hbs = totalhbs + hbs
			gps = totalgps + gps
			bns = totalbns + bns
			tallys = totaltally + tallys
			diffs = totaldiff + diffs

		}
		fmt.Printf("%20v", dist)
		fmt.Printf("%5v", hbs)
		fmt.Printf("%5v", gps)
		fmt.Printf("%5v", bns)
		fmt.Printf("%5v", tallys)
		fmt.Printf("%5v", diffs)
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
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

	for dist := range totalDist {

		for _, sheet := range xlFile.Sheets {
			for _, row := range sheet.Rows {
				var key, value, distVal string
				for c, cell := range row.Cells {
					if c == 0 {
						distVal = cell.String()
					} else if c == 1 {
						key = cell.String()
					} else if c == 2 {
						value = cell.String()
					}
				}

				if distVal == dist {
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
