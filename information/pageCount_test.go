package information

import (
	"fmt"
	samplePDFs "pdfcpuSamples/samplePDFs"
	"testing"
)

func TestPageCount(t *testing.T) {
	want := 1
	fmt.Println(samplePDFs.OnePageFile)

	if got, _ := CountPagesOfPDFFile(samplePDFs.OnePageFile); got != want {
		t.Errorf("CountPagesOfPDFFile() = %v, want %v", got, want)
	}

}
