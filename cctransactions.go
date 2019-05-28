package intacct

import (
	"encoding/xml"
)

type CCTransaction struct {
	XMLName xml.Name `xml:"cctransaction"`

	RecordNo        string              `xml:"RECORDNO"`
	Financialentity string              `xml:"FINANCIALENTITY"`
	Description     string              `xml:"DESCRIPTION"`
	Description2    string              `xml:"DESCRIPTION2"`
	RecordID        string              `xml:"RECORDID"`
	WhenCreated     string              `xml:"WHENCREATED"`
	Currency        string              `xml:"CURRENCY"`
	Total           string              `xml:"TOTALENTERED"`
	Item            []CCTransactionItem `xml:"ccpayitems"`
}

type CCTransactionItem struct {
	XMLName xml.Name `xml:"ccpayitem"`

	Description   string `xml:"DESCRIPTION"`
	PaymentAmount string `xml:"PAYMENTAMOUNT"`
	ItemID        string `xml:"ITEMID"`
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
