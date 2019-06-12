package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty"
)

func main() {
	// ip := "11.124.25.3"
	username := "visionadmin"
	password := "MRTDB@zy"

	resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// https://11.124.25.3/piwebapi/points/multiple?path=\\ZY-MRTDB\ZYX00D003PT1207

	res, err := resty.R().
		SetBasicAuth(username, password).
		Get("https://11.124.25.3/piwebapi/points/multiple?path=" + "\\\\" + "ZY-MRTDB" + "\\ZYX00C055FQQNSUM")

	piResult := WebID{}

	fmt.Println("***", string(res.Body()))
	time.Sleep(time.Second * 10)
	err = json.Unmarshal(res.Body(), &piResult)
	if err != nil {
		time.Sleep(time.Second * 5)
		fmt.Println("getStreamsValue请求解序列化失败")
		return
	}
	fmt.Println("---->", piResult.Items[0].Object.WebID)
	fmt.Println("success")
	time.Sleep(time.Second * 10)

}

func test01() {
	fmt.Println("---->,dev01分支下提交代码")
}

type WebID struct {
	Links struct {
	} `json:"Links"`
	Items []struct {
		Identifier     string `json:"Identifier"`
		IdentifierType string `json:"IdentifierType"`
		Object         struct {
			WebID            string  `json:"WebId"`
			ID               int     `json:"Id"`
			Name             string  `json:"Name"`
			Path             string  `json:"Path"`
			Descriptor       string  `json:"Descriptor"`
			PointClass       string  `json:"PointClass"`
			PointType        string  `json:"PointType"`
			DigitalSetName   string  `json:"DigitalSetName"`
			EngineeringUnits string  `json:"EngineeringUnits"`
			Span             float64 `json:"Span"`
			Zero             float64 `json:"Zero"`
			Step             bool    `json:"Step"`
			Future           bool    `json:"Future"`
			DisplayDigits    int     `json:"DisplayDigits"`
			Links            struct {
				Self             string `json:"Self"`
				DataServer       string `json:"DataServer"`
				Attributes       string `json:"Attributes"`
				InterpolatedData string `json:"InterpolatedData"`
				RecordedData     string `json:"RecordedData"`
				PlotData         string `json:"PlotData"`
				SummaryData      string `json:"SummaryData"`
				Value            string `json:"Value"`
				EndValue         string `json:"EndValue"`
			} `json:"Links"`
		} `json:"Object"`
	} `json:"Items"`
}
