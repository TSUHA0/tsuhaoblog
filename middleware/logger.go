package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	path := "log/log"
	logger := logrus.New()

	src, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("failed open log file, %v\n", err)
	}

	//logger.Out = src 此处写入log/log文件 是总记录文件
	logger.Out = src
	// Only log the warning severity or above.
	logger.SetLevel(logrus.DebugLevel)

	logWriter, err := rotatelogs.New(
		path+"%Y%m%d.log", //分割日志文件
		//MaxAge and RotationCount cannot be both set  两者不能同时设置
		rotatelogs.WithMaxAge(7*24*time.Hour),     //clear 最小分钟为单位
		rotatelogs.WithRotationTime(24*time.Hour), //rotate 最小为1分钟轮询。默认60s  低于1分钟就按1分钟来
	)
	if err != nil {
		fmt.Printf("failed to create rotatelogs: %v\n", err)
	}

	writeMap := lfshook.WriterMap{
		logrus.TraceLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendT := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000)))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}

		dataSize := c.Writer.Size()
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		method := c.Request.Method
		uri := c.Request.RequestURI
		usrAgent := c.Request.UserAgent()

		entry := logger.WithFields(logrus.Fields{
			"SpendTime": spendT,
			"HostName":  hostName,
			"DataSize":  dataSize,
			"Status":    statusCode,
			"ClientIp":  clientIp,
			"Method":    method,
			"URI":       uri,
			"usrAgent":  usrAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
