package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

type Data struct {
	State    int    `json:"state"`
	Message  string `json:"message"`
	Total    int    `json:"total"`
	PageNum  int    `json:"pageNum"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Tflag    int    `json:"Tflag"`
	Result   []struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		DetailsLink string `json:"detailsLink"`
		VideoLink   string `json:"videoLink"`
		Date        string `json:"date"`
		Week        string `json:"week"`
		Red         string `json:"red"`
		Blue        string `json:"blue"`
		Blue2       string `json:"blue2"`
		Sales       string `json:"sales"`
		Poolmoney   string `json:"poolmoney"`
		Content     string `json:"content"`
		Addmoney    string `json:"addmoney"`
		Addmoney2   string `json:"addmoney2"`
		Msg         string `json:"msg"`
		Z2Add       string `json:"z2add"`
		M2Add       string `json:"m2add"`
		Prizegrades []struct {
			Type      int    `json:"type"`
			Typenum   string `json:"typenum"`
			Typemoney string `json:"typemoney"`
		} `json:"prizegrades"`
		Zj1 string `json:"zj1,omitempty"`
		Mj1 string `json:"mj1,omitempty"`
		Zj6 string `json:"zj6,omitempty"`
		Mj6 string `json:"mj6,omitempty"`
	} `json:"result"`
}

func main() {
	url := "http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=&issueStart=&issueEnd=&dayStart=&dayEnd=&pageNo=1&pageSize=1555&week=&systemType=PC"
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"),
	)

	c.OnResponse(func(r *colly.Response) {
		r.Save("./data.json")

	})

	err := c.Visit(url)
	if err != nil {
		return
	}

	file, err := os.ReadFile("./data.json")
	if err != nil {
		log.Println(err)
		return
	}
	var data Data
	err = json.Unmarshal(file, &data)
	if err != nil {
		return
	}
	fmt.Println(data.Result)

}
