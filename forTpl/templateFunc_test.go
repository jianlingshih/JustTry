package forTpl

import (
	"testing"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func Test_DealWithRegexp(t *testing.T) {

	beego.Debug("Hello World!", GetWeekday(1523967273))
	beego.Debug("+++++++", SnowFlakeId())
}

// func Test_GetPeriodByLesson(t *testing.T) {
// 	cc, _, _ := GetPeriodByLesson(1523967273)
// 	beego.Debug("Hello World!", cc)
// }

// func Test_GetDateHM(t *testing.T) {
// 	cc := GetDateHM(1523967273)
// 	beego.Debug("Hello World!", cc)
// }

// func Test_GetPrice(t *testing.T) {
// 	cc := GetPrice(100)
// 	beego.Debug("Hello World!", cc)
// }
