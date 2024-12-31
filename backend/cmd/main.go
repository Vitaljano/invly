package main

import (
	"fmt"
	"net/http"

	"github.com/Vitaljano/invly/backend/config"
	"github.com/Vitaljano/invly/backend/internal/auth"
	"github.com/Vitaljano/invly/backend/pkg/middleware"
	"github.com/Vitaljano/invly/backend/pkg/pdfgen"
)

func main() {
	conf := config.Load()

	mux := http.NewServeMux()

	//Handler
	auth.NewAuthHandler(mux)

	//Middleware
	stack := middleware.Chain(
		middleware.Logging,
	)

	in := pdfgen.Invoice{
		InvoiceTitle:  "INVOICE",
		InvoiceNumber: 130,
		InvoiceDate:   "2012-03-1",
		BillTo: pdfgen.BillTo{
			Name:    "Fake",
			Address: "Oassis dubai",
		},
		BillFrom: pdfgen.BillFrom{
			Name:    "Name Second Name",
			Address: "London 32-72",
			Bank:    "Revolut",
			Swift:   "sw4567324",
			BankNr:  "90435672355235",
		},
		Works: []pdfgen.Work{
			{
				Id:          23,
				Description: "Done some tasks",
				Price:       2.700,
				WorkTime:    "7.7h",
			},
			{
				Id:          24,
				Description: "Done some tasks",
				Price:       2.700,
				WorkTime:    "7.7h",
			},
			{
				Id:          23,
				Description: "Done some tasks",
				Price:       2.700,
				WorkTime:    "7.7h",
			},
		},
	}

	go pdfgen.GenerateInvoice(in, pdfgen.InvoiceOptions{
		SaveFolder: conf.PdfInvoiceFolder,
	})

	addr := fmt.Sprintf(":%s", conf.Port)
	server := http.Server{
		Addr:    addr,
		Handler: stack(mux),
	}

	fmt.Println("Server start on port", addr)
	server.ListenAndServe()
}
