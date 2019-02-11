package intacct

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Vendor struct {
	XMLName xml.Name `xml:"vendor"`

	VendorID              string `xml:"VENDORID"`
	Name                  string `xml:"NAME"`
	Printas               string `xml:"DISPLAYCONTACT>PRINTAS"`
	Companyname           string `xml:"DISPLAYCONTACT>COMPANYNAME"`
	Taxable               string `xml:"DISPLAYCONTACT>TAXABLE"`
	Taxgroup              string `xml:"DISPLAYCONTACT>TAXGROUP"`
	Prefix                string `xml:"DISPLAYCONTACT>PREFIX"`
	Firstname             string `xml:"DISPLAYCONTACT>FIRSTNAME"`
	Lastname              string `xml:"DISPLAYCONTACT>LASTNAME"`
	Initial               string `xml:"DISPLAYCONTACT>INITIAL"`
	Phone1                string `xml:"DISPLAYCONTACT>PHONE1"`
	Phone2                string `xml:"DISPLAYCONTACT>PHONE2"`
	Cellphone             string `xml:"DISPLAYCONTACT>CELLPHONE"`
	Pager                 string `xml:"DISPLAYCONTACT>PAGER"`
	Fax                   string `xml:"DISPLAYCONTACT>FAX"`
	Email1                string `xml:"DISPLAYCONTACT>EMAIL1"`
	Email2                string `xml:"DISPLAYCONTACT>EMAIL2"`
	Url1                  string `xml:"DISPLAYCONTACT>URL1"`
	Url2                  string `xml:"DISPLAYCONTACT>URL2"`
	Address1              string `xml:"DISPLAYCONTACT>MAILADDRESS>ADDRESS1"`
	Address2              string `xml:"DISPLAYCONTACT>MAILADDRESS>ADDRESS2"`
	City                  string `xml:"DISPLAYCONTACT>MAILADDRESS>CITY"`
	State                 string `xml:"DISPLAYCONTACT>MAILADDRESS>STATE"`
	Zip                   string `xml:"DISPLAYCONTACT>MAILADDRESS>ZIP"`
	Country               string `xml:"DISPLAYCONTACT>MAILADDRESS>COUNTRY"`
	Onetime               bool   `xml:"ONETIME"`
	Status                string `xml:"STATUS"`
	Hidedisplaycontact    string `xml:"HIDEDISPLAYCONTACT"`
	Vendtype              string `xml:"VENDTYPE"`
	Parentid              string `xml:"PARENTID"`
	Glgroup               string `xml:"GLGROUP"`
	Taxid                 string `xml:"TAXID"`
	Name1099              string `xml:"NAME1099"`
	Form1099type          string `xml:"FORM1099TYPE"`
	Form1099box           string `xml:"FORM1099BOX"`
	Supdocid              string `xml:"SUPDOCID"`
	Apaccount             string `xml:"APACCOUNT"`
	Creditlimit           string `xml:"CREDITLIMIT"`
	Onhold                string `xml:"ONHOLD"`
	Donotcutcheck         string `xml:"DONOTCUTCHECK"`
	Comments              string `xml:"COMMENTS"`
	Currency              string `xml:"CURRENCY"`
	Contactinfo           string `xml:"CONTACTINFO>CONTACTNAME"`
	Payto                 string `xml:"PAYTO>CONTACTNAME"`
	ReturnTo              string `xml:"CONTACTNAME>RETURNTO"`
	Paymethodkey          string `xml:"PAYMETHODKEY"`
	Mergepaymentreq       string `xml:"MERGEPAYMENTREQ"`
	Paymentnotify         bool   `xml:"PAYMENTNOTIFY"`
	Billingtype           string `xml:"BILLINGTYPE"`
	Paymentpriority       string `xml:"PAYMENTPRIORITY"`
	Termname              string `xml:"TERMNAME"`
	Displaytermdiscount   string `xml:"DISPLAYTERMDISCOUNT"`
	Achenabled            string `xml:"ACHENABLED"`
	Achbankroutingnumber  string `xml:"ACHBANKROUTINGNUMBER"`
	Achaccountnumber      string `xml:"ACHACCOUNTNUMBER"`
	Achaccounttype        string `xml:"ACHACCOUNTTYPE"`
	Achremittancetype     string `xml:"ACHREMITTANCETYPE"`
	Vendoraccountno       string `xml:"VENDORACCOUNTNO"`
	Displayacctnocheck    string `xml:"DISPLAYACCTNOCHECK"`
	Objectrestriction     string `xml:"OBJECTRESTRICTION"`
	Restrictedlocations   string `xml:"RESTRICTEDLOCATIONS"`
	Restricteddepartments string `xml:"RESTRICTEDDEPARTMENTS"`
	Customfield1          string `xml:"CUSTOMFIELD1"`
}

type Vendors struct {
	Client
}

func (vendors Vendors) makeRequest(list interface{}) ([]Vendor, string, error) {
	get := Function{
		ControlID: "testControlID",
		Method:    list,
	}
	// Create a new request using the Client
	req, err := vendors.Client.NewRequest(get)
	if err != nil {
		return nil, "", err
	}

	resp, err := vendors.Client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	var body Response
	if err = xml.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, "", err
	}

	// Check the response for errors
	if err = vendors.Client.CheckResponseErrors(body); err != nil {
		return nil, "", err
	}

	// TODO pull out status code and body status checks into client
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
		bodyString := string(body)
		return nil, "", fmt.Errorf("non-200 status code: %d, error: %s", resp.StatusCode, bodyString)
	}

	return body.Operation.Result.Data.Vendors, body.Operation.Result.Data.ResultId, nil
}
func (vendors Vendors) List(limit int) ([]Vendor, error) {
	list := ReadByQuery{
		Object:   "vendor",
		Fields:   "VENDORID,NAME,STATUS",
		Query:    "",
		Pagesize: 1000,
	}

	vendorsList, next, err := vendors.makeRequest(list)
	if err != nil {
		return vendorsList, err
	}

	if len(vendorsList) >= limit {
		return vendorsList[:limit], nil
	}

	for next != "" {
		list := ReadMore{
			ResultId: next,
		}
		var err error
		var vendorsPage []Vendor
		vendorsPage, next, err = vendors.makeRequest(list)
		if err != nil {
			return vendorsList, err
		}
		for _, bill := range vendorsPage {
			vendorsList = append(vendorsList, bill)
			if len(vendorsList) >= limit {
				return vendorsList[:limit], nil
			}
		}
	}

	return vendorsList, nil
}
