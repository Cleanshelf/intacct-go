package intacct

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bill struct {
	XMLName xml.Name `xml:"apbill"`

	RecordNO                  string     `xml:"RECORDNO"`
	RecordType                string     `xml:"RECORDTYPE"`
	RecordID                  string     `xml:"RECORDID"`
	FinancialEntity           string     `xml:"FINANCIALENTITY"`
	State                     string     `xml:"STATE"`
	Rawstate                  string     `xml:"RAWSTATE"`
	VendorID                  string     `xml:"VENDORID"`
	VendorName                string     `xml:"VENDORNAME"`
	Form1099type              string     `xml:"FORM1099TYPE"`
	Form1099box               string     `xml:"FORM1099BOX"`
	Vendtype1099type          string     `xml:"VENDTYPE1099TYPE"`
	Trx_entitydue             string     `xml:"TRX_ENTITYDUE"`
	DocNumber                 string     `xml:"DOCNUMBER"`
	Description               string     `xml:"DESCRIPTION"`
	Description2              string     `xml:"DESCRIPTION2"`
	Termname                  string     `xml:"TERMNAME"`
	Termkey                   string     `xml:"TERMKEY"`
	Termvalue                 string     `xml:"TERMVALUE"`
	WhenCreated               string     `xml:"WHENCREATED"`
	WhenPosted                string     `xml:"WHENPOSTED"`
	WhenDiscount              string     `xml:"WHENDISCOUNT"`
	WhenDue                   string     `xml:"WHENDUE"`
	WhenPaid                  string     `xml:"WHENPAID"`
	Recpaymentdate            string     `xml:"RECPAYMENTDATE"`
	Paymentpriority           string     `xml:"PAYMENTPRIORITY"`
	Onhold                    bool       `xml:"ONHOLD"`
	Basecurr                  string     `xml:"BASECURR"`
	Currency                  string     `xml:"CURRENCY"`
	Exch_rate_date            string     `xml:"EXCH_RATE_DATE"`
	Exch_rate_type_id         string     `xml:"EXCH_RATE_TYPE_ID"`
	Exchange_rate             string     `xml:"EXCHANGE_RATE"`
	TotalEntered              float64    `xml:"TOTALENTERED"`
	TotalSelected             float64    `xml:"TOTALSELECTED"`
	TotalPaid                 float64    `xml:"TOTALPAID"`
	TotalDue                  float64    `xml:"TOTALDUE"`
	TrxTotalEntered           float64    `xml:"TRX_TOTALENTERED"`
	TrxTotalSelected          float64    `xml:"TRX_TOTALSELECTED"`
	TrxTotalPaid              float64    `xml:"TRX_TOTALPAID"`
	TrxTotalDue               float64    `xml:"TRX_TOTALDUE"`
	Billtopaytocontactname    string     `xml:"BILLTOPAYTOCONTACTNAME"`
	Shiptoreturntocontactname string     `xml:"SHIPTORETURNTOCONTACTNAME"`
	Billtopaytokey            string     `xml:"BILLTOPAYTOKEY"`
	Shiptoreturntokey         string     `xml:"SHIPTORETURNTOKEY"`
	Prbatch                   string     `xml:"PRBATCH"`
	Prbatchkey                string     `xml:"PRBATCHKEY"`
	Modulekey                 string     `xml:"MODULEKEY"`
	Schopkey                  string     `xml:"SCHOPKEY"`
	Systemgenerated           string     `xml:"SYSTEMGENERATED"`
	Auwhencreated             string     `xml:"AUWHENCREATED"`
	Whenmodified              string     `xml:"WHENMODIFIED"`
	Createdby                 string     `xml:"CREATEDBY"`
	Modifiedby                string     `xml:"MODIFIEDBY"`
	Due_in_days               string     `xml:"DUE_IN_DAYS"`
	SupdocID                  string     `xml:"SUPDOCID"`
	Megaentitykey             string     `xml:"MEGAENTITYKEY"`
	MegaentityID              string     `xml:"MEGAENTITYID"`
	MegaentityName            string     `xml:"MEGAENTITYNAME"`
	Vendor_notes              string     `xml:"VENDOR_NOTES"`
	Voucher_number            string     `xml:"VOUCHER_NUMBER"`
	Exchratedate              string     `xml:"EXCHRATEDATE"`
	Exchratetype              string     `xml:"EXCHRATETYPE"`
	Exchrate                  string     `xml:"EXCHRATE"`
	Paytocontactname          string     `xml:"PAYTOCONTACTNAME"`
	Returntocontactname       string     `xml:"RETURNTOCONTACTNAME"`
	BillItems                 []BillItem `xml:"APBILLITEMS"`
}

type BillItem struct {
	XMLName xml.Name `xml:"apbillitem"`

	Line_no           int     `xml:LINE_NO"`
	Accountno         string  `xml:ACCOUNTNO"`
	Accountlabel      string  `xml:ACCOUNTLABEL"`
	Offsetglaccountno string  `xml:OFFSETGLACCOUNTNO"`
	Trx_amount        float64 `xml:TRX_AMOUNT"`
	EntryDescription  string  `xml:ENTRYDESCRIPTION"`
	Form1099          bool    `xml:FORM1099"`
	Form1099type      string  `xml:FORM1099TYPE"`
	Form1099box       string  `xml:FORM1099BOX"`
	Billable          bool    `xml:BILLABLE"`
	Allocation        string  `xml:ALLOCATION"`
	LocationID        string  `xml:LOCATIONID"`
	DepartmentID      string  `xml:DEPARTMENTID"`
	ProjectID         string  `xml:PROJECTID"`
	CustomerID        string  `xml:CUSTOMERID"`
	VendorID          string  `xml:VENDORID"`
	EmployeeID        string  `xml:EMPLOYEEID"`
	ItemID            string  `xml:ITEMID"`
	ClassID           string  `xml:CLASSID"`
	ContractID        string  `xml:CONTRACTID"`
	WarehouseID       string  `xml:WAREHOUSEID"`
	// Gldim int `xml:GLDIM"`
	Custom string `xml:Custom"`
}
type Bills struct {
	Client
}

func (bills Bills) List(vendorID string, limit int) ([]Bill, error) {
	list := ReadByQuery{
		Object:   "APBILL",
		Fields:   "*", //TODO
		Query:    "VENDORID = '" + vendorID + "'",
		Pagesize: limit,
	}

	get := Function{
		ControlID: "testControlID",
		Method:    list,
	}
	// Create a new request using the Client
	req, err := bills.Client.NewRequest(get)
	if err != nil {
		return nil, err
	}

	resp, err := bills.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body Response
	if err = xml.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	// Check the response for errors
	if err = bills.Client.CheckResponseErrors(body); err != nil {
		return nil, err
	}

	// TODO pull out status code and body status checks into client
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		bodyString := string(body)
		return nil, fmt.Errorf("non-200 status code: %d, error: %s", resp.StatusCode, bodyString)
	}

	return body.Operation.Result.Data.Bills, nil
}
