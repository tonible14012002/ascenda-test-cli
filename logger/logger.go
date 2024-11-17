package logger

import "fmt"

type ILogger interface {
	Print(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type logger struct{}

func New() ILogger {
	return &logger{}
}

func (l *logger) Print(msg string) {
	fmt.Println(msg)
}

func (l *logger) Info(msg string) {
	fmt.Println("INFO: ", msg)
}

func (l *logger) Warn(msg string) {
	fmt.Println("WARN: ", msg)
}

func (l *logger) Error(msg string) {
	fmt.Println("ERROR: ", msg)
}
