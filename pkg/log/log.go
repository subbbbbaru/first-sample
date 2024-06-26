// package log

// import (
// 	"log"
// 	"os"
// )

// var (
// 	InfoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
// 	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
// )

package log

import (
	"io"
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

// InitLoggers инициализирует логгеры с заданными выходными потоками
func InitLoggers(infoHandle io.Writer, errorHandle io.Writer) {
	infoLogger = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info возвращает логгер для информационных сообщений
func Info() *log.Logger {
	return infoLogger
}

// Error возвращает логгер для сообщений об ошибках
func Error() *log.Logger {
	return errorLogger
}

// Инициализация логгеров по умолчанию при старте
func init() {
	InitLoggers(os.Stdout, os.Stderr)
}
