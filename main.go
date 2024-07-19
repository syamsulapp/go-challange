package main

import (
	"github.com/thecodingmachine/gotenberg-go-client/v7"
)

func main() {
	c := &gotenberg.Client{Hostname: "http://localhost:3000"}
	docs, _ := gotenberg.NewDocumentFromPath("example.docx", "D:/go-challange/docs/example.docx")
	req := gotenberg.NewOfficeRequest(docs)
	c.Store(req, "D:/go-challange/pdf/example.pdf")
}
