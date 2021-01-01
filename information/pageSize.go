package information

/*
  Every page within a PDF might have different dimensions (aka page size).
  That makes e.g. stamping page numbers a little more difficult

  We therefore collect a slice of dimensions of every page, and check
  if some pages differ in size.

*/

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type dim struct {
	width  float64
	height float64
}

var pageDimensions []dim

func GetPageDimensions(fname string) []dim {

	pageCount, _ := CountPagesOfPDFFile(fname)

	pdfcpuDims, _ := api.PageDimsFile(fname)

	for i := 0; i < pageCount; i++ {
		pageDimensions[i].width = pdfcpuDims[i].Width
		pageDimensions[i].height = pdfcpuDims[i].Height
	}

	return pageDimensions

}
