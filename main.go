package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lottery-analysis/dal/model"
)

func main() {

	c := colly.NewCollector()
	c.OnResponse(func(r *colly.Response) {
		// 存储json到.json文件
		r.Save("./data/data.json")
		// 使用map[string]interface{}解析任意结构的JSON数据
		var data1 map[string]interface{}
		// 解组json->struct
		err := json.Unmarshal(r.Body, &data1)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		if _, exists := data1["result"]; exists {
			// 福彩
			var response model.FcResponse
			var fcResults []model.FcResult
			err := json.Unmarshal(r.Body, &response)
			if err != nil {
				return
			}
			// 遍历result
			for _, v := range response.Result {
				var fcPrizegrades []model.FcPrizegrade
				for _, p := range v.Prizegrades {
					fcPrizegrade := model.FcPrizegrade{
						Type:      p.Type,
						Typenum:   p.Typenum,
						Typemoney: p.Typemoney,
					}
					fcPrizegrades = append(fcPrizegrades, fcPrizegrade)
				}
				fcResult := model.FcResult{
					Name:        v.Name,
					Code:        v.Code,
					DetailsLink: v.DetailsLink,
					VideoLink:   v.VideoLink,
					Date:        v.Date,
					Week:        v.Week,
					Red:         v.Red,
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
					Prizegrades: fcPrizegrades,
				}
				fcResults = append(fcResults, fcResult)
			}
			// 保存数据到mysql
			dsn := "root:root123@tcp(127.0.0.1:3306)/fc?charset=utf8mb4&parseTime=True&loc=Local"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			// 自动迁移同步
			db.AutoMigrate(&model.FcResult{}, &model.FcPrizegrade{})

			// fmt.Println(fcResults)
			// 逐条插入
			for _, fcResult := range fcResults {
				// 开启事务
				err = db.Transaction(func(tx *gorm.DB) error {
					// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
					var maxID uint
					// COALESCE(MAX(id), 0) 用于在 MAX(id) 返回空值时将其替换为零。这样，即使表是新建的且没有任何记录，也不会出现空值的问题。
					if err := tx.Model(&model.FcResult{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID).Error; err != nil {

						return err
					}
					// 手动分配新的ID
					fcResult.ID = maxID + 1
					// 插入数据
					if err := tx.Create(&fcResult).Error; err != nil {
						// 返回任何错误都会回滚事务
						return err
					}
					// 返回 nil 提交事务
					return nil
				})
			}

		} else {
			// 大乐透
			// fmt.Println("================")
			// fmt.Println(data1["value"])
		}

		// for k, v := range data1 {
		// 	fmt.Println(k, v)
		// }

	})
	urlMap := map[string]string{
		"ssq": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&pageNo=1&pageSize=2000&systemType=PC",
		// "3d":  "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=3d&pageNo=1&pageSize=2000&systemType=PC",
		// "qlc": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=qlc&pageNo=1&pageSize=2000&systemType=PC",
		// "kl8": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=kl8&pageNo=1&pageSize=2000&systemType=PC",
		// "dlt": "https://webapi.sporttery.cn/gateway/lottery/getHistoryPageListV1.qry?gameNo=85&provinceId=0&pageSize=20&isVerify=1&pageNo=1",
	}
	for _, v := range urlMap {
		c.Visit(v)
	}

	// 等待所有访问完成
	c.Wait()
}
