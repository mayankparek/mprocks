package lib

import "encoding/xml"

type SourceCsvLine struct {
	ACPDC_rollno string
	RollNo       string
	Name         string
	PassYear     string
}

type Subject struct {
	Text              string `xml:",chardata"`
	Name              string `xml:"name,attr"`
	Code              string `xml:"code,attr"`
	MarksTheory       string `xml:"marksTheory,attr"`
	MarksMaxTheory    string `xml:"marksMaxTheory,attr"`
	MarksPractical    string `xml:"marksPractical,attr"`
	MarksMaxPractical string `xml:"marksMaxPractical,attr"`
	MarksTotal        string `xml:"marksTotal,attr"`
	MarksMax          string `xml:"marksMax,attr"`
	Gp                string `xml:"gp,attr"`
	GpMax             string `xml:"gpMax,attr"`
	Grade             string `xml:"grade,attr"`
}

type Certificate struct {
	XMLName  xml.Name `xml:"Certificate"`
	Text     string   `xml:",chardata"`
	Number   string   `xml:"number,attr"`
	Status   string   `xml:"status,attr"`
	IssuedTo struct {
		Text   string `xml:",chardata"`
		Person struct {
			Name   string `xml:"name,attr"`
			Dob    string `xml:"dob,attr"`
			Gender string `xml:"gender,attr"`
			Phone  string `xml:"phone,attr"`
			Email  string `xml:"email,attr"`
		} `xml:"Person"`
	} `xml:"IssuedTo"`
	CertificateData struct {
		Text   string `xml:",chardata"`
		School struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			Code       string `xml:"code,attr"`
			RegionName string `xml:"regionName,attr"`
			RegionCode string `xml:"regionCode,attr"`
		} `xml:"School"`
		Examination struct {
			Text       string `xml:",chardata"`
			Name       string `xml:"name,attr"`
			CenterCode string `xml:"centerCode,attr"`
			Month      string `xml:"month,attr"`
			Year       string `xml:"year,attr"`
		} `xml:"Examination"`
		Performance struct {
			Text       string `xml:",chardata"`
			Result     string `xml:"result,attr"`
			MarksTotal string `xml:"marksTotal,attr"`
			MarksMax   string `xml:"marksMax,attr"`
			Percentage string `xml:"percentage,attr"`
			Cgpa       string `xml:"cgpa,attr"`
			CgpaMax    string `xml:"cgpaMax,attr"`
			Subjects   struct {
				Text    string    `xml:",chardata"`
				Subject []Subject `xml:"Subject"`
			} `xml:"Subjects"`
		} `xml:"Performance"`
	} `xml:"CertificateData"`
}

func main() {

}
