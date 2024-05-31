package main

import (
	"encoding/json"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"log"
	"strings"
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

type Result struct {
	NameCode    string `json:"nameCode" gorm:"primarykey"`
	Name        string `json:"name" `
	Code        string `json:"code"`
	DetailsLink string `json:"detailsLink"`
	VideoLink   string `json:"videoLink"`
	Date        string `json:"date"`
	Week        string `json:"week"`
	Red         string `json:"red"`
	R1          string `json:"r1"`
	R2          string `json:"r2"`
	R3          string `json:"r3"`
	R4          string `json:"r4"`
	R5          string `json:"r5"`
	R6          string `json:"r6"`
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
	Zj1         string `json:"zj1,omitempty"`
	Mj1         string `json:"mj1,omitempty"`
	Zj6         string `json:"zj6,omitempty"`
	Mj6         string `json:"mj6,omitempty"`
	Prizegrades string `json:"prizegrades"`
}

func main() {
	url := "http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=&issueStart=&issueEnd=&dayStart=&dayEnd=&pageNo=1&pageSize=1555&week=&systemType=PC"
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"),
	)
	c.OnError(func(r *colly.Response, err error) {
		log.Println(err)
	})

	c.OnResponse(func(r *colly.Response) {
		// 保存响应内容到 json文件

		r.Save("./data.json")

		// 解析json到对象
		var data Data
		err := json.Unmarshal(r.Body, &data)
		if err != nil {
			log.Println(err)
			return
		}

		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
		db.AutoMigrate(&Result{})
		for _, v := range data.Result {
			var (
				prizegrades string
				nameCode    string
				r1          string
				r2          string
				r3          string
				r4          string
				r5          string
				r6          string
			)
			r6 = ""
			nameCode = v.Name + "-" + v.Code
			for _, p := range v.Prizegrades {
				prizegrades += fmt.Sprintf("奖项: %d,数量: %s,金额: %s - ", p.Type, p.Typenum, p.Typemoney)

			}
			numbers := strings.Split(v.Red, ",")
			r1 = numbers[0]
			r2 = numbers[1]
			r3 = numbers[2]
			r4 = numbers[3]
			r5 = numbers[4]
			r6 = numbers[5]

			db.Create(&Result{
				NameCode:    nameCode,
				Name:        v.Name,
				Code:        v.Code,
				DetailsLink: v.DetailsLink,
				VideoLink:   v.VideoLink,
				Date:        v.Date,
				Week:        v.Week,
				Red:         v.Red,
				R1:          r1,
				R2:          r2,
				R3:          r3,
				R4:          r4,
				R5:          r5,
				R6:          r6,
				Blue:        v.Blue,
				Blue2:       v.Blue2,
				Sales:       v.Sales,
				Poolmoney:   v.Poolmoney,
				Content:     v.Content,
				Addmoney:    v.Addmoney,
				Addmoney2:   v.Addmoney2,
				Msg:         v.Msg,
				Z2Add:       v.Z2Add,
				M2Add:       v.M2Add,
				Zj1:         v.Zj1,
				Mj1:         v.Mj1,
				Zj6:         v.Zj6,
				Mj6:         v.Mj6,
				Prizegrades: prizegrades,
			})
		}

	})

	err := c.Visit(url)
	if err != nil {
		return
	}

}
