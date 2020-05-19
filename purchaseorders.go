package intacct

import (
	"encoding/xml"
)

type PODocument struct {
	XMLName xml.Name `xml:"podocument"`

	RecordNo     string `xml:"RECORDNO"`
	DocNo        string `xml:"DOCNO"`
	DocID        string `xml:"DOCID"`
	PONumber     string `xml:"PONUMBER"`
	WhenCreated  string `xml:"WHENCREATED"`
	CustVendID   string `xml:"CUSTVENDID"`
	CustVendName string `xml:"CUSTVENDNAME"`
	DocParClass  string `xml:"DOCPARCLASS"`
	Note         string `xml:"NOTE"`
	Message      string `xml:"MESSAGE"`
	Total        string `xml:"TOTAL"`
	Currency     string `xml:"CURRENCY"`
}

type PODocuments struct {
	Client
}

func (poDocuments PODocuments) List(vendorID string, fromDate string, limit int) ([]PODocument, error) {
	itemList := make([]PODocument, 0)

	list := ReadByQuery{
		Object:   "PODOCUMENT",
		Fields:   "*",
		Query:    "WHENCREATED > '" + fromDate + "'",
		Pagesize: 1000,
	}

	if vendorID != "" {
		list.Query = "CUSTVENDID ='" + vendorID + "' AND " + list.Query
	}

	data, next, err := poDocuments.Client.makeRequestByQuery(list)
	if err != nil {
		return itemList, err
	}

	itemList = data.PODocuments

	for next != "" {
		list := ReadMore{
			ResultId: next,
		}
		var err error
		var pageData *Data
		pageData, next, err = poDocuments.Client.makeRequestByQuery(list)
		if err != nil {
			return itemList, err
		}

		page := pageData.PODocuments
		for _, item := range page {
			itemList = append(itemList, item)
			if len(itemList) >= limit {
				return itemList[:limit], nil
			}
		}
	}

	return itemList, nil
}
