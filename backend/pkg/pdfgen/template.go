package pdfgen

import (
	"fmt"
	"github.com/go-pdf/fpdf"
	"path/filepath"
	"strconv"
)

const pageWidth = 210.0

func NewFrontendTemplate(invoice Invoice) *fpdf.Fpdf {
	marginX := 10.0
	marginY := 20.0

	pdf := fpdf.New("P", "mm", "A4", filepath.Join("pkg", "pdfgen", "assets", "fonts"))
	pdf.AddUTF8Font("calibri", "R", "calibri-regular.ttf")
	pdf.AddUTF8Font("calibri", "B", "calibri-bold.ttf")
	pdf.SetMargins(marginX, marginY, marginX)
	pdf.AddPage()

	imagePath := filepath.Join("pkg", "pdfgen", "assets", "img", "frontend.png")

	fmt.Println(imagePath)

	// if _, err := os.Stat(imagePath); os.IsNotExist(err) {
	// 	return fmt.Errorf("file does not exist: %s", imagePath)
	// }
	// img := pdf.GetImageInfo(imagePath)
	// fmt.Printf("%v",img)
	imageWidth := 50.0
	pdf.ImageOptions(imagePath, marginX, marginY-5, imageWidth, 0, false, fpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "")

	titleFontSize := 36.0
	pdf.SetFont("calibri", "B", titleFontSize)

	titleStr := invoice.InvoiceTitle
	titleWidth := pdf.GetStringWidth(titleStr)

	fmt.Println(titleWidth)
	pdf.SetX(pageWidth - marginX - titleWidth)
	pdf.Cell(0, 0, titleStr)

	pdf.SetXY(pageWidth-marginX-titleWidth, titleFontSize-8)
	pdf.SetFont("calibri", "B", 11)
	pdf.Cell(0, 0, "#"+strconv.Itoa(invoice.InvoiceNumber))

	// BillFrom
	pdf.SetFont("calibri", "R", 11)

	lineHeight := 5.0
	x := marginX
	y := 45.0

	pdf.SetXY(x, y)
	pdf.Cell(150, lineHeight, invoice.BillFrom.Name)

	y += lineHeight
	pdf.SetXY(x, y)
	pdf.Cell(150, lineHeight, "Address: "+invoice.BillFrom.Address)

	y += lineHeight
	pdf.SetXY(x, y)
	pdf.Cell(150, lineHeight, "Bank: "+invoice.BillFrom.Bank)

	y += lineHeight
	pdf.SetXY(x, y)
	pdf.Cell(150, lineHeight, "Swift: "+invoice.BillFrom.Swift)

	y += lineHeight
	pdf.SetXY(x, y)
	pdf.Cell(150, lineHeight, "BankNr: "+invoice.BillFrom.BankNr)

	pdf.SetFont("calibri", "B", 12)
	y += 15
	initialY := y
	pdf.SetXY(x, y)
	pdf.Cell(150, lineHeight, "Bill To")

	pdf.SetFont("calibri", "R", 11)

	y += lineHeight
	pdf.SetXY(x, y)
	pdf.Cell(150, lineHeight, invoice.BillTo.Name)

	y += lineHeight
	pdf.SetXY(x, y)
	pdf.MultiCell(50, lineHeight, "Address: "+invoice.BillTo.Address, "0", "", false)

	//Invoice number
	xRight := 150.0 // Set this to the X position of the right column
	y = initialY
	pdf.SetXY(xRight, y)
	pdf.Cell(150, lineHeight, "Invoice#: "+strconv.Itoa(invoice.InvoiceNumber))

	y += lineHeight
	pdf.SetXY(xRight, y)
	pdf.Cell(150, lineHeight, "Invoice Date: "+invoice.InvoiceDate)
	// Table

	y += 40.
	rowHeight := 10.
	idWidth := 15.
	descWidth := 115.
	timeWidth := 20.
	totalWidth := 40.

	pdf.SetXY(10, y)
	pdf.SetFillColor(0, 0, 0)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont("calibri", "B", 12)
	pdf.CellFormat(idWidth, rowHeight, "Id", "", 0, "C", true, 0, "")

	pdf.SetXY(25, y)
	pdf.CellFormat(descWidth, rowHeight, "Description", "", 0, "C", true, 0, "")

	pdf.SetXY(140, y)
	pdf.CellFormat(timeWidth, rowHeight, "Time", "", 0, "C", true, 0, "")

	pdf.SetXY(160, y)
	pdf.CellFormat(totalWidth, rowHeight, "Total", "", 0, "C", true, 0, "")

	y += rowHeight
	pdf.SetTextColor(0, 0, 0)
	for i, item := range invoice.Works {
		pdf.SetXY(10, y)
		pdf.SetFont("calibri", "R", 12)
		pdf.SetFillColor(255, 255, 255) // Set fill color to white

		// Print Id
		pdf.CellFormat(idWidth, rowHeight, fmt.Sprint(i+1), "0", 0, "C", true, 0, "")

		pdf.SetX(25) // Move X position for the next cell
		// Print Description
		pdf.CellFormat(descWidth, rowHeight, item.Description, "0", 0, "L", true, 0, "")

		pdf.SetX(140) // Move X position for the next cell
		// Print Time
		pdf.CellFormat(timeWidth, rowHeight, item.WorkTime, "0", 0, "C", true, 0, "")

		pdf.SetX(160) // Move X position for the next cell
		// Print Total
		pdf.CellFormat(totalWidth, rowHeight, "", "0", 0, "C", true, 0, "")

		pdf.Line(10, y, 200, y)
		y += rowHeight
	}
	y += 1

	pdf.Line(10, y, 200, y)

	y += 0.1

	pdf.SetXY(160, y)
	pdf.SetFillColor(249, 249, 251)
	pdf.CellFormat(totalWidth, rowHeight, "total", "0", 0, "C", true, 0, "")

	return pdf
}
