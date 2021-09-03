package lib

import (
	"fmt"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Config struct {
	Endpoint            string            `json:"endpoint"`
	SourceFileName      string            `json:"sourceFileName"`
	DestinationFileName string            `json:"destinationFileName"`
	Header              map[string]string `json:"header"`
}

func maskValue(s string) string {
	if s == "" {
		s = "N/A"
	}
	return s
}

func PrepareDataInExcel(config Config, data SourceCsvLine, xmlParsedData Certificate, subject Subject) {

	f, err := excelize.OpenFile(config.DestinationFileName)
	if err != nil {
		return
	}
	sheetName := "CBSE_DATA_PULL_REQUEST_RESPONSE"
	// index := f.NewSheet(sheetName)
	count := 1
	rows, err := f.Rows(sheetName)
	if err != nil {
		fmt.Println("Error while fetching row from sheet", err)
		return
	}
start:
	if rows.Next() {
		count++
		goto start
	}
	s := strconv.Itoa(count)
	f.SetCellValue(sheetName, "A"+s, data.ACPDC_rollno)
	f.SetCellValue(sheetName, "B"+s, data.RollNo)
	f.SetCellValue(sheetName, "C"+s, data.Name)
	f.SetCellValue(sheetName, "D"+s, data.PassYear)

	f.SetCellValue(sheetName, "E"+s, "N/A")

	f.SetCellValue(sheetName, "F"+s, maskValue(xmlParsedData.Number))
	f.SetCellValue(sheetName, "G"+s, maskValue(xmlParsedData.Status))
	f.SetCellValue(sheetName, "H"+s, maskValue(xmlParsedData.IssuedTo.Person.Name))
	f.SetCellValue(sheetName, "I"+s, maskValue(xmlParsedData.IssuedTo.Person.Dob))
	f.SetCellValue(sheetName, "J"+s, maskValue(xmlParsedData.IssuedTo.Person.Gender))
	f.SetCellValue(sheetName, "K"+s, maskValue(xmlParsedData.IssuedTo.Person.Phone))
	f.SetCellValue(sheetName, "L"+s, maskValue(xmlParsedData.IssuedTo.Person.Email))
	f.SetCellValue(sheetName, "M"+s, maskValue(xmlParsedData.CertificateData.School.Name))
	f.SetCellValue(sheetName, "N"+s, maskValue(xmlParsedData.CertificateData.School.Code))
	f.SetCellValue(sheetName, "O"+s, maskValue(xmlParsedData.CertificateData.School.RegionName))
	f.SetCellValue(sheetName, "P"+s, maskValue(xmlParsedData.CertificateData.School.RegionCode))
	f.SetCellValue(sheetName, "Q"+s, maskValue(xmlParsedData.CertificateData.Examination.Name))
	f.SetCellValue(sheetName, "R"+s, maskValue(xmlParsedData.CertificateData.Examination.CenterCode))
	f.SetCellValue(sheetName, "S"+s, maskValue(xmlParsedData.CertificateData.Examination.Month))
	f.SetCellValue(sheetName, "T"+s, maskValue(xmlParsedData.CertificateData.Examination.Year))
	f.SetCellValue(sheetName, "U"+s, maskValue(xmlParsedData.CertificateData.Performance.Result))
	f.SetCellValue(sheetName, "V"+s, maskValue(xmlParsedData.CertificateData.Performance.MarksTotal))
	f.SetCellValue(sheetName, "W"+s, maskValue(xmlParsedData.CertificateData.Performance.MarksMax))
	f.SetCellValue(sheetName, "X"+s, maskValue(xmlParsedData.CertificateData.Performance.Percentage))
	f.SetCellValue(sheetName, "Y"+s, maskValue(xmlParsedData.CertificateData.Performance.Cgpa))
	f.SetCellValue(sheetName, "Z"+s, maskValue(xmlParsedData.CertificateData.Performance.CgpaMax))
	f.SetCellValue(sheetName, "AA"+s, maskValue(subject.Name))
	f.SetCellValue(sheetName, "AB"+s, maskValue(subject.Code))
	f.SetCellValue(sheetName, "AC"+s, maskValue(subject.MarksTheory))
	f.SetCellValue(sheetName, "AD"+s, maskValue(subject.MarksMaxTheory))
	f.SetCellValue(sheetName, "AE"+s, maskValue(subject.MarksPractical))
	f.SetCellValue(sheetName, "AF"+s, maskValue(subject.MarksMaxPractical))
	f.SetCellValue(sheetName, "AG"+s, maskValue(subject.MarksTotal))
	f.SetCellValue(sheetName, "AH"+s, maskValue(subject.MarksMax))
	f.SetCellValue(sheetName, "AI"+s, maskValue(subject.Gp))
	f.SetCellValue(sheetName, "AJ"+s, maskValue(subject.GpMax))
	f.SetCellValue(sheetName, "AK"+s, maskValue(subject.Grade))

	f.SetRowHeight(sheetName, count, 20)
	// f.SetActiveSheet(index)
	if err := f.SaveAs(config.DestinationFileName); err != nil {
		fmt.Println(err)
	}
}

func DumpError(config Config, data SourceCsvLine, errResponse string) {

	f, err := excelize.OpenFile(config.DestinationFileName)
	if err != nil {
		return
	}
	sheetName := "CBSE_DATA_PULL_REQUEST_RESPONSE"
	// index := f.NewSheet(sheetName)
	count := 1
	rows, err := f.Rows(sheetName)
	if err != nil {
		fmt.Println("Error while fetching row from sheet", err)
		return
	}
start:
	if rows.Next() {
		count++
		goto start
	}
	s := strconv.Itoa(count)
	f.SetCellValue(sheetName, "A"+s, data.ACPDC_rollno)
	f.SetCellValue(sheetName, "B"+s, data.RollNo)
	f.SetCellValue(sheetName, "C"+s, data.Name)
	f.SetCellValue(sheetName, "D"+s, data.PassYear)

	f.SetCellValue(sheetName, "E"+s, errResponse)

	f.SetRowHeight(sheetName, count, 20)
	// f.SetActiveSheet(index)
	if err := f.SaveAs(config.DestinationFileName); err != nil {
		fmt.Println(err)
	}
}

func PrepareHeadersInExcel(config Config) {
	if _, err := os.Stat(config.DestinationFileName); !os.IsNotExist(err) {
		os.Remove(config.DestinationFileName)
		// path/to/whatever does not exist
	}
	f := excelize.NewFile()
	sheetName := "CBSE_DATA_PULL_REQUEST_RESPONSE"
	index := f.NewSheet(sheetName)

	f.SetCellValue(sheetName, "A1", "Req_ACPDC_Roll_No")
	f.SetCellValue(sheetName, "B1", "Req_Roll_No")
	f.SetCellValue(sheetName, "C1", "Req_Name")
	f.SetCellValue(sheetName, "D1", "Req_PassYear")

	f.SetCellValue(sheetName, "E1", "ERROR")

	f.SetCellValue(sheetName, "F1", "Res_Number")
	f.SetCellValue(sheetName, "G1", "Res_Status")
	f.SetCellValue(sheetName, "H1", "Res_Person_Name")
	f.SetCellValue(sheetName, "I1", "Res_Person_DOB")
	f.SetCellValue(sheetName, "J1", "Res_Person_Gender")
	f.SetCellValue(sheetName, "K1", "Res_Person_Phone")
	f.SetCellValue(sheetName, "L1", "Res_Person_Email")
	f.SetCellValue(sheetName, "M1", "Res_School_Name")
	f.SetCellValue(sheetName, "N1", "Res_School_Code")
	f.SetCellValue(sheetName, "O1", "Res_School_RegionName")
	f.SetCellValue(sheetName, "P1", "Res_School_RegionCode")
	f.SetCellValue(sheetName, "Q1", "Res_Exam_Name")
	f.SetCellValue(sheetName, "R1", "Res_Exam_CenterCode")
	f.SetCellValue(sheetName, "S1", "Res_Exam_Month")
	f.SetCellValue(sheetName, "T1", "Res_Exam_Year")
	f.SetCellValue(sheetName, "U1", "Res_Performance_Result")
	f.SetCellValue(sheetName, "V1", "Res_Performance_TotalMarks")
	f.SetCellValue(sheetName, "W1", "Res_Performance_MarksMax")
	f.SetCellValue(sheetName, "X1", "Res_Performance_Percentage")
	f.SetCellValue(sheetName, "Y1", "Res_Performance_Cgpa")
	f.SetCellValue(sheetName, "Z1", "Res_Performance_CgpaMax")
	f.SetCellValue(sheetName, "AA1", "Res_Subject_Name")
	f.SetCellValue(sheetName, "AB1", "Res_Subject_Code")
	f.SetCellValue(sheetName, "AC1", "Res_Subject_MarksTheory")
	f.SetCellValue(sheetName, "AD1", "Res_Subject_MarksMaxTheory")
	f.SetCellValue(sheetName, "AE1", "Res_Subject_MarksPractical")
	f.SetCellValue(sheetName, "AF1", "Res_Subject_MarksMaxPractical")
	f.SetCellValue(sheetName, "AG1", "Res_Subject_MarksTotal")
	f.SetCellValue(sheetName, "AH1", "Res_Subject_MarksMax")
	f.SetCellValue(sheetName, "AI1", "Res_Subject_Gp")
	f.SetCellValue(sheetName, "AJ1", "Res_Subject_GpMax")
	f.SetCellValue(sheetName, "AK1", "Res_Subject_Grade")

	f.SetRowHeight(sheetName, 1, 20)
	f.SetColWidth(sheetName, "A", "D", 20)
	f.SetActiveSheet(index)
	if err := f.SaveAs(config.DestinationFileName); err != nil {
		fmt.Println(err)
	}
}
