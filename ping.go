package main

import (
	"strconv"
)

var callCount int

func Ping() string {

	callCount++
	result := "pong #" + strconv.Itoa(callCount)
	return result
}
