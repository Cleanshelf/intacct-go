package intacct

import (
	"encoding/xml"
)

type APPayment struct {
	XMLName xml.Name `xml:"appymt"`

	RECORDNO                 string `xml:"RECORDNO"`
	PRBATCHKEY               string `xml:"PRBATCHKEY"`
	RECORDTYPE               string `xml:"RECORDTYPE"`
	RECORDID                 string `xml:"RECORDID"`
	FINANCIALENTITY          string `xml:"FINANCIALENTITY"`
	FINANCIALENTITYTYPE      string `xml:"FINANCIALENTITYTYPE"`
	FINANCIALACCOUNTNAME     string `xml:"FINANCIALACCOUNTNAME"`
	FINANCIALACCOUNTCURRENCY string `xml:"FINANCIALACCOUNTCURRENCY"`
	STATE                    string `xml:"STATE"`
	RAWSTATE                 string `xml:"RAWSTATE"`
	PAYMENTDATE              string `xml:"PAYMENTDATE"`
	PAYMENTMETHOD            string `xml:"PAYMENTMETHOD"`
	PAYMENTMETHODKEY         string `xml:"PAYMENTMETHODKEY"`
	PAYMENTREQUESTMETHOD     string `xml:"PAYMENTREQUESTMETHOD"`
	ENTITY                   string `xml:"ENTITY"`
	VENDORID                 string `xml:"VENDORID"`
	VENDORNAME               string `xml:"VENDORNAME"`
	DOCNUMBER                string `xml:"DOCNUMBER"`
	DESCRIPTION              string `xml:"DESCRIPTION"`
	DESCRIPTION2             string `xml:"DESCRIPTION2"`
	WHENCREATED              string `xml:"WHENCREATED"`
	WHENPAID                 string `xml:"WHENPAID"`
	BASECURR                 string `xml:"BASECURR"`
	CURRENCY                 string `xml:"CURRENCY"`
	EXCHRATEDATE             string `xml:"EXCH_RATE_DATE"`
	EXCHRATETYPEID           string `xml:"EXCH_RATE_TYPE_ID"`
	EXCHANGERATE             string `xml:"EXCHANGE_RATE"`
	TOTALENTERED             string `xml:"TOTALENTERED"`
	TOTALSELECTED            string `xml:"TOTALSELECTED"`
	TOTALPAID                string `xml:"TOTALPAID"`
	TOTALDUE                 string `xml:"TOTALDUE"`
	TRXTOTALENTERED          string `xml:"TRX_TOTALENTERED"`
	TRXTOTALSELECTED         string `xml:"TRX_TOTALSELECTED"`
	TRXTOTALPAID             string `xml:"TRX_TOTALPAID"`
	TRXTOTALDUE              string `xml:"TRX_TOTALDUE"`
	BILLTOPAYTOCONTACTNAME   string `xml:"BILLTOPAYTOCONTACTNAME"`
	PRBATCH                  string `xml:"PRBATCH"`
	AMOUNTTOPAY              string `xml:"AMOUNTTOPAY"`
	ACTION                   string `xml:"ACTION"`
	AUWHENCREATED            string `xml:"AUWHENCREATED"`
	WHENMODIFIED             string `xml:"WHENMODIFIED"`
	CREATEDBY                string `xml:"CREATEDBY"`
	MODIFIEDBY               string `xml:"MODIFIEDBY"`
	USERKEY                  string `xml:"USERKEY"`
	CLEARED                  string `xml:"CLEARED"`
	CLRDATE                  string `xml:"CLRDATE"`
	STATUS                   string `xml:"STATUS"`
	SYSTEMGENERATED          string `xml:"SYSTEMGENERATED"`
	PAYMENTPRIORITY          string `xml:"PAYMENTPRIORITY"`
	BILLTOPAYTOKEY           string `xml:"BILLTOPAYTOKEY"`
	ONHOLD                   string `xml:"ONHOLD"`
	PARENTPAYMENTKEY         string `xml:"PARENTPAYMENTKEY"`
	WHENPOSTED               string `xml:"WHENPOSTED"`
	MEGAENTITYKEY            string `xml:"MEGAENTITYKEY"`
	MEGAENTITYID             string `xml:"MEGAENTITYID"`
	MEGAENTITYNAME           string `xml:"MEGAENTITYNAME"`
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

	for next != "" {
		list := ReadMore{
			ResultId: next,
		}

		var err error
		var pageData *Data
		pageData, next, err = apPayments.Client.makeRequestByQuery(list)
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
