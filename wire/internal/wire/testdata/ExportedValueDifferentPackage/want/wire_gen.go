// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	os "os"
)

// Injectors from wire.go:

func injectedFile() *os.File {
	file := _wireFileValue
	return file
}

var (
	_wireFileValue = os.Stdout
)
