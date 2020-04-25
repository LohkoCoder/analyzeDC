package main

import (
	"flag"
	"os"

	"fmt"
	"github.com/Luxurioust/excelize"
	"strconv"
	"strings"
	"time"
)

var timeMap = make(map[int]time.Time)
var transferMap = make(map[int]float64)
var dailyInterest float64 = 0.000658
var accumulatedInterest float64 = 0
var filePath string
var sheetName string
var deadline [3]int64

func main() {

	cmd()
	//创建excel文件
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	//cell := xlsx.GetCellValue("Sheet1", "B2")
	//fmt.Println(cell)

	// Get all the rows in the Sheet1.
	initData(xlsx, sheetName)
	analyze()
}

func initData(xlsx *excelize.File, sheetName string) {
	rows := xlsx.GetRows(sheetName)
	for index, row := range rows {

		timeMap[index] = strToDate(row[1])
		transferMap[index], _ = strconv.ParseFloat(row[2], 64)
	}
}

func strToDate(strTime string) time.Time {
	newTime := strings.Split(strTime, "-")
	year, _ := strconv.Atoi("20"+newTime[2])
	month, _ := strconv.Atoi(newTime[0])
	newMon := time.January
	switch month {
	case 1:
		newMon = time.January
	case 2:
		newMon = time.February
	case 3:
		newMon = time.March
	case 4:
		newMon = time.April
	case 5:
		newMon = time.May
	case 6:
		newMon = time.June
	case 7:
		newMon = time.July
	case 8:
		newMon = time.August
	case 9:
		newMon = time.September
	case 10:
		newMon = time.October
	case 11:
		newMon = time.November
	case 12:
		newMon = time.December
	}
	day, _ := strconv.Atoi(newTime[1])
	the_time := time.Date(year, newMon, day, 0, 0, 0, 0, time.UTC)
	return the_time
}

func analyze(){
	debet := transferMap[0]

	for i := 1; i < len(transferMap); i++ {
		day := int(timeMap[i].Sub(timeMap[i-1]).Hours()/24)
		transfer := transferMap[i]
		if transfer < 0 {
			accumulatedInterest += Round2(dailyInterest * float64(day) * debet + transfer)
			if accumulatedInterest < 0 {
				debet = Round2(debet + accumulatedInterest)
				accumulatedInterest = 0
			}
		} else {
			accumulatedInterest = Round2(accumulatedInterest + dailyInterest * float64(day) * debet)
			debet = Round2(debet + transfer)
		}
		fmt.Printf("日期: %s 累计未还利息: %g 累计未还本金: %g\n", timeMap[i], accumulatedInterest, debet)

	}
	fmt.Println("最终应还本金: ", debet)
	deadline1 := time.Date(int(deadline[0]), time.Month(deadline[1]), int(deadline[2]), 0, 0, 0, 0, time.UTC)
	deadlineGap := int(deadline1.Sub(timeMap[len(timeMap)-1]).Hours()/24)
	accumulatedInterest = Round2(accumulatedInterest + dailyInterest * float64(deadlineGap) * debet)
	fmt.Println("到2019年6月17日累计未还利息: ", accumulatedInterest)
	fmt.Println("到2019年6月17日累计未还本息和: ", debet+accumulatedInterest)
}

func Round2(value float64) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(2)+"f", value)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

func cmd() {
	var di string
	var fp string
	var sn string
	var dl string
	flag.StringVar(&fp, "f", "", "文件路径")
	flag.StringVar(&di, "di", "", "日利息默认为0.0658%")
	flag.StringVar(&sn, "sn", "", "表格名")
	flag.StringVar(&dl, "dl", "", "计息日期")
	flag.Parse()
	filePath = fp
	dailyInterest1, err := strconv.ParseFloat(di, 64)
	dailyInterest = dailyInterest1
	if err != nil {
		fmt.Println("dailyInterest parse error", err)
		os.Exit(400)
	}
	sheetName = sn
	extractDate(dl)
}

func extractDate(date string) {
	date1 := strings.Split(date, ".")
	for i := 0; i<3; i++ {
		deadline[i], _ = strconv.ParseInt(date1[i], 10, 64)
	}
}