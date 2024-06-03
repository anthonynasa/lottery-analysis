package main

import (
	"encoding/json"
	"fmt"
	"github.com/glebarez/sqlite"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"log"
	"lottery-analysis/dal/model"
	"strings"
)

func main() {
	// 收集器
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
			"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36"),
		// colly.Async(true),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		// 保存响应内容到 json文件
		r.Save("./data/data.json")

		// 保存数据到数据库
		DoData(r)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error visiting", r.Request.URL, ":", err)
	})

	// 开始访问
	urls := []string{
		"https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&pageNo=1&pageSize=2000&systemType=PC",
		"https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=3d&pageNo=1&pageSize=2000&systemType=PC",
		"https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=qlc&pageNo=1&pageSize=2000&systemType=PC",
		// "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=kl8&pageNo=1&pageSize=2000&systemType=PC",
	}

	for _, url := range urls {
		fmt.Println(url)
		// 访问url
		err := c.Visit(url)
		if err != nil {

		}
	}
	// 等待所有访问完成
	c.Wait()

}

// DoData
//
//	@Description: 保存数据到数据库
//	@param r
func DoData(r *colly.Response) {
	// 解析json到对象
	var data model.Data
	err := json.Unmarshal(r.Body, &data)
	if err != nil {
		log.Println(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	db.AutoMigrate(&model.Result{})
	for _, v := range data.Result {
		nameCode := v.Name + "-" + v.Code
		prizegrades, r1, r2, r3, r4, r5, r6, r7 := "", "", "", "", "", "", "", ""

		for _, p := range v.Prizegrades {
			prizegrades += fmt.Sprintf("奖项: %d,数量: %s,金额: %s - ", p.Type, p.Typenum, p.Typemoney)

		}
		numbers := strings.Split(v.Red, ",")
		for i, strNum := range numbers {
			// 根据索引分别赋值给对应的变量
			switch i {
			case 0:
				r1 = strNum
			case 1:
				r2 = strNum
			case 2:
				r3 = strNum
			case 3:
				r4 = strNum
			case 4:
				r5 = strNum
			case 5:
				r6 = strNum
			case 7:
				r7 = strNum
			}
		}

		db.Create(&model.Result{
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
			R7:          r7,
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
}
