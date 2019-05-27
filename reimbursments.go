package intacct

import (
	"encoding/xml"
)

type EPPayment struct {
	XMLName xml.Name `xml:"eppayment"`

	RecordNo    string `xml:"RECORDNO"`
	WhenCreated string `xml:"WHENCREATED"`
	Description string `xml:"DESCRIPTION"`
	WhenPaid    string `xml:"WHENPAID"`
	Currency    string `xml:"CURRENCY"`
	Total       string `xml:"TOTALENTERED"`
}

type EPPayments struct {
	Client
}

func (epPayments EPPayments) List(fromDate string, limit int) ([]EPPayment, error) {
	itemList := make([]EPPayment, 0)

	list := ReadByQuery{
		Object:   "EPPAYMENT",
		Fields:   "*",
		Query:    "WHENCREATED > '" + fromDate + "'",
		Pagesize: 1000,
	}

	data, next, err := epPayments.Client.makeRequestByQuery(list)
	if err != nil {
		return itemList, err
	}

	itemList = data.EPPayments

	for next != "" {
		list := ReadMore{
			ResultId: next,
		}
		var err error
		var pageData *Data
		pageData, next, err = epPayments.Client.makeRequestByQuery(list)
		if err != nil {
			return itemList, err
		}

		page := pageData.EPPayments
		for _, item := range page {
			itemList = append(itemList, item)
			if len(itemList) >= limit {
				return itemList[:limit], nil
			}
		}
	}

	return itemList, nil
}
