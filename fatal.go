package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Fatal calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Fatal(args ...interface{}) {
	mutex.Lock()
	logStr := fmt.Sprint(bold, red, "[FATAL]", reset)
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
	os.Exit(statusExit)
}

// Fatalf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, args ...interface{}) {
	mutex.Lock()
	logStr := fmt.Sprint(bold, red, "[FATAL]", reset)
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
	os.Exit(statusExit)
}

// Fatalln calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Fatalln(args ...interface{}) {
	mutex.Lock()
	logStr := fmt.Sprint(bold, red, "[FATAL]", reset)
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
	os.Exit(statusExit)
}
