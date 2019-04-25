package intacct

import (
	"encoding/xml"
)

type APRecurBill struct {
	XMLName     xml.Name    `xml:"aprecurbill"`
	Description string      `xml:"DESCRIPTION"`
	RecordID    string      `xml:"RECORDID"`
	DocNumber   string      `xml:"DOCNUMBER"`
	StartDate   IntacctDate `xml:"STARTDATE"`
	Mode        string      `xml:"MODENEW"`
	Internal    string      `xml:"INTERVAL"`

	RecurBillItems APRecurBillItems `xml:"RECURBILLITEMS"`
}

type IntacctDate struct {
	Year  string `xml:"YEAR"`
	Month string `xml:"MONTH"`
	Day   string `xml:"DAY"`
}

type APRecurBillItems struct {
	Amount string `xml:"AMOUNT"`
}

type APRecurBills struct {
	Client
}

func (apRecurBills APRecurBills) List(vendorID string, fromDate string, limit int) ([]APRecurBill, error) {
	itemList := make([]APRecurBill, 0)

	list := ReadByQuery{
		Object:   "APBILL",
		Fields:   "*", //TODO
		Query:    "VENDORID='" + vendorID + "' AND WHENCREATED >= '" + fromDate + "'",
		Pagesize: 1000,
	}

	data, next, err := apRecurBills.Client.makeRequestByQuery(list)
	if err != nil {
		return itemList, err
	}

	itemList = data.APRecurBills

	if len(itemList) >= limit {
		return itemList[:limit], nil
	}

	for next != "" {
		list := ReadMore{
			ResultId: next,
		}

		var err error
		pageData, _, err := apRecurBills.Client.makeRequestByQuery(list)
		if err != nil {
			return itemList, err
		}

		page := pageData.APRecurBills
		for _, item := range page {
			itemList = append(itemList, item)
			if len(itemList) >= limit {
				return itemList[:limit], nil
			}
		}
	}

	return itemList, nil
}
