package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"log"
	"pdfcpuSamples/information"
)

func main() {
	nrPages, err := api.PageCountFile("./samplepdfs/sample-A4-portrait-3pgs.pdf")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ThreePagePDF returned %v pages", nrPages)

	fname := "./samplepdfs/sample-A4-portrait-1pg.pdf"

	dims, err := api.PageDimsFile(fname)

	if err!=nil{
		log.Println("FAIL: PageDimsFile returned error", err)
	}

	fmt.Printf( "Dimensions %v \n", dims )
	fmt.Printf( "type of dims %T \n", dims)
	fmt.Printf( "length of dims %v \n", len(dims))

	fmt.Printf( "width= %v \n", dims[0].Width)

	dims2 := information.GetPageDimensions( fname )
	fmt.Printf( "Dimensions %v \n", dims2 )
	fmt.Printf( "type of dims %T \n", dims2)
	fmt.Printf( "length of dims %v \n", len(dims2))
	fmt.Printf( "capacity of dims %v \n", cap(dims2))



}
