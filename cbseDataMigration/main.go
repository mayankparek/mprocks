package main

import (
	"cbseMigrationProject/lib"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		os.Exit(0)
	}
	config := lib.Config{}

	err = json.Unmarshal([]byte(configFile), &config)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(config)
	lines, err := ReadCsv(config.SourceFileName)
	if err != nil {
		panic(err)
	}
	// Create a Resty Client
	client := resty.New()
	for k, v := range config.Header {
		client.SetHeader(k, v)
	}
	// Loop through lines & turn into object

	postBodyRaw := `{
		"txnId": "f7f1469c-29b0-4325-9dfc-c567200a70f7",
		"format": "xml",
		"certificateParameters": {
			"year": "%s",
			"rollno": "%s",
			"FullName": "%s"
		},
		"consentArtifact": {
			"consent": {
				"consentId": "ea9c43aa-7f5a-4bf3-a0be-e1caa24737ba",
				"timestamp": "2021-08-07T13:11:49.750Z",
				"dataConsumer": {
					"id": "string"
				},
				"dataProvider": {
					"id": "string"
				},
				"purpose": {
					"description": "string"
				},
				"user": {
					"idType": "string",
					"idNumber": "string",
					"mobile": "9999999999",
					"email": "dummy@yopmail.com"
				},
				"data": {
					"id": "string"
				},
				"permission": {
					"access": "string",
					"dateRange": {
						"from": "2021-08-07T13:11:49.750Z",
						"to": "2021-08-07T13:11:49.750Z"
					},
					"frequency": {
						"unit": "string",
						"value": 0,
						"repeats": 0
					}
				}
			},
			"signature": {
				"signature": "string"
			}
		}
	}`

	lib.PrepareHeadersInExcel(config)
	for indx, line := range lines {
		if indx == 0 {
			continue
		}
		<-time.After(2 * time.Second)
		fmt.Println("Processing record ................ ", indx)

		data := lib.SourceCsvLine{
			ACPDC_rollno: line[0],
			RollNo:       line[1],
			Name:         line[2],
			PassYear:     line[3],
		}
		requestBody := fmt.Sprintf(postBodyRaw, data.PassYear, data.RollNo, data.Name)

		response, err := client.R().SetBody(requestBody).Post(config.Endpoint)
		if err != nil {
			fmt.Println("ERROR: ----------1 ", err)
			os.Exit(2)
			return
		}
		statusCode := response.StatusCode()
		if !(statusCode >= 200 && statusCode < 300) {
			fmt.Println("ERROR: ----------2 ", string(response.Body()), data, err)
			if response.Body() == nil {
				lib.DumpError(config, data, err.Error())
				continue
			}
			lib.DumpError(config, data, string(response.Body()))
			continue
		}
		xmlData := lib.Certificate{}
		err = xml.Unmarshal(response.Body(), &xmlData)

		if err != nil {
			fmt.Println("ERROR: ----------3 ", err)
			os.Exit(4)
			return
		}

		for _, subject := range xmlData.CertificateData.Performance.Subjects.Subject {
			lib.PrepareDataInExcel(config, data, xmlData, subject)
		}

	}
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
