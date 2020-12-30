package information

import (
	"pdfcpuSamples/samplepdfs"
	"testing"
)

const prefix = samplepdfs.DirPrefix

var testCases = []struct {
	dirName     string
	expectedNrOfPDFs int
}{
	{ // empty Dir
		prefix + "EmptyDir",
		0,
	},
	{ // OnePDF
		prefix + "OnePDF",
		1,
	},
	{ // Four files, TWO PDFs
		dirName: prefix + "FourFilesTwoPDFs",
		expectedNrOfPDFs: 2,
	},
}
func TestCountPDFsInDir(t *testing.T) {


	for _,d  := range testCases {
		got := CountPDFsInDir(d.dirName)
		if  got != d.expectedNrOfPDFs {
			t.Errorf("FAIL: directory %s expected %v, got %v PDF files", d.dirName, d.expectedNrOfPDFs, got )
		}
	}

}
