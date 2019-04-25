package intacct

import (
	"encoding/xml"
)

type CCTransaction struct {
	XMLName xml.Name `xml:"cctransaction"`

	RecordNo        string `xml:"RECORDNO"`
	Financialentity string `xml:"FINANCIALENTITY"`
	WhenCreated     string `xml:"WHENCREATED"`
	Description     string `xml:"DESCRIPTION"`
	Description2    string `xml:"DESCRIPTION2"`
	WhenPaid        string `xml:"WHENPAID"`
	Currency        string `xml:"CURRENCY"`
	Total           string `xml:"TOTALENTERED"`
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
		Pagesize: 1000,
	}

	data, next, err := ccTransactions.Client.makeRequestByQuery(list)
	if err != nil {
		return itemList, err
	}

	itemList = data.CCTransactions

	if len(itemList) >= limit {
		return itemList[:limit], nil
	}

	for next != "" {
		list := ReadMore{
			ResultId: next,
		}
		var err error
		pageData, _, err := ccTransactions.Client.makeRequestByQuery(list)
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

