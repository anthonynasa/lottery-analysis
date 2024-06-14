package model

import "gorm.io/gorm"

type FcResponse struct {
	State    int      `json:"state"`
	Message  string   `json:"message"`
	Total    int      `json:"total"`
	PageNum  int      `json:"pageNum"`
	PageNo   int      `json:"pageNo"`
	PageSize int      `json:"pageSize"`
	Tflag    int      `json:"Tflag"`
	Result   []Result `json:"result"`
}
type Result struct {
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
}

type FcResult struct {
	gorm.Model
	Name        string         `json:"name"`
	Code        string         `json:"code" gorm:"size:64;uniqueIndex"`
	DetailsLink string         `json:"detailsLink"`
	VideoLink   string         `json:"videoLink"`
	Date        string         `json:"date"`
	Week        string         `json:"week"`
	Red         string         `json:"red"`
	Blue        string         `json:"blue"`
	Blue2       string         `json:"blue2"`
	Sales       string         `json:"sales"`
	Poolmoney   string         `json:"poolmoney"`
	Content     string         `json:"content"`
	Addmoney    string         `json:"addmoney"`
	Addmoney2   string         `json:"addmoney2"`
	Msg         string         `json:"msg"`
	Z2Add       string         `json:"z2add"`
	M2Add       string         `json:"m2add"`
	Prizegrades []FcPrizegrade `json:"prizegrades"`
}
type FcPrizegrade struct {
	gorm.Model
	Type       int    `json:"type"`
	Typenum    string `json:"typenum"`
	Typemoney  string `json:"typemoney"`
	FcResultId int    `json:"fc_result_id" `
}

