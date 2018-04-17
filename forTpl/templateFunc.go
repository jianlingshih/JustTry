package forTpl

import (
	"strconv"
	"time"
)

//转换价格  分==>元
func GetPrice(price interface{}) string {
	value1, ok := price.(int64)
	var re float64
	if ok {
		re = float64(value1) / 100.0
	} else {
		value2, ok := price.(int)
		if ok {
			re = float64(value2) / 100.0
		}
	}

	priceFormat := strconv.FormatFloat(re, 'f', -1, 32)

	return priceFormat
}

//根据时间戳获取日期字符创
func GetDateHM(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}
