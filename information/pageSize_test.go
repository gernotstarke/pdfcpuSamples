package information

import (
	"log"
	samplePDFs "pdfcpuSamples/samplepdfs"
	"testing"
)

func TestPageSize(t *testing.T) {

	var oneFileDims []dim

	// get page dimensions for single-page file
	log.Printf( "calling GetPageDimensions")
	oneFileDims = GetPageDimensions(samplePDFs.OnePageFile);

	if len(oneFileDims) != 1 {
		log.Printf("FAIL: length of page dimension list expected to be 1, but was %v", len(oneFileDims))
	}

		log.Print( "Height = ", oneFileDims[0].height)
	log.Print( "Width = ", oneFileDims[0].width)

	if oneFileDims[0].height != PageSizes["A4"].Height {
		t.Errorf("FAIL: page height should be A4, but is %v", oneFileDims[0].height)
	}
}

