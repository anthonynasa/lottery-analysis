package model

import "gorm.io/gorm"

type FcResponse struct {
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
	} `json:"result"`
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

type TcResponse struct {
	DataFrom     string `json:"dataFrom"`
	EmptyFlag    bool   `json:"emptyFlag"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
	Value        struct {
		LastPoolDraw struct {
			LotteryDrawNum       string `json:"lotteryDrawNum"`
			LotteryDrawResult    string `json:"lotteryDrawResult"`
			LotteryDrawTime      string `json:"lotteryDrawTime"`
			LotteryGameName      string `json:"lotteryGameName"`
			LotteryGameNum       string `json:"lotteryGameNum"`
			PoolBalanceAfterdraw string `json:"poolBalanceAfterdraw"`
			PrizeLevelList       []struct {
				AwardType         int    `json:"awardType"`
				Group             string `json:"group"`
				LotteryCondition  string `json:"lotteryCondition"`
				PrizeLevel        string `json:"prizeLevel"`
				Sort              int    `json:"sort"`
				StakeAmount       string `json:"stakeAmount"`
				StakeAmountFormat string `json:"stakeAmountFormat"`
				StakeCount        string `json:"stakeCount"`
				TotalPrizeamount  string `json:"totalPrizeamount"`
			} `json:"prizeLevelList"`
		} `json:"lastPoolDraw"`
		List []struct {
			DrawFlowFund            string        `json:"drawFlowFund"`
			DrawFlowFundRj          string        `json:"drawFlowFundRj"`
			DrawPdfUrl              string        `json:"drawPdfUrl"`
			EstimateDrawTime        string        `json:"estimateDrawTime"`
			IsGetKjpdf              int           `json:"isGetKjpdf"`
			IsGetXlpdf              int           `json:"isGetXlpdf"`
			LotteryDrawNum          string        `json:"lotteryDrawNum"`
			LotteryDrawResult       string        `json:"lotteryDrawResult"`
			LotteryDrawStatus       int           `json:"lotteryDrawStatus"`
			LotteryDrawStatusNo     string        `json:"lotteryDrawStatusNo"`
			LotteryDrawTime         string        `json:"lotteryDrawTime"`
			LotteryEquipmentCount   int           `json:"lotteryEquipmentCount"`
			LotteryGameName         string        `json:"lotteryGameName"`
			LotteryGameNum          string        `json:"lotteryGameNum"`
			LotteryGamePronum       int           `json:"lotteryGamePronum"`
			LotteryNotice           int           `json:"lotteryNotice"`
			LotteryNoticeShowFlag   int           `json:"lotteryNoticeShowFlag"`
			LotteryPaidBeginTime    string        `json:"lotteryPaidBeginTime"`
			LotteryPaidEndTime      string        `json:"lotteryPaidEndTime"`
			LotteryPromotionFlag    int           `json:"lotteryPromotionFlag"`
			LotteryPromotionFlagRj  int           `json:"lotteryPromotionFlagRj"`
			LotterySaleBeginTime    string        `json:"lotterySaleBeginTime"`
			LotterySaleEndTimeUnix  int           `json:"lotterySaleEndTimeUnix"`
			LotterySaleEndtime      string        `json:"lotterySaleEndtime"`
			LotterySuspendedFlag    int           `json:"lotterySuspendedFlag"`
			LotteryUnsortDrawresult string        `json:"lotteryUnsortDrawresult"`
			MatchList               []interface{} `json:"matchList"`
			PdfType                 int           `json:"pdfType"`
			PoolBalanceAfterdraw    string        `json:"poolBalanceAfterdraw"`
			PoolBalanceAfterdrawRj  string        `json:"poolBalanceAfterdrawRj"`
			PrizeLevelList          []struct {
				AwardType         int    `json:"awardType"`
				Group             string `json:"group"`
				LotteryCondition  string `json:"lotteryCondition"`
				PrizeLevel        string `json:"prizeLevel"`
				Sort              int    `json:"sort"`
				StakeAmount       string `json:"stakeAmount"`
				StakeAmountFormat string `json:"stakeAmountFormat"`
				StakeCount        string `json:"stakeCount"`
				TotalPrizeamount  string `json:"totalPrizeamount"`
			} `json:"prizeLevelList"`
			PrizeLevelListRj  []interface{} `json:"prizeLevelListRj"`
			RuleType          int           `json:"ruleType"`
			SurplusAmount     string        `json:"surplusAmount"`
			SurplusAmountRj   string        `json:"surplusAmountRj"`
			TermList          []interface{} `json:"termList"`
			TermResultList    []interface{} `json:"termResultList"`
			TotalSaleAmount   string        `json:"totalSaleAmount"`
			TotalSaleAmountRj string        `json:"totalSaleAmountRj"`
			Verify            int           `json:"verify"`
			VtoolsConfig      struct {
			} `json:"vtoolsConfig"`
		} `json:"list"`
		PageNo   int `json:"pageNo"`
		PageSize int `json:"pageSize"`
		Pages    int `json:"pages"`
		Total    int `json:"total"`
	} `json:"value"`
}

type TcResult struct {
	gorm.Model
	LotteryGameName       string       `json:"lotteryGameName"`
	LotteryGameNum        string       `json:"lotteryGameNum"`
	LotteryDrawNum        string       `json:"lotteryDrawNum" gorm:"size:64;uniqueIndex"`
	DrawPdfUrl            string       `json:"drawPdfUrl"`
	LotteryDrawTime       string       `json:"lotteryDrawTime"`
	LotteryEquipmentCount int          `json:"lotteryEquipmentCount"`
	LotteryDrawResult     string       `json:"lotteryDrawResult"`
	TotalSaleAmount       string       `json:"totalSaleAmount"`
	PoolBalanceAfterdraw  string         `json:"poolBalanceAfterdraw"`
	PrizeLevelList        []TcPrizeLevel `json:"prizeLevelList"`
}

type TcPrizeLevel struct {
	gorm.Model
	AwardType         int    `json:"awardType"`
	Group             string `json:"group"`
	LotteryCondition  string `json:"lotteryCondition"`
	PrizeLevel        string `json:"prizeLevel"`
	Sort              int    `json:"sort"`
	StakeAmount       string `json:"stakeAmount"`
	StakeAmountFormat string `json:"stakeAmountFormat"`
	StakeCount        string `json:"stakeCount"`
	TotalPrizeamount  string `json:"totalPrizeamount"`
	TcResultId  string `json:"tc_result_id"`
}
