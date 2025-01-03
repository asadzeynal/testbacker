package dataintegritychecker

import (
	"asadzeynal/testbacker/pkg/parsers"
	"io"
	"strconv"
	"time"
)

type Checker struct {
	interval time.Duration
}

func New(interval time.Duration) *Checker {
	return &Checker{
		interval: interval,
	}
}

type IncorrectDiffData struct {
	PrevRow []string
	CurrRow []string
}

func (c *Checker) Check(filePath string, timeColumnName string) ([]IncorrectDiffData, error) {
	parser := parsers.NewCSVParser()
	err := parser.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	titleRow, err := parser.NextRow()
	if err != nil {
		return nil, err
	}

	timeColIndex := -1
	for i, v := range titleRow {
		if v == timeColumnName {
			timeColIndex = i
		}
	}

	result := make([]IncorrectDiffData, 0)
	var prevRow, currRow []string
	var prevTime, currTime time.Time
	for {
		prevRow = currRow
		currRow, err = parser.NextRow()
		if err == io.EOF {
			break
		}

		parsedTime, err := strconv.Atoi(currRow[timeColIndex])
		if err != nil {
			return nil, err
		}

		currTime = time.Unix(int64(parsedTime), 0)

		if !c.isCorrectInterval(prevTime, currTime) {
			result = append(result, IncorrectDiffData{
				PrevRow: prevRow,
				CurrRow: currRow,
			})
		}
	}
	return result, err
}

func (c *Checker) isCorrectInterval(prev time.Time, curr time.Time) bool {
	// first row
	if prev.Equal(time.Time{}) {
		return true
	}
	return prev.Add(c.interval).Equal(curr)
}
