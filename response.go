package intacct

import (
	"encoding/xml"
)

// Status consts
const (
	Failure = "failure"
	Success = "success"
)

// TODO Set result type method?
type Response struct {
	XMLName   xml.Name `xml:"response"`
	Control   Control
	Operation ResultOperation
	Errors    []Error `xml:"errormessage"`
}

type ResultOperation struct {
	XMLName        xml.Name `xml:"operation"`
	Authentication Authentication
	Content        Content `xml:"content"`
	Result         Result  `xml:"result"`
}

// TODO Or use delayed parsing?
type Result struct {
	Status    string `xml:"status"`
	Function  string `xml:"function"`
	ControlID string `xml:"controlid"`
	// TODO listtype?
	Data Data `xml:"data"`
}

type Data struct {
	NumRemaining int    `xml:"numRemaining,attr"`
	ResultId     string `xml:"resultId,attr"`
	TotalCount   int    `xml:"totalcount,attr"`

	Invoices   []Invoice   `xml:"invoice"`
	Vendors    []Vendor    `xml:"vendor"`
	Bills      []Bill      `xml:"apbill"`
	APPayments []APPayment `xml:"appymt"`
	Customers  []Customer  `xml:"customer"`
}
