package intacct

import (
	"encoding/xml"
	"strconv"
	"time"
)

const (
	Version3 = "3.0"
)

type Control struct {
	XMLName           xml.Name `xml:"control"`
	Status            string   `xml:"status,omitempty"` // Only response
	SenderID          string   `xml:"senderid"`
	Password          string   `xml:"password"`
	ControlID         string   `xml:"controlid"`
	UniqueID          bool     `xml:"uniqueid"`
	DTDVersion        string   `xml:"dtdversion"`
	Includewhitespace bool     `xml:"includewhitespace"`
}

func NewControlV2(sender, password string) Control {
	return Control{
		SenderID:   sender,
		Password:   password,
		ControlID:  strconv.FormatInt(time.Now().Unix(), 10),
		DTDVersion: Version3,
	}
}

type Login struct {
	XMLName    xml.Name `xml:"login"`
	UserID     string   `xml:"userid"`
	CompanyID  string   `xml:"companyid"`
	Password   string   `xml:"password"`
	LocationID string   `xml:"locationid,omitempty"`
}

func NewLogin(user, company, password string) Login {
	return Login{
		UserID:    user,
		CompanyID: company,
		Password:  password,
	}
}

// TODO Relationship to login?
type Authentication struct {
	XMLName xml.Name `xml:"authentication"`
	Login   Login
	// TODO omitempty doesn't work for zero timestamps
	// SessionTimestamp time.Time `xml:"sessiontimestamp,omitempty"`
}
