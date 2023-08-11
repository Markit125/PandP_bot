package readFiles

import (
	"bufio"
	"fmt"
	"os"
	errs "pandp/src/errors"
)

func ReadFileInSliceFromFile(slice *[]string, fileName string) error {

	if fileName == "" {
		return &errs.AbsentFileError{}
	}

	textFile, err := os.Open(fileName)
	if err != nil {
		return &errs.OpenFileError{
			Message: fmt.Sprintf("Cannot open file %s", fileName),
		}
	}

	defer func() {
		if err := textFile.Close(); err != nil {
			panic(err)
		}
	}()

	fileScanner := bufio.NewScanner(textFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if len(fileScanner.Text()) <= 6 {
			return &errs.ReadError{
				Message: fmt.Sprintf("Corrupted line in file %s", fileName),
			}
		}
		*slice = append(*slice, (fileScanner.Text())[6:])
	}

	return nil
}
