package dataintegritychecker

import (
	"asadzeynal/testbacker/pkg/parsers"
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

func (c *Checker) Check(filePath string) error {
	parser := parsers.NewCSVParser()
	parser.OpenFile(filePath)

	_, err := parser.NextRow()
	if err != nil {
		return err
	}
	return nil
}

func (c *Checker) isCorrectInterval(prev time.Time, curr time.Time) bool {
	return prev.Add(c.interval).Equal(curr)
}
