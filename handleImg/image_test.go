package handleImg

import (
	"fmt"
	"testing"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//根据unionid 获取头像链接
func Test_GetHeadUrl(t *testing.T) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:111111@tcp(127.0.0.1:3306)/nncms?charset=utf8", 30)
	orm.Debug = true
	unionid := "oR0D703X3FWPxtkwVYS_djaBxdTk"
	headurl, _ := GetHeadUrl(unionid)
	beego.Debug("headUrl", headurl)

}

//获取头像图片
func Test_GetImgByUrl(t *testing.T) {
	url := "http://thirdwx.qlogo.cn/mmopen/vi_32/Rd8PyOa5ichtKtGR4JcodU0GQxUPAUyXP70kTjSBXTf1Phj8lOrEJJTOPNABLqEicZCDk3qmxjeuiasCqt2Vol96w/132"
	fileName := "../static/data/qrcode/tmp/cc.jpg"
	GetImgByUrl(url, fileName)
}

//生成二维码
func Test_GenerateQrCode(t *testing.T) {
	err := GenerateQrCode("https://www.sogou.com", 256, "test001.png")
	if err != nil {
		fmt.Println("生成二维码 失败，详情", err.Error())
		return
	}
	fmt.Println("生成二维码成功")
}

func Test_AddOnePic(t *testing.T) {
	img := "test.jpeg"         //主图片
	pic := "test001.png"       //要插入的贴图
	fileName := "zxcvb002.jpg" //生成图片的路径
	width := 80                //缩放参数
	height := 80               //缩放参数
	pointX := 290              //固定的坐标参数
	pointY := 290              //固定的坐标参数
	_, err := AddOnePic(img, pic, fileName, width, height, pointX, pointY)
	if err != nil {
		fmt.Println("合成图片失败")
	} else {
		fmt.Println("合成图片成功")
	}

	ttc := "test.ttc"        //字体存放位置
	content := "测试数据007"     //要添加的文字内容
	fontSize := float64(8.8) //要添加的文字的字号
	err = writeOnImg(fileName, ttc, content, fontSize, 360, 330)
	if err != nil {
		fmt.Println("添加文字失败")
	} else {
		fmt.Println("添加文字成功")
	}
}
