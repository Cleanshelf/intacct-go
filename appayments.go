package intacct

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Appymtdetails        []APPaymentDetail `xml:"APPYMTDETAILS"`
}

type APPaymentDetail struct {
	XMLName xml.Name `xml:"appymtdetail"`

	Recordkey               string `xml:"RECORDKEY"`
	Entrykey                string `xml:"ENTRYKEY"`
	Entrycurrency           string `xml:"ENTRYCURRENCY"`
	Posadjkey               string `xml:"POSADJKEY"`
	Posadjentrykey          string `xml:"POSADJENTRYKEY"`
	Trx_paymentamount       string `xml:"TRX_PAYMENTAMOUNT"`
	Inlinekey               string `xml:"INLINEKEY"`
	Inlineentrykey          string `xml:"INLINEENTRYKEY"`
	Trx_inlineamount        string `xml:"TRX_INLINEAMOUNT"`
	Discountdate            string `xml:"DISCOUNTDATE"`
	Adjustmentkey           string `xml:"ADJUSTMENTKEY"`
	Adjustmententrykey      string `xml:"ADJUSTMENTENTRYKEY"`
	Trx_adjustmentamount    string `xml:"TRX_ADJUSTMENTAMOUNT"`
	Advancekey              string `xml:"ADVANCEKEY"`
	Advanceentrykey         string `xml:"ADVANCEENTRYKEY"`
	Trx_postedadvanceamount string `xml:"TRX_POSTEDADVANCEAMOUNT"`
}
type APPayments struct {
	Client
}

func (apPayments APPayments) List() ([]APPayment, error) {
	list := ReadByQuery{
		Object:   "APPYMT",
		Fields:   "*",
		Query:    "",
		Pagesize: 1000,
	}

	get := Function{
		ControlID: "testControlID",
		Method:    list,
	}
	// Create a new request using the Client
	req, err := apPayments.Client.NewRequest(get)
	if err != nil {
		return nil, err
	}

	resp, err := apPayments.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var body Response
	if err = xml.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	// Check the response for errors
	if err = apPayments.Client.CheckResponseErrors(body); err != nil {
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

	return body.Operation.Result.Data.APPayments, nil
}
