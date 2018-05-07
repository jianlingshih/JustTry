package handleImg

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/disintegration/imaging"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	qrcode "github.com/skip2/go-qrcode"
	"golang.org/x/image/font"
)

// GetHeadUrl 根据unionid 获取客户详情
func GetHeadUrl(unionid string) (re string, err error) {
	o := orm.NewOrm()
	rs := o.Raw("SELECT headimgurl FROM cms_customers WHERE unionid=?", unionid)
	err = rs.QueryRow(&re)
	return
}

//根据url获取图片
func GetImgByUrl(imagPath, fileName string) {
	//通过http请求获取图片的流文件
	resp, _ := http.Get(imagPath)
	body, _ := ioutil.ReadAll(resp.Body)
	out, _ := os.Create(fileName)
	io.Copy(out, bytes.NewReader(body))
	return
}

//生成二维码  参数1: url链接   参数2:生成二维码图片的大小  参数3：生成二维码的路径
func GenerateQrCode(url string, size int, filename string) (err error) {
	err = qrcode.WriteFile(url, qrcode.Medium, size, filename)
	if err != nil {
		beego.Debug("Generate qrcode failed")
	}
	return
}
func getFontFamily(ttcStr string) (*truetype.Font, error) {
	// 这里需要读取中文字体，否则中文文字会变成方格
	fontBytes, err := ioutil.ReadFile(ttcStr)
	if err != nil {
		fmt.Println("read file error:", err)
		return &truetype.Font{}, err
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("parse font error:", err)
		return &truetype.Font{}, err
	}

	return f, err
}

// 给图片贴图
func AddOnePic(img, pic, fileName string, width, height, pointX, pointY int) (string, error) {
	m, err := imaging.Open(pic) //插入的贴图
	if err != nil {
		fmt.Printf("open1 file failed")
	}

	bm, err := imaging.Open(img) //主图片
	if err != nil {
		fmt.Printf("open2 file failed")
	}
	// 图片按比例缩放
	dst := imaging.Resize(m, width, height, imaging.Lanczos)
	// 将图片粘贴到背景图的固定位置
	result := imaging.Overlay(bm, dst, image.Pt(pointX, pointY), 1)
	err = imaging.Save(result, fileName)
	if err != nil {
		return "", err
	}

	return fileName, nil
}

//给图片添加文字
func writeOnImg(fileName, ttcStr, content string, fontSize float64, pointX, pointY int) error {
	m, err := imaging.Open(fileName) //获取需要添加文字的图片
	if err != nil {
		fmt.Printf("open1 file failed")
	}
	var dpi = flag.Float64("dpi", 256, "screen resolution")

	target := imaging.Clone(m)
	c := freetype.NewContext()

	c.SetDPI(*dpi)
	c.SetClip(target.Bounds())
	c.SetDst(target)
	c.SetHinting(font.HintingFull)

	// 设置文字颜色、字体、字大小
	c.SetSrc(image.NewUniform(color.RGBA{R: 240, G: 240, B: 245, A: 180}))
	c.SetFontSize(fontSize)

	fontFam, err := getFontFamily(ttcStr)
	if err != nil {
		fmt.Println("get font family error")
	}
	c.SetFont(fontFam)

	pt := freetype.Pt(pointX, pointY)

	_, err = c.DrawString(content, pt)
	if err != nil {
		fmt.Printf("draw error: %v \n", err)
	}
	err = imaging.Save(target, fileName)
	if err != nil {
	}
	return err
}
