package date_utils

import "time"

const (
	mysqlDateFormat =  "2006-01-02 15:04:05"
)

func GetNowDbFormat() string {
	return time.Now().UTC().Format(mysqlDateFormat)
}
