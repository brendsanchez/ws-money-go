package config

import (
	"github.com/sirupsen/logrus"
	"time"
)

type customFormatter struct {
	logrus.Formatter
	Location *time.Location
}

// Format to accept timeZone
func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Time = entry.Time.In(f.Location)
	return f.Formatter.Format(entry)
}

func InitLogrus(cfg *Config) {
	if cfg.Logger.IsDebug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	logrus.SetReportCaller(cfg.Logger.ReportCaller)
	logrus.SetFormatter(jsonFormatter())
}

func jsonFormatter() logrus.Formatter {
	return &customFormatter{
		Formatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "logTime",
				logrus.FieldKeyFile: "logCaller",
				logrus.FieldKeyFunc: "logFunc",
				logrus.FieldKeyMsg:  "message",
			},
			TimestampFormat: "2006-01-02 15:04:05",
			PrettyPrint:     false,
		},
		Location: timeZone(),
	}
}

func timeZone() *time.Location {
	loc, err := time.LoadLocation("America/Buenos_Aires")
	if err != nil {
		logrus.Fatal(err)
	}
	return loc
}
