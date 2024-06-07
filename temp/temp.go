package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"lottery-analysis/dal/model"
)

func main() {

	var data model.FcResponse
	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		println(r.StatusCode)
		println(r.Body)
		// 数据处理
		// r.Save("./data/data.json")

		json.Unmarshal(r.Body, &data)
		// fmt.Println(data.Result)
		for k, v := range data.Result {
			fmt.Println(k, v)
		}

	})
	c.Visit("https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&pageNo=1&pageSize=2000&systemType=PC")
	urlMap := map[string]string{
		"3d":  "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=3d&pageNo=1&pageSize=2000&systemType=PC",
		"ssq": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&pageNo=1&pageSize=2000&systemType=PC",
		"qlc": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=qlc&pageNo=1&pageSize=2000&systemType=PC",
		"kl8": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=kl8&pageNo=1&pageSize=2000&systemType=PC",
		"dlt": "https://webapi.sporttery.cn/gateway/lottery/getHistoryPageListV1.qry?gameNo=85&provinceId=0&pageSize=2000&isVerify=1&pageNo=1",
	}
	for k, v := range urlMap {
		c.Visit(v)
		fmt.Println(k, v)
	}

}
