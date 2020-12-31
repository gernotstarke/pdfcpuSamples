package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	nrPages, err := api.PageCountFile("./samplepdfs/sample-A4-portrait-3pgs.pdf")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ThreePagePDF returned %v pages", nrPages)
}
