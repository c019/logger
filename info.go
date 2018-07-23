package logger

import (
	"fmt"
	"strings"
	"time"
)

// Info calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Info(args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(blue, "[INFO]", reset)
	} else {
		logStr = fmt.Sprint("[INFO]")
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

// Infof calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(blue, "[INFO]", reset)
	} else {
		logStr = fmt.Sprint("[INFO]")
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

// Infoln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) {
	mutex.Lock()
	var logStr string
	if !fileSave {
		logStr = fmt.Sprint(blue, "[INFO]", reset)
	} else {
		logStr = fmt.Sprint("[INFO]")
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
