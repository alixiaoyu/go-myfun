package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/zuijinbuzai/excelclaim/excel"
)

type EventStatus struct {
	ID        string
	Name      string
	EventFlow string
	url       string
}

type ToExcel struct {
	Res []Result `json:"-,omitempty"`
}
type Result struct {
	StepNo    int    `json:"stepNo,omitempty"`
	Name      string `json:"name,omitempty"`
	StartTime string `json:"startTime,omitempty"`
	EndTime   string `json:"endTime,omitempty"`
	Status    int    `json:"status,omitempty"`
}

type InSlice []string

func main() {
	var result EventStatus
	var ex ToExcel
	db, _ := gorm.Open("mysql", "root:lxy123456@/test?charset=utf8&parseTime=True&loc=Local")
	err := db.Table("t_event_status").Where("id = ?", 1).First(&result).Error
	if err != nil {
		fmt.Println("数据库查询错误")
		return
	}
	err = json.Unmarshal([]byte(result.EventFlow), &ex.Res)
	if err != nil {
		fmt.Println("JSON序列化失败 error", err)
		return
	}
	var exSlice []InSlice
	for _, v := range ex.Res {
		slce := InSlice{}
		name := v.Name
		etime := v.EndTime
		stime := v.StartTime
		sNo := strconv.Itoa(v.StepNo)
		var status string
		if v.Status == 1 {
			status = "应急已处理"
		} else {
			status = "应急未处理"
		}
		slce = append(slce, sNo, name, stime, etime, status)
		exSlice = append(exSlice, slce)

	}
	// resp, err := CreateExcel("1", exSlice)
	resp, err := ExceGen(exSlice)
	fmt.Println("------>", resp, err)
}

func ExceGen(ins []InSlice) (string, error) {
	f := excelize.NewFile()

	sheet := excel.NewSheet(f, "应急时间", 5, 28)
	sheet.SetAllColsWidth(7, 50, 30, 30, 12)
	excelStyle2 := excel.NewExcelStyle(11, 0, true)
	sheet.WriteRow("编号", "名称", "开始时间", "结束时间", "处理状态")
	for _, v := range ins {
		if v[3] == "" {
			v[3] = "NULL"
		}
		sheet.WriteRow(v[0], v[1], v[2], v[3], v[4]).ApplyItem(1, excelStyle2)
	}

	err := f.SaveAs("./test.xlsx")
	if err != nil {
		return "failed", err
	}
	return "success", nil

}

func CreateExcel(id string, ins []InSlice) (string, error) {
	fname := fmt.Sprintf("%s.xls", time.Now().Format("20060102")+id)
	f, err := os.Create(fname)
	if err != nil {
		return "创建Excel失败", err
	}
	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	w.Write([]string{"编号", "名称", "开始时间", "结束时间", "处理状态"})
	for _, v := range ins {

		w.Write(v)
	}
	w.Flush()
	return "success", nil
}
