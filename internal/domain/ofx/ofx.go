package ofx

import (
	"encoding/xml"
	"fmt"
	"os"
)

type OFX struct {
	XMLName      xml.Name       `xml:"OFX"`
	Signon       SignonResponse `xml:"SIGNONMSGSRSV1>SONRS"`
	BankResponse BankResponse   `xml:"BANKMSGSRSV1>STMTTRNRS>STMTRS"`
}

type SignonResponse struct {
	Org string `xml:"FI>ORG"`
	FID string `xml:"FI>FID"`
}

type BankResponse struct {
	AccountID    string           `xml:"BANKACCTFROM>ACCTID"`
	StartDate    string           `xml:"BANKTRANLIST>DTSTART"`
	EndDate      string           `xml:"BANKTRANLIST>DTEND"`
	Transactions []OFXTransaction `xml:"BANKTRANLIST>STMTTRN"`
	Balance      float64          `xml:"LEDGERBAL>BALAMT"`
}

type OFXTransaction struct {
	Type        string  `xml:"TRNTYPE"`
	Date        string  `xml:"DTPOSTED"`
	Amount      float64 `xml:"TRNAMT"`
	ID          string  `xml:"FITID"`
	Description string  `xml:"MEMO"`
}

func ParseOFXFile(path string) (*OFX, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var ofx OFX
	err = xml.Unmarshal(data, &ofx)
	if err != nil {
		return nil, err
	}

	return &ofx, nil
}
