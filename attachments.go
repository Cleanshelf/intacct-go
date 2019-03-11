package intacct

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Supdoc struct {
	XMLName xml.Name `xml:"supdoc"`

	Recordno     string `xml:"recordno"`
	Supdocid     string `xml:"supdocid"`
	Supdocname   string `xml:"supdocname"`
	Folder       string `xml:"folder"`
	Description  string `xml:"description"`
	Creationdate string `xml:"creationdate"`
	Createdby    string `xml:"createdby"`

	Attachments AttachmentsXML `xml:"attachments"`
}

type AttachmentsXML struct {
	XMLName xml.Name `xml:"attachments"`

	Attachment []Attachment `xml:"attachment"`
}

type Attachment struct {
	XMLName xml.Name `xml:"attachment"`

	AttachmentName string `xml:"attachmentname"`
	AttachmentType string `xml:"attachmenttype"`
	AttachmentData string `xml:"attachmentdata"`
}

type Attachments struct {
	Client
}

func (attachments Attachments) makeRequest(list interface{}) ([]Supdoc, error) {
	var docs []Supdoc
	get := &Function{
		ControlID: "testControlID",
		Method:    list,
	}
	// Create a new request using the Client
	req, err := attachments.Client.NewRequest(get)
	if err != nil {
		return docs, err
	}

	resp, err := attachments.Client.Do(req)
	if err != nil {
		return docs, err
	}
	defer resp.Body.Close()

	var body Response
	if err = xml.NewDecoder(resp.Body).Decode(&body); err != nil {
		return docs, err
	}
	// Check the response for errors
	if err = attachments.Client.CheckResponseErrors(body); err != nil {
		return docs, err
	}

	// TODO pull out status code and body status checks into client
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return docs, err
		}
		bodyString := string(body)
		return docs, fmt.Errorf("non-200 status code: %d, error: %s", resp.StatusCode, bodyString)
	}
	return body.Operation.Result.Data.Supdocs, nil
}

func (attachments Attachments) Get(key string) (Supdoc, error) {
	query := Get{
		XMLName: xml.Name{Local: "get"},
		Object:  "supdoc",
		Key:     key,
	}

	docs, err := attachments.makeRequest(query)
	if err != nil {
		return Supdoc{}, err
	}

	if len(docs) > 0 {
		return docs[0], nil
	}

	return Supdoc{}, nil
}
