package forTpl

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/zheng-ji/goSnowFlake"
)

// SnowFlakeId 生成雪花ID
func SnowFlakeId() int64 {
	rwork := int64(rand.Intn(100))
	iw, _ := goSnowFlake.NewIdWorker(rwork)
	if id, err := iw.NextId(); err != nil {
		return 0
	} else {
		return id
	}
}

//MD5处理
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//base 64 加解密
func base64Encode(src string) string {
	return string([]byte(base64.StdEncoding.EncodeToString([]byte(src)))[:])
}
func base64Decode(src string) string {
	re, _ := base64.StdEncoding.DecodeString(src)
	return string(re[:])
}

//对map进行排序
func MapSort() {
	// To create a map as input
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "c"
	m[0] = "b"

	// To store the keys in slice in sorted order
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// To perform the opertion you want
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

//slice 去重
func RemoveRep(s []int) (result []int) {
	s = []int{1, 2, 3, 4, 5, 6, 2, 3}
	start := time.Now()
	m := make(map[int]bool)
	for _, v := range s {
		if _, ok := m[v]; !ok {
			result = append(result, v)
			m[v] = true
		}
	}
	fmt.Println("花费时间:", fmt.Sprintf("%vms", (time.Now().UnixNano()-start.UnixNano())/1e+6))
	beego.Debug(result)
	return
}

//正则处理
func DealWithRegexp() {
	reg, _ := regexp.Compile("^(\\d+)(\\w+)")
	if reg.MatchString("20min") == true {
		submatch := reg.FindStringSubmatch("20min")
		beego.Debug(submatch[2])
	}
}

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
func GetMidnightTimestamp(day int) (re int64) {
	t := time.Now()
	re = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).AddDate(0, 0, day).Unix()
	return
}

//根据时间戳获取日期字符创
func GetDateHM(timestamp int64) string {
	if timestamp <= 0 {
		return ""
	}
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04")
}

//根据时间戳获取周几
func GetWeekday(timestamp int64) string {
	tm2 := time.Unix(timestamp, 0)
	week := tm2.Weekday().String()
	weekCN := GetWeekCN(week)
	return weekCN

}
func GetWeekCN(weekEn string) string {
	var txt string

	switch weekEn {

	case "Sunday":
		txt = "周日"
	case "Monday":
		txt = "周一"
	case "Tuesday":
		txt = "周二"
	case "Wednesday":
		txt = "周三"
	case "Thursday":
		txt = "周四"
	case "Friday":
		txt = "周五"
	case "Saturday":
		txt = "周六"
	}

	return txt
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
//{"A":{"name":"上海虹桥国际机场  ","latlng":"31.19668,121.3376"},"B":{"name":"上海长途汽车客运总站  ","latlng":"31.25217,121.45472"}}
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

/** 获取活动详情
// [{"name":"Safety Check","type":"T1","content":["A10010","A10030","A10035"],"time":10}]
// var sessioncontent []SessionContent
// 	err = json.Unmarshal([]byte(content), &sessioncontent)

func GetActivityDetail(sessioncontent []SessionContent, lang string) ([]SessionContentDetail, AthleteNow) {
	var athleteNow AthleteNow
	detail := make([]SessionContentDetail, len(sessioncontent))
	for k, v := range sessioncontent {
		if v.Type == "T4" {
			athleteNow.HasAthleteNow = 1
			athleteNow.AthleteNowId = v.Name
			y := v.Content.([]interface{})
			athleteNow.AthleteNowActivity = make([]string, len(y))
			for i := range y {
				athleteNow.AthleteNowActivity[i] = y[i].(string)
			}
		}
		count := int64(0)
		detail[k].Name = v.Name
		detail[k].Type = v.Type
		detail[k].Time = v.Time
		switch v.Content.(type) {
		case string:
			var sessionActivity []ActivityDetail
			activity, _ := GetActivity(v.Content.(string))

			sessionActivity = append(sessionActivity, GetActivityAfterSplited(activity, lang))
			detail[k].Content = sessionActivity
		case []interface{}:
			a := v.Content.([]interface{})
			sessionActivity := make([]ActivityDetail, len(a))
			var sign string
			for i := range a {
				activity, _ := GetActivity(a[i].(string))
				sessionActivity[i] = GetActivityAfterSplited(activity, lang)
				if sign == sessionActivity[i].Skill_lang {
					sessionActivity[i].Display = "none"
					count++
				}
				sign = sessionActivity[i].Skill_lang
			}
			detail[k].Content = sessionActivity
			detail[k].Rowspan = 2*int64(len(a)) - count
		}

	}
	return detail, athleteNow

}
*/
