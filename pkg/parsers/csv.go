package parsers

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

type csvParser struct {
	reader *csv.Reader
}

func NewCSVParser() *csvParser {
	return &csvParser{}
}

func (p *csvParser) OpenFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %w", err)
	}

	defer file.Close()
	p.reader = csv.NewReader(file)
	return nil
}

func (p *csvParser) NextRow() ([]string, error) {
	row, err := p.reader.Read()
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, io.EOF
		}
		return nil, fmt.Errorf("could not read next row: %w", err)
	}
	return row, nil
}
