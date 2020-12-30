package information

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// CountPDFsInDir returns the number of PDF files in the given directory
func CountPDFsInDir(dirName string) int {

	var files []string
	var fCount int

	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {

		// exclude dirs
		if info.IsDir() {
			return nil
		}

		// count only PDF files
		if strings.ToUpper( filepath.Ext(path)) == ".PDF" {
			fCount++
			files = append(files, info.Name())
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}

	return fCount
}
