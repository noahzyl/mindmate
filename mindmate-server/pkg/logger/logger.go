/*
 * Define logic of the logger
 */

package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siruspen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"
)

// Initialize the logger
func init() {
	// Set the format of logs as JSON
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", // Set format of timestamp
	})
	// Do not record caller of logs
	logrus.SetReportCaller(false)
}

// Write a log (info level)
func Write(message string, fileName string) {
	LogToFile(logrus.InfoLevel, fileName)
	logrus.Info(message)
}

func Debug(fields logrus.Fields, args ...interface{}) { // ... means variadic parameters
	LogToFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args) // WithFields() can add custom fields the log
}

func Info(fields logrus.Fields, args ...interface{}) {
	LogToFile(logrus.InfoLevel, "info")
	logrus.WithFields(fields).Info(args)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	LogToFile(logrus.WarnLevel, "warn")
	logrus.WithFields(fields).Warn(args)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	LogToFile(logrus.FatalLevel, "fatal")
	logrus.WithFields(fields).Fatal(args)
}

func Error(fields logrus.Fields, args ...interface{}) {
	LogToFile(logrus.ErrorLevel, "error")
	logrus.WithFields(fields).Error(args)
}

func Panic(fields logrus.Fields, args ...interface{}) {
	LogToFile(logrus.PanicLevel, "panic")
	logrus.WithFields(fields).Panic(args)
}

func Trace(fields logrus.Fields, args ...interface{}) {
	LogToFile(logrus.TraceLevel, "trace")
	logrus.WithFields(fields).Trace(args)
}

// Create a log file
func setLogFile(logName string) string {
	// Check whether the directory of log files exists
	if _, err := os.Stat("./runtime/log"); os.IsNotExist(err) {
		// If the directory not exists, then create it
		err = os.MkdirAll("./runtime/log", 0777)
		if err != nil { // Error
			panic(fmt.Errorf("create log dir ./runtime/log error: %s", err))
		}
	}

	// Set log file name
	timeStr := time.Now().Format("2006-01-02") // Date
	// path.Join() will automatically handle different operating system path separators
	fileName := path.Join("./runtime/log", logName+"_"+timeStr+".log")

	return fileName
}

// Write a log
func LogToFile(level logrus.Level, logName string) {
	// Open log file and write the log
	fileName := setLogFile(logName)
	var err error
	// Use 0644 to prevent other people to change the log
	/* The opened file pointer is not directly assigned to a new variable,
	 * but instead, it is assigned to os.Stderr, overwriting the standard error stream with the file pointer.
	 * This means that the logs will not only be written to the standard error stream (os.Stderr),
	 * but the log content will also be written to the specified log file.
	 */
	os.Stderr, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file error:", err)
	}
	logrus.SetOutput(os.Stderr) // Set output target
	logrus.SetLevel(level)
}

// Write a request log
func LogRequest() gin.LoggerConfig {
	fileName := setLogFile("success")
	// Open log file
	var err error // Variable shadowing
	os.Stderr, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file error:", err)
	}

	// Set log config of gin
	conf := gin.LoggerConfig{
		// Set format of the log (by defining a function)
		Formatter: func(param gin.LogFormatterParams) string { // gin.LogFormatterParams contains details of a request
			return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
				param.TimeStamp.Format("2006-01-02 15:04:05"),
				param.ClientIP,
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
			)
		},
		// Set output targets
		Output: io.MultiWriter(os.Stdout, os.Stderr),
	}

	return conf
}

// Write an error log
func LogError(ctx *gin.Context) {
	defer func() {
		// If catchy a panic, write an error log
		if err := recover(); err != nil {
			fileName := setLogFile("error")
			// Open log file
			file, errFile := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if errFile != nil {
				fmt.Println("open log file error:", errFile)
			}
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			file.WriteString("panic error time: " + timeStr + "\n")
			file.WriteString(fmt.Sprintf("%v\n", err)) // Error info
			file.WriteString("stacktrace from panic: " + string(debug.Stack()) + "\n")
			file.Close()
			ctx.JSON(http.StatusOK, gin.H{
				"code":    500, // Internal error of the server
				"message": fmt.Sprintf("%v", err),
			})
			ctx.Abort() // Stop handling requests
		}
	}()
	ctx.Next() // If no error, continue to handle requests
}
