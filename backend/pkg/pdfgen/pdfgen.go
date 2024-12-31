package pdfgen

import (
	"fmt"
	"os"
	"time"
)

type BillTo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type BillFrom struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Bank    string `json:"bank"`
	Swift   string `json:"swift"`
	BankNr  string `json:"bankNr"`
}

type Work struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	WorkTime    string  `json:"workTime"`
}

type Invoice struct {
	InvoiceTitle  string   `json:"invoiceTitle"`
	InvoiceNumber int      `json:"invoiceNumber"`
	InvoiceDate   string   `json:"invoiceDate"`
	BillTo        BillTo   `json:"billTo"`
	BillFrom      BillFrom `json:"billFrom"`
	Works         []Work   `json:"works"`
}

type InvoiceOptions struct {
	SaveFolder string
}

func GenerateInvoice(invoice Invoice, opt InvoiceOptions) error {

	pdf := NewFrontendTemplate(invoice)

	formattedDate := time.Now().Format("2006-01-02")
	folderName := time.Now().Format("2006")
	filename := fmt.Sprintf("%s/%s/invoice-%s-%d.pdf", opt.SaveFolder, folderName, formattedDate, invoice.InvoiceNumber)

	if err := os.MkdirAll(fmt.Sprintf("%s/%s", opt.SaveFolder, folderName), os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return err
	}

	err := pdf.OutputFileAndClose(filename)

	return err

}
