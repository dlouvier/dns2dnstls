package main

import "log"

// Debugger - Small debugger function to save lines of code
func Debugger(e error) {
	if e != nil {
		log.Fatal(e.Error())
	}
}
