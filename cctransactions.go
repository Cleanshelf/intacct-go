package intacct

import (
	"encoding/xml"
)

type CCTransaction struct {
	XMLName xml.Name `xml:"cctransaction"`

	RECORDNO        string `xml:"RECORDNO"`
	FINANCIALENTITY string `xml:"FINANCIALENTITY"`
	WHENCREATED     string `xml:"WHENCREATED"`
	RECORDID        string `xml:"RECORDID"`
	RECORDTYPE      string `xml:"RECORDTYPE"`
	DESCRIPTION2    string `xml:"DESCRIPTION2"`
	DESCRIPTION     string `xml:"DESCRIPTION"`
	PRBATCHKEY      string `xml:"PRBATCHKEY"`
	BASECURR        string `xml:"BASECURR"`
	CURRENCY        string `xml:"CURRENCY"`
	EXCHRATEDATE    string `xml:"EXCH_RATE_DATE"`
	EXCHRATETYPEID  string `xml:"EXCH_RATE_TYPE_ID"`
	EXCHANGERATE    string `xml:"EXCHANGE_RATE"`
	TOTALENTERED    string `xml:"TOTALENTERED"`
	TRXTOTALENTERED string `xml:"TRX_TOTALENTERED"`
	TOTALPAID       string `xml:"TOTALPAID"`
	TRXTOTALPAID    string `xml:"TRX_TOTALPAID"`
	WHENPAID        string `xml:"WHENPAID"`
	REVERSALKEY     string `xml:"REVERSALKEY"`
	REVERSALDATE    string `xml:"REVERSALDATE"`
	REVERSEDKEY     string `xml:"REVERSEDKEY"`
	REVERSEDDATE    string `xml:"REVERSEDDATE"`
	STATE           string `xml:"STATE"`
	RAWSTATE        string `xml:"RAWSTATE"`
	CLEARED         string `xml:"CLEARED"`
	PAYMENTKEY      string `xml:"PAYMENTKEY"`
	AUWHENCREATED   string `xml:"AUWHENCREATED"`
	WHENMODIFIED    string `xml:"WHENMODIFIED"`
	CREATEDBY       string `xml:"CREATEDBY"`
	MODIFIEDBY      string `xml:"MODIFIEDBY"`
	MEGAENTITYKEY   string `xml:"MEGAENTITYKEY"`
	MEGAENTITYID    string `xml:"MEGAENTITYID"`
	MEGAENTITYNAME  string `xml:"MEGAENTITYNAME"`
}

type CCTransactions struct {
	Client
}

func (ccTransactions CCTransactions) List(fromDate string, limit int) ([]CCTransaction, error) {
	itemList := make([]CCTransaction, 0)

	list := ReadByQuery{
		Object:   "CCTRANSACTION",
		Fields:   "*",
		Query:    "WHENCREATED >= '" + fromDate + "'",
		Pagesize: 200,
	}

	data, next, err := ccTransactions.Client.makeRequestByQuery(list)
	if err != nil {
		return itemList, err
	}

	itemList = data.CCTransactions
	total := data.TotalCount
	for total > len(itemList) && next != "" {
		list := ReadMore{
			ResultId: next,
		}

		var err error
		var pageData *Data
		pageData, next, err = ccTransactions.Client.makeRequestByQuery(list)
		if err != nil {
			return itemList, err
		}

		page := pageData.CCTransactions
		for _, item := range page {
			itemList = append(itemList, item)
			if len(itemList) >= limit {
				return itemList[:limit], nil
			}
		}
	}

	return itemList, nil
}
