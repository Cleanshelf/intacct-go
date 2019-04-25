package intacct

import (
	"encoding/xml"
)

type APPayment struct {
	XMLName xml.Name `xml:"appymt"`

	Financialentity      string            `xml:"FINANCIALENTITY"`
	Paymentmethod        string            `xml:"PAYMENTMETHOD"`
	Paymentrequestmethod string            `xml:"PAYMENTREQUESTMETHOD"`
	Exch_rate_date       string            `xml:"EXCH_RATE_DATE"`
	Exch_rate_type_id    string            `xml:"EXCH_RATE_TYPE_ID"`
	VendorID             string            `xml:"VENDORID"`
	DocNumber            string            `xml:"DOCNUMBER"`
	Description          string            `xml:"DESCRIPTION"`
	Paymentdate          string            `xml:"PAYMENTDATE"`
	Currency             string            `xml:"CURRENCY"`
	Basecurr             string            `xml:"BASECURR"`
	Amounttopay          string            `xml:"AMOUNTTOPAY"`
	Action               string            `xml:"ACTION"`
}

type APPayments struct {
	Client
}

func (apPayments APPayments) List(vendorID string, fromDate string, limit int) ([]APPayment, error) {
	itemList := make([]APPayment, 0)

	list := ReadByQuery{
		Object:   "APPYMT",
		Fields:   "*",
		Query:    "VENDORID='" + vendorID + "' AND WHENCREATED >= '" + fromDate + "'",
		Pagesize: 1000,
	}

	data, next, err := apPayments.Client.makeRequestByQuery(list)
	if err != nil {
		return itemList, err
	}

	itemList = data.APPayments

	if len(itemList) >= limit {
		return itemList[:limit], nil
	}

	for next != "" {
		list := ReadMore{
			ResultId: next,
		}

		var err error
		pageData, _, err := apPayments.Client.makeRequestByQuery(list)
		if err != nil {
			return itemList, err
		}

		page := pageData.APPayments
		for _, item := range page {
			itemList = append(itemList, item)
			if len(itemList) >= limit {
				return itemList[:limit], nil
			}
		}
	}

	return itemList, nil
}
