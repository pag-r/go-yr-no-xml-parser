package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Forecasts struct {
	XMLName   xml.Name   `xml:"weatherdata"`
	Forecasts []Forecast `xml:"forecast>tabular>time"`
}

type Forecast struct {
	XMLName  xml.Name `xml:"time"`
	TimeFrom string   `xml:"from,attr"`
	TimeTo   string   `xml:"to,attr"`
	Icon     Icon     `xml:"symbol"`
}

type Icon struct {
	XMLName xml.Name `xml:"symbol"`
	Name    string   `xml:"name,attr"`
}

func panic(err error) {
	if err != nil {
		fmt.Printf("%s\n", err)
	}
}

func main() {

	var xmlFilePath string
	if len(os.Args) > 1 {
		xmlFilePath = os.Args[1]
		file, err := os.Open(xmlFilePath)
		panic(err)
		defer file.Close()

		b, err := ioutil.ReadAll(file)
		panic(err)

		var forecasts Forecasts
		err = xml.Unmarshal([]byte(b), &forecasts)
		panic(err)

		for i := range forecasts.Forecasts {
			now := time.Now().Format(time.RFC3339)
			from := forecasts.Forecasts[i].TimeFrom
			to := forecasts.Forecasts[i].TimeTo
			if now > from && now < to {
				fmt.Printf("%s > %s < %s - %s\n",
					from, now, to, forecasts.Forecasts[i].Icon.Name)
			}
		}
	} else {
		panic(fmt.Errorf("XML file is required"))
	}
}
