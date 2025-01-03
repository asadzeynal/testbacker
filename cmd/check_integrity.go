package main

import (
	dataintegritychecker "asadzeynal/testbacker/internal/data_integrity_checker"
	"fmt"
	"log"
	"os"
	"time"
)

const ERR_INVALID_PATH = "invalid filepath, please provide a valid csv file path"

func main() {
	csvFileName := os.Args[1]
	if csvFileName == "" {
		log.Fatalf(ERR_INVALID_PATH)
	}

	c := dataintegritychecker.New(time.Second)
	res, err := c.Check(csvFileName, "Unix Timestamp")
	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, v := range res {
		fmt.Printf("inconsistency between timestamps %s, %s", v.PrevRow[1], v.CurrRow[1])
	}
}
