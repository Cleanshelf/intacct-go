package intacct

import (
	"encoding/xml"
)

// TODO Or Common method and change the Name field...?
type GetList struct {
	XMLName xml.Name `xml:"get_list"`
	Object  string   `xml:"object,attr"`
	ListParams
}

type Get struct {
	XMLName xml.Name `xml:"get"`

	Object string `xml:"object,attr"`
	Key    string `xml:"key,attr"`
}

type ReadByQuery struct {
	XMLName  xml.Name `xml:"readByQuery"`
	Object   string   `xml:"object"`
	Fields   string   `xml:"fields"`
	Query    string   `xml:"query"`
	Pagesize int      `xml:"pagesize"`
}

type ReadMore struct {
	XMLName  xml.Name `xml:"readMore"`
	ResultId string   `xml:"resultId"`
}

// TODO params should include filters, sorts, max items?
// Filters and sorts can be attached to the params or passed directly to List
// TODO How to support multiple or nested filter expressions?
type ListParams struct {
	MaxItems uint64  `xml:"maxitems,attr"`
	Filter   Logical `xml:"filter,omitempty"`
	Sorts    Sorts   `xml:"sorts"`
}

// Merge will merge two ListParams - the given values (if non-zero) take
// precedence
func (l ListParams) Merge(other ListParams) ListParams {
	if other.MaxItems > 0 {
		l.MaxItems = other.MaxItems
	}

	l.Filter.Filters = append(l.Filter.Filters, other.Filter.Filters...)

	if len(other.Sorts) > 0 {
		l.Sorts = other.Sorts
	}
	return l
}
