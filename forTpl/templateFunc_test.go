package forTpl

import (
	"testing"

	"github.com/astaxie/beego"

	_ "github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func Test_GetPrice(t *testing.T) {
	cc := GetPrice(100)
	beego.Debug("Hello World!", cc)
}
