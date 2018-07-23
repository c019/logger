package logger

import (
	"fmt"
	"strings"
	"time"
)

// Debug calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Debug(args ...interface{}) {
	if debugAtive {
		mutex.Lock()
		logStr := fmt.Sprint(green, "[DEBUG]", reset)
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
}

// Debugf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, args ...interface{}) {
	if debugAtive {
		mutex.Lock()
		logStr := fmt.Sprint(green, "[DEBUG]", reset)
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
}

// Debugln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Debugln(args ...interface{}) {
	if debugAtive {
		mutex.Lock()
		logStr := fmt.Sprint(green, "[DEBUG]", reset)
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
}
