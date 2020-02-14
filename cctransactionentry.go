package intacct

import (
	"encoding/xml"
	"strings"
)

// RECORDKEY is foreign key of CCTransaction.RECORDNO
type CCTransactionEntry struct {
	XMLName xml.Name `xml:"cctransactionentry"`

	RECORDNO        string `xml:"RECORDNO"`
	RECORDKEY       string `xml:"RECORDKEY"`
	RECORDTYPE      string `xml:"RECORDTYPE"`
	ACCOUNTKEY      string `xml:"ACCOUNTKEY"`
	ACCOUNTNO       string `xml:"ACCOUNTNO"`
	ACCOUNTTITLE    string `xml:"ACCOUNTTITLE"`
	ACCOUNTLABELKEY string `xml:"ACCOUNTLABELKEY"`
	ACCOUNTLABEL    string `xml:"ACCOUNTLABEL"`
	AMOUNT          string `xml:"AMOUNT"`
	TRXAMOUNT       string `xml:"TRX_AMOUNT"`
	DEPT            string `xml:"DEPT"`
	DEPARTMENTID    string `xml:"DEPARTMENTID"`
	DEPARTMENTNAME  string `xml:"DEPARTMENTNAME"`
	LOCATION        string `xml:"LOCATION"`
	LOCATIONID      string `xml:"LOCATIONID"`
	LOCATIONNAME    string `xml:"LOCATIONNAME"`
	DESCRIPTION     string `xml:"DESCRIPTION"`
	EXCHRATEDATE    string `xml:"EXCH_RATE_DATE"`
	EXCHRATETYPEID  string `xml:"EXCH_RATE_TYPE_ID"`
	EXCHANGERATE    string `xml:"EXCHANGE_RATE"`
	LINEITEM        string `xml:"LINEITEM"`
	LINENO          string `xml:"LINE_NO"`
	CURRENCY        string `xml:"CURRENCY"`
	BASECURR        string `xml:"BASECURR"`
	STATUS          string `xml:"STATUS"`
	TOTALPAID       string `xml:"TOTALPAID"`
	TRXTOTALPAID    string `xml:"TRX_TOTALPAID"`
	DEPARTMENTKEY   string `xml:"DEPARTMENTKEY"`
	LOCATIONKEY     string `xml:"LOCATIONKEY"`
	WHENCREATED     string `xml:"WHENCREATED"`
	WHENMODIFIED    string `xml:"WHENMODIFIED"`
	CREATEDBY       string `xml:"CREATEDBY"`
	MODIFIEDBY      string `xml:"MODIFIEDBY"`
	PROJECTDIMKEY   string `xml:"PROJECTDIMKEY"`
	PROJECTID       string `xml:"PROJECTID"`
	PROJECTNAME     string `xml:"PROJECTNAME"`
	CUSTOMERDIMKEY  string `xml:"CUSTOMERDIMKEY"`
	CUSTOMERID      string `xml:"CUSTOMERID"`
	CUSTOMERNAME    string `xml:"CUSTOMERNAME"`
	VENDORDIMKEY    string `xml:"VENDORDIMKEY"`
	VENDORID        string `xml:"VENDORID"`
	VENDORNAME      string `xml:"VENDORNAME"`
	EMPLOYEEDIMKEY  string `xml:"EMPLOYEEDIMKEY"`
	EMPLOYEEID      string `xml:"EMPLOYEEID"`
	EMPLOYEENAME    string `xml:"EMPLOYEENAME"`
	ITEMDIMKEY      string `xml:"ITEMDIMKEY"`
	ITEMID          string `xml:"ITEMID"`
	ITEMNAME        string `xml:"ITEMNAME"`
	CLASSDIMKEY     string `xml:"CLASSDIMKEY"`
	CLASSID         string `xml:"CLASSID"`
	CLASSNAME       string `xml:"CLASSNAME"`
}

type CCTransactionEntries struct {
	Client
}

func (ccTransactionEntries CCTransactionEntries) ListByTransactionIds(transactionIDs []string) ([]CCTransactionEntry, error) {
	if transactionIDs == nil || len(transactionIDs) == 0 {
		return make([]CCTransactionEntry, 0), nil
	}

	query := ReadByQuery{
		Object:   "CCTRANSACTIONENTRY",
		Fields:   "*",
		Query:    "RECORDKEY IN ('" + strings.Join(transactionIDs, "','") + "')",
		Pagesize: 200,
	}

	data, next, err := ccTransactionEntries.Client.makeRequestByQuery(query)
	if err != nil {
		return make([]CCTransactionEntry, 0), err
	}

	ccEntries := data.CCTransactionEntries
	total := data.TotalCount
	for total > len(ccEntries) && next != "" {
		query := ReadMore{
			ResultId: next,
		}

		var err error
		var pageData *Data
		pageData, next, err = ccTransactionEntries.Client.makeRequestByQuery(query)
		if err != nil {
			return ccEntries, err
		}

		entries := pageData.CCTransactionEntries
		for _, entry := range entries {
			ccEntries = append(ccEntries, entry)
		}
	}

	return ccEntries, nil
}
