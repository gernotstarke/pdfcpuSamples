package pageCount

import "testing"

func TestPageCount(t *testing.T) {
	want := 1
	onePageFile := "./sample-A4-portrait-1pg.pdf"
	if got, _ := CountPagesOfPDFFile(onePageFile); got != want {
		t.Errorf("CountPagesOfPDFFile() = %v, want %v", got, want)
	}

}