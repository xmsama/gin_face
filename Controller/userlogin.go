package Controller

import (
	"face/Global"
	"face/Models"
	"face/Utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"net/http"
	"time"
)

type LoginResp struct {
	Code int `json:"code"`
	Data struct {
		User struct {
			ID          int    `json:"ID"`
			UserName    string `json:"userName"`
			NickName    string `json:"nickName"`
			SideMode    string `json:"sideMode"`
			HeaderImg   string `json:"headerImg"`
			BaseColor   string `json:"baseColor"`
			ActiveColor string `json:"activeColor"`
			AuthorityID string `json:"authorityId"`
			Authority   struct {
				DefaultRouter string `json:"defaultRouter"`
			} `json:"authority"`
		} `json:"user"`
		Token     string `json:"token"`
		ExpiresAt string `json:"expiresAt"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func Login(c *gin.Context) {
	data, _ := c.GetRawData()
	var datamap map[string]string

	err := Utils.UnmarshalJSON(c, data, &datamap)
	if err != nil {
		return
	}

	// 设置Payload
	claims := jwt.MapClaims{
		"account": "example_account",
		"time":    time.Now().Unix(),
	}

	// 设置过期时间
	expirationTime := time.Now().Add(5 * time.Minute)
	claims["exp"] = expirationTime.Unix()

	// 创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 设置签名密钥并签名
	signingKey := []byte("your-secret-key")
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Println("JWT签名失败：", err)
		return
	}

	// 输出JWT字符串
	fmt.Println(tokenString)

	fmt.Println(datamap)
	var LoginResp LoginResp
	LoginResp.Data.User.ID = 1
	LoginResp.Data.User.UserName = "管理员"
	LoginResp.Data.User.NickName = "管理员"
	LoginResp.Data.User.SideMode = "light"
	LoginResp.Data.User.HeaderImg = "@/assets/userimage.png"
	LoginResp.Data.User.BaseColor = "#fff"
	//LoginResp.Data.User.
	LoginResp.Data.User.ActiveColor = "#1890ff"
	LoginResp.Data.Token = tokenString
	LoginResp.Data.User.Authority.DefaultRouter = "dashboard"
	jsonStr, err := json.Marshal(LoginResp)
	c.String(http.StatusOK, string(jsonStr))
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 0,
	//	"msg":  "芜湖~",
	//})

}

// AddUser 添加管理员
func AddUser(c *gin.Context) {

}

func Captcha(c *gin.Context) {
	db := Global.DB

	var Count int64
	db.Model(Global.CaptchaModel).Count(&Count)

	config := base64Captcha.DriverString{
		Height:     30,
		Width:      100,
		NoiseCount: 0,
		Length:     5,
		Source:     "1234567890abcdefghijklmnopqrstuvwxyz",
		BgColor:    &color.RGBA{R: 255, G: 255, B: 255, A: 255},

		Fonts: nil,
	}
	captcha := base64Captcha.NewCaptcha(&config, base64Captcha.DefaultMemStore)
	id, b64, _ := captcha.Generate()
	key := captcha.Store.Get(id, true)
	NewCp := Models.Captcha{Base64: b64, Result: key}
	db.Create(&NewCp)
	c.JSON(http.StatusOK, gin.H{
		"code":          0,
		"captchaId":     NewCp.Id,
		"picPath":       b64,
		"captchaLength": 5,
	})

}
