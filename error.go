package logger

import (
	"fmt"
	"strings"
	"time"
)

// Error calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Error(args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(red, "[ERROR]", reset)
	} else {
		logStr = fmt.Sprint("[ERROR]")
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

// Errorf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(red, "[ERROR]", reset)
	} else {
		logStr = fmt.Sprint("[ERROR]")
	}
	logStr += fmt.Sprint(" ")
	if clockAtive {
		logStr += fmt.Sprint(time.Now().Format(logClock))
		logStr += fmt.Sprint(" ")
	}
	logStr += fmt.Sprintf(format, args...)
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

// Errorln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(red, "[ERROR]", reset)
	} else {
		logStr = fmt.Sprint("[ERROR]")
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

func errorInterno(args ...interface{}) {
	mutex.Lock()
	logStr := fmt.Sprint(red, "[ERROR]", reset)
	logStr += fmt.Sprint(" ")
	if clockAtive {
		logStr += fmt.Sprint(time.Now().Format(logClock))
		logStr += fmt.Sprint(" ")
	}
	logStr += fmt.Sprint(args...)
	logStr += fmt.Sprint("\n")
	mutex.Unlock()
	print(logStr)
}
