package logger

import (
	"fmt"
	"strings"
	"time"
)

// Panic calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Panic(args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(bold, red, "[PANIC]", reset)
	} else {
		logStr = fmt.Sprint("[PANIC]")
	}
	logStr += fmt.Sprint(" ")
	if clockAtive {
		logStr += fmt.Sprint(time.Now().Format(logClock))
		logStr += fmt.Sprint(" ")
	}
	logStr += fmt.Sprint(args...)
	mutex.Unlock()
	panic(logStr)
}

// Panicf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Panicf(format string, args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(bold, red, "[PANIC]", reset)
	} else {
		logStr = fmt.Sprint("[PANIC]")
	}
	logStr += fmt.Sprint(" ")
	if clockAtive {
		logStr += fmt.Sprint(time.Now().Format(logClock))
		logStr += fmt.Sprint(" ")
	}
	logStr += fmt.Sprintf(format, args...)
	if strings.Contains(format[len(format)-1:], "\n") {
		logStr += fmt.Sprint(format[:len(format)-1])
	}
	mutex.Unlock()
	panic(logStr)
}

// Panicln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Panicln(args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(bold, red, "[PANIC]", reset)
	} else {
		logStr = fmt.Sprint("[PANIC]")
	}
	logStr += fmt.Sprint(" ")
	if clockAtive {
		logStr += fmt.Sprint(time.Now().Format(logClock))
		logStr += fmt.Sprint(" ")
	}
	logStr += fmt.Sprint(args...)
	mutex.Unlock()
	panic(logStr)
}
