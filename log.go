package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

const (
	logFilePath, logFileName = "/var/log/webserver", "webserver.log"
)

var (
	Logger  *logrus.Logger
	CstZone = time.FixedZone("CST", 8*3600)
)

type StandardFormatter struct{}

// Format 格式化日志条目为标准格式的字节切片。
// 它使用当前时间戳、日志级别和日志消息来构建日志条目。
// 参数:
//
//	entry - 要格式化的日志条目。
//
// 返回值:
//
//	格式化后的日志条目的字节切片和可能的错误。
func (s *StandardFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().In(CstZone).Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

// LoggerToFile 是一个 Gin 中间件，用于将日志记录到文件。
// 它会初始化日志记录器，并设置日志文件的分割、保存时间和切割时间间隔。
// 日志记录器会记录每个请求的开始时间、结束时间、延迟时间、请求方法、请求 URI、状态码和客户端 IP。
// 返回的 gin.HandlerFunc 会在每个请求处理完成后记录这些信息。
func LoggerToFile() gin.HandlerFunc {
	fileName := path.Join(logFilePath, logFileName)

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println("open log file err:", err)
	}

	if Logger == nil {
		Logger = logrus.New()
	}

	Logger.Out = src

	Logger.SetLevel(logrus.InfoLevel)

	writer, err := rotateLogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d",

		// 生成软链，指向最新日志文件
		rotateLogs.WithLinkName(fileName),

		// 最大保存时间
		rotateLogs.WithMaxAge(time.Hour*time.Duration(24*7)),

		// 设置日志切割时间间隔(1天)
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		log.Fatalln("rotateLogs.New err:", err)
	}

	Logger.SetOutput(writer)
	Logger.SetFormatter(new(StandardFormatter))
	Logger.Info("logs init success")

	return func(c *gin.Context) {
		startTime := time.Now().In(CstZone)
		c.Next()
		endTime := time.Now().In(CstZone)
		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		Logger.Infof("| %3d | %13v | %10s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri)
	}
}
