package information

import (
	"pdfcpuSamples/samplePDFs"
	"testing"
)

func TestValidate(t *testing.T) {
	// positive check: valid PDF should return true
	valid, _ := ValidatePDFFile(samplePDFs.OnePageFile)

	if valid != true {
		t.Errorf("ValidatePDFFile returned false, should be true: %v", samplePDFs.OnePageFile)
	}

	// negative check: non-existing file should return false
	nonExistingFileName := "/some/non/existing/path/r492jgfuth/file42"
	valid, _ = ValidatePDFFile(nonExistingFileName + ".pdf")

	if valid == true {
		t.Errorf("ValidatePDFFile returned true, should be false for nonexisting file: %v", nonExistingFileName)
	}

	// negative check: wrong extension should fail
	badExtensionFile := nonExistingFileName + ".XZY"
	valid, _ = ValidatePDFFile(badExtensionFile)
	if valid == true {
		t.Errorf("ValidatePDFFile returned true, should be false for non-PDF extension: %v", badExtensionFile)
	}

	// negative check: markdown file with pdf extension should fail
	valid, _ = ValidatePDFFile(samplePDFs.DisguisedMarkdownFile)
	if valid != false {
		t.Errorf("ValidatePDFFile returned true, should be false for bad PDF file: %v", samplePDFs.DisguisedMarkdownFile)
	}
}
