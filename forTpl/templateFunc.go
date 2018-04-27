package forTpl

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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
func GetDateAdd() {
	//Add方法和Sub方法是相反的，获取t0和t1的时间距离d是使用Sub，将t0加d获取t1就是使用Add方法
	k := time.Now()
	//一天之前
	d, _ := time.ParseDuration("-24h")
	fmt.Println(k.Add(d))
	//一周之前
	fmt.Println(time.Unix(k.Add(d*7).Unix(), 0).Format("20060102"))
	//一月之前
	fmt.Println(k.Add(d * 30))

}

//根据时间戳获取日期字符创
func GetDateHM(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}

// GetPeriodByLesson 根据lesson的开课时间获取period 如 201803B
func GetPeriodByLesson(timestamp int64) (period string, start int64, end int64) {
	DATE_FORMAT := "2006-01-02"
	year := time.Unix(timestamp, 0).Year()
	month := time.Unix(timestamp, 0).Month()
	day := time.Unix(timestamp, 0).Day()
	period = time.Unix(timestamp, 0).Format("200601")

	theMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	loc, _ := time.LoadLocation("Local")
	starttime, _ := time.ParseInLocation(DATE_FORMAT, theMonth.AddDate(0, 0, 0).Format(DATE_FORMAT), loc)
	endtime, _ := time.ParseInLocation(DATE_FORMAT, theMonth.AddDate(0, 1, 0).Format(DATE_FORMAT), loc)
	middletime, _ := time.ParseInLocation(DATE_FORMAT, theMonth.AddDate(0, 0, 15).Format(DATE_FORMAT), loc)
	if day > 15 {
		start = middletime.Unix()
		end = endtime.Unix()
		period += "B"
	} else {
		start = starttime.Unix()
		end = middletime.Unix()
		period += "A"
	}
	return
}

//初始化ID号
func GenerateSKUID(skuid string, maxid int64) string {
	maxid++
	maxlen := 5 - len(fmt.Sprintf("%d", maxid))
	for i := 0; i < maxlen; i++ {
		skuid += "0"
	}
	skuid += fmt.Sprintf("%d", maxid)
	return skuid
}

//获取集合点A&B
type PointInfo struct {
	Name   string `json:"name"`
	Latlng string `json:"latlng"`
}

func GetRellyStr(points string) (pointA, pointB string) {
	var data map[string]PointInfo
	if err := json.Unmarshal([]byte(points), &data); err != nil {
		return
	} else {
		pointA = data["A"].Name + " 坐标:" + data["A"].Latlng
		if len(data) > 1 {
			pointB = data["B"].Name + " 坐标:" + data["B"].Latlng
		}
	}
	return
}

func GetRellyPoints(pointA, pointB string) (re string) {

	points := make(map[string]PointInfo)
	var point PointInfo
	a := strings.Split(pointA, "坐标:")
	b := strings.Split(pointB, "坐标:")
	if len(a) == 1 {
		return
	}
	point.Name = a[0]
	point.Latlng = a[1]

	points["A"] = point
	if len(b) > 1 {
		point.Name = b[0]
		point.Latlng = b[1]
		points["B"] = point
	}

	j, _ := json.Marshal(points)
	re = string(j)
	return
}
