package main

import (
	"os"
	"log"
	"flag"
)

var log0 *log.Logger = NewNullLog()
var log1 *log.Logger = NewNullLog()
var log2 *log.Logger = NewNullLog()
var log3 *log.Logger = NewNullLog()
var log4 *log.Logger = NewNullLog()
var log5 *log.Logger = NewNullLog()
var log6 *log.Logger = NewNullLog()
var log7 *log.Logger = NewNullLog()
var log8 *log.Logger = NewNullLog()
var log9 *log.Logger = NewNullLog()

var logLevel int

func init() {
	flag.IntVar(&logLevel, "loglevel", 5, "Sets the amount of messages from the server that are logged. 5 is a sane default, 10 will be too fast to read, and 1 will almost never give you any messages")
}

func NewNullLog() *log.Logger {
	nfile, err := os.Open(os.DevNull, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("WTF. Could not open null device!")
	}
	nlog = log.New(nfile, "", 0)
	return nlog
}

func NewNormalLog() *log.Logger {
	return log.New(os.Stderr, "", 0)
}

func InitLogs() {
	if logLevel > 9 {
		log9 = NewNormalLog()
	}
	if logLevel > 8 {
		log8 = NewNormalLog()
	}
	if logLevel > 7 {
		log7 = NewNormalLog()
	}
	if logLevel > 6 {
		log6 = NewNormalLog()
	}
	if logLevel > 5 {
		log5 = NewNormalLog()
	}
	if logLevel > 4 {
		log4 = NewNormalLog()
	}
	if logLevel > 3 {
		log3 = NewNormalLog()
	}
	if logLevel > 2 {
		log2 = NewNormalLog()
	}
	if logLevel > 1 {
		log1 = NewNormalLog()
	}
	if logLevel > 0 {
		log0 = NewNormalLog()
	}
}