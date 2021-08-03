package parser

import (
	"log"
	"strconv"
	"strings"
)

func toINT(s string) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return int64(i)
}

func toINT16(s string) int16 {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalln(err)
	}
	return int16(i)
}

func toFLO(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return float64(i)
}

func trim(s string) string {
	return strings.Trim(s, " \t")
}
