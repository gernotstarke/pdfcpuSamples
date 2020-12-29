package pageCount

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"errors"
)

// countPagesOfODFFile counts the pages of fileName.
func CountPagesOfPDFFile( pdfFileName string) (int, error) {

	// use default configuration for pdfcpu ("nil")
	err := api.ValidateFile( pdfFileName, nil)

	if err != nil {
		return -1, errors.New("CountPagesOfPDFFile: invalid PDF: %v" )
	}

	return api.PageCountFile(pdfFileName)

}
