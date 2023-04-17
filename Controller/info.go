package Controller

import (
	"face/Global"
	"face/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"time"

	"net/http"
	"os"
	"strconv"
)

type LSysInfo struct {
	MemAll         uint64
	MemFree        uint64
	MemUsed        uint64
	MemUsedPercent float64
	Days           int64
	Hours          int64
	Minutes        int64
	Seconds        int64

	CpuUsedPercent float64
	OS             string
	Arch           string
	CpuCores       int
}

func GetInfo() (string, string, string) {
	c, _ := cpu.Info()
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Get Error"
	}
	physicalCnt, _ := cpu.Counts(false)
	logicalCnt, _ := cpu.Counts(true)
	var modelname string
	if len(c) > 1 {
		for _, sub_cpu := range c {
			modelname = sub_cpu.ModelName
		}
	} else {
		sub_cpu := c[0]
		modelname = sub_cpu.ModelName
		//cores := sub_cpu.Cores
		//fmt.Printf("CPU: %v   %v cores \n", modelname, cores)
	}

	unit := uint64(1024 * 1024) // MB
	v, _ := mem.VirtualMemory()
	var info LSysInfo
	info.MemAll = v.Total
	info.MemAll /= unit
	var Memory float64
	Memory = float64(info.MemAll)
	Memory, _ = strconv.ParseFloat(fmt.Sprintf("%.1f", Memory/1024), 64)
	Ram := fmt.Sprintf("%.1f", Memory)
	modelname += " (" + strconv.Itoa(physicalCnt) + "核心 " + strconv.Itoa(logicalCnt) + "線程)"
	return hostname, modelname, Ram + "G"
}

func GetIndexInfo(c *gin.Context) {
	type L struct {
		ServerName     string
		GoVer          string
		Ram            string
		CPU            string
		OsVersion      string
		TodaySuccess   int
		TodayFail      int
		LastDaySuccess int
		LastDayFail    int
	}
	type Response struct {
		Code int         `json:"code"`
		Data interface{} `json:"data"`
		Msg  string      `json:"msg"`
	}

	Hostname, Cpu, Memory := GetInfo()
	db := Global.DB
	var TodayInfo Models.Info
	var LastDayInfo Models.Info
	yesterday := time.Now().AddDate(0, 0, -1)
	db.Where("date = ?", time.Now().Format("2006-01-02")).Take(&TodayInfo)
	db.Where("date = ?", yesterday.Format("2006-01-02")).Take(&LastDayInfo)
	list := []L{{ServerName: Hostname, GoVer: "go1.19.1", OsVersion: "Ubuntu-22/33", CPU: Cpu, Ram: Memory, LastDayFail: LastDayInfo.Fail,
		LastDaySuccess: LastDayInfo.Success, TodaySuccess: TodayInfo.Success, TodayFail: TodayInfo.Fail}}
	//List1.List = list

	c.JSON(http.StatusOK, Response{
		0,
		list,
		"",
	})
}
