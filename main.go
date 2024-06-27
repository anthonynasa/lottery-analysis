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

	urlMap := map[string]string{
		"ssq": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&pageNo=1&pageSize=2000&systemType=PC",
		// "3d":  "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=3d&pageNo=1&pageSize=2000&systemType=PC",
		// "qlc": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=qlc&pageNo=1&pageSize=2000&systemType=PC",
		// "kl8": "https://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=kl8&pageNo=1&pageSize=2000&systemType=PC",
		"dlt": "https://webapi.sporttery.cn/gateway/lottery/getHistoryPageListV1.qry?gameNo=85&pageSize=200&isVerify=1&pageNo=1",
	}
	for _, v := range urlMap {

		doCollect( v)
	}
}

// 封装访问
func doCollect(url string) {
	c := colly.NewCollector(
	// colly.Async(true),
	)
	c.OnResponse(func(r *colly.Response) {
		// 存储json到.json文件
		r.Save("./data/data.json")
		// 使用map[string]interface{}解析任意结构的JSON数据
		var data map[string]interface{}
		// 解组json->struct
		err := json.Unmarshal(r.Body, &data)
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err)
			return
		}

		if _, exists := data["result"]; exists {
			// 福彩
			var fcResponse model.FcResponse
			var fcResults []model.FcResult
			err := json.Unmarshal(r.Body, &fcResponse)
			if err != nil {
				fmt.Println("fc解组json出错:", err.Error())
				return
			}
			// 遍历result
			for _, v := range fcResponse.Result {
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
			// 1. 创建并连接到数据库
			dsn := "root:root123@tcp(127.0.0.1:3306)/"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				fmt.Println("Failed to connect to database:", err)
				return
			}
			// 2. 创建数据库
			dbName := "fc"
			err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
			if err != nil {
				fmt.Println("Failed to create database:", err)
				return
			}
			dsn = fmt.Sprintf("root:root123@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)

			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

			// 3. 自动迁移同步
			db.AutoMigrate(&model.FcResult{}, &model.FcPrizegrade{})

			// 4. 逐条插入
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
			var tcResponse model.TcResponse
			var tcResults []model.TcResult

			err := json.Unmarshal(r.Body, &tcResponse)
			if err != nil {
				fmt.Println("dlt解析出错:", err.Error())
				return
			}
			// fmt.Println(tcResponse)
			for _, v := range tcResponse.Value.List {
				var tcPrizeLevelList []model.TcPrizeLevel
				for _, p := range v.PrizeLevelList {
					tcPrizeLevel := model.TcPrizeLevel{
						AwardType:         p.AwardType,
						Group:             p.Group,
						LotteryCondition:  p.LotteryCondition,
						PrizeLevel:        p.PrizeLevel,
						Sort:              p.Sort,
						StakeAmount:       p.StakeAmount,
						StakeAmountFormat: p.StakeAmountFormat,
						StakeCount:        p.StakeCount,
						TotalPrizeamount:  p.TotalPrizeamount,
					}
					tcPrizeLevelList = append(tcPrizeLevelList, tcPrizeLevel)
				}
				tcResult := model.TcResult{
					LotteryGameName:       v.LotteryGameName,
					LotteryGameNum:        v.LotteryGameNum,
					LotteryDrawNum:        v.LotteryDrawNum,
					DrawPdfUrl:            v.DrawPdfUrl,
					LotteryDrawTime:       v.LotteryDrawTime,
					LotteryEquipmentCount: v.LotteryEquipmentCount,
					LotteryDrawResult:     v.LotteryDrawResult,
					TotalSaleAmount:       v.TotalSaleAmount,
					PoolBalanceAfterdraw:  v.PoolBalanceAfterdraw,
					PrizeLevelList:        tcPrizeLevelList,
				}
				tcResults = append(tcResults, tcResult)

			}

			// 保存数据到mysql
			// 1. 创建并连接到数据库
			dsn := "root:root123@tcp(127.0.0.1:3306)/"
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				fmt.Println("Failed to connect to database:", err)
				return
			}
			// 2. 创建数据库
			dbName := "tc"
			err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName).Error
			if err != nil {
				fmt.Println("Failed to create database:", err)
				return
			}
			dsn = fmt.Sprintf("root:root123@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbName)

			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

			// 自动迁移同步
			db.AutoMigrate(&model.TcResult{}, &model.TcPrizeLevel{})

			// fmt.Println(fcResults)
			// 逐条插入
			for _, tcResult := range tcResults {
				// 开启事务
				err = db.Transaction(func(tx *gorm.DB) error {
					// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
					var maxID uint
					// COALESCE(MAX(id), 0) 用于在 MAX(id) 返回空值时将其替换为零。这样，即使表是新建的且没有任何记录，也不会出现空值的问题。
					if err := tx.Model(&model.TcResult{}).Select("COALESCE(MAX(id), 0)").Scan(&maxID).Error; err != nil {

						return err
					}
					// 手动分配新的ID
					tcResult.ID = maxID + 1
					// 插入数据
					if err := tx.Create(&tcResult).Error; err != nil {
						// 返回任何错误都会回滚事务
						return err
					}
					// 返回 nil 提交事务
					return nil
				})
			}

		}

	})
	c.Visit(url)
	// 等待所有访问完成
	c.Wait()
}
