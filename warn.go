package logger

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Warn calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Warn(args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(yellow, "[WARN]", reset)
	} else {
		logStr = fmt.Sprint("[WARN]")
	}
	logStr += fmt.Sprint(" ")
	if clockAtive {
		logStr += fmt.Sprint(time.Now().Format(logClock))
		logStr += fmt.Sprint(" ")
	}
	logStr += fmt.Sprint(args...)
	logStr += fmt.Sprint("\n")
	mutex.Unlock()
	if fileSave {
		save(logStr)
	} else {
		print(logStr)
	}
}

// Warnf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(yellow, "[WARN]", reset)
	} else {
		logStr = fmt.Sprint("[WARN]")
	}
	logStr += fmt.Sprint(" ")
	if clockAtive {
		logStr += fmt.Sprint(time.Now().Format(logClock))
		logStr += fmt.Sprint(" ")
	}
	log.Printf(format, args...)
	if !strings.Contains(format[len(format)-1:], "\n") {
		logStr += fmt.Sprint("\n")
	}
	mutex.Unlock()
	if fileSave {
		save(logStr)
	} else {
		print(logStr)
	}
}

// Warnln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Warnln(args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(yellow, "[WARN]", reset)
	} else {
		logStr = fmt.Sprint("[WARN]")
	}
	logStr += fmt.Sprint(" ")
	if clockAtive {
		logStr += fmt.Sprint(time.Now().Format(logClock))
		logStr += fmt.Sprint(" ")
	}
	logStr += fmt.Sprintln(args...)
	mutex.Unlock()
	if fileSave {
		save(logStr)
	} else {
		print(logStr)
	}
}
