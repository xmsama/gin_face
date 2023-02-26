package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndexInfo(c *gin.Context) {
	type L struct {
		AccountTotal    int64
		MachineTotal    int64
		CardTotal       int64
		NoCardMachine   int64
		NoOutCard       int64
		RunMachine      int64
		ErrorAccount    int64
		CompleteAccount int64
		Man             int64
		Woman           int64
		ServerName      string
		HTTPVer         string
		WEBName         string
		GoVer           string
		Ram             string
		CPU             string
		OsVersion       string
		Disk            string
		Heart           string
	}
	type Response struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}
	list := []L{{AccountTotal: 1, MachineTotal: 1, CardTotal: 1, NoCardMachine: 1, NoOutCard: 1,
		RunMachine: 1, ErrorAccount: 1, CompleteAccount: 1, Man: 1, Woman: 1, Ram: "12/22G", Heart: "OK",
		ServerName: "Ubuntu-ss22", GoVer: "go1.19.1", HTTPVer: "HTTP/1.1", WEBName: "No", OsVersion: "Ubuntu-22/33", CPU: "I7 22322344544", Disk: "已用123123/12312"}}
	//List1.List = list

	c.JSON(http.StatusOK, Response{
		0,
		list,
		"",
	})
}
