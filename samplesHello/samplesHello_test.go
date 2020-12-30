package samplesHello

import "testing"

func TestHello(t *testing.T) {
	want := "Welcome the pdfcpu example 'samplesHello'"
	if got := pdfCountHello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}

func TestAnother(t *testing.T) {
	want := "Hello, world."
	if got := quoteHello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
