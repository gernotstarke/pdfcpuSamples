package information

/*
  Every page within a PDF might have different dimensions (aka page size).
  That makes e.g. stamping page numbers a little more difficult

  We therefore collect a slice of dimensions of every page, and check
  if some pages differ in size.

*/

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"log"
)

type dim struct {
	width  float64
	height float64
}


func GetPageDimensions(fname string) []dim {

	pageCount, _ := CountPagesOfPDFFile(fname)
	log.Printf("pagecount of %v was %v", fname, pageCount)

    var pageDimensions []dim
	var currentPageDim dim

	pdfcpuDims, err := api.PageDimsFile(fname)

	if err!=nil{
		log.Printf("Error %v", err)
	}

	for i := 0; i < pageCount; i++ {
		currentPageDim.width = pdfcpuDims[i].Width
		currentPageDim.height = pdfcpuDims[i].Height
		pageDimensions = append(pageDimensions, currentPageDim)
	}

	return pageDimensions

}
