package model

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
	R7          string `json:"r7"`
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
