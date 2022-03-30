package main

import "C"
import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
	"unsafe"

	"github.com/gin-gonic/gin"
	hik "github.com/zwwland/hkweb/hikvision"
)

type Service struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

type SearchRecordReq struct {
	StartAt     string  `json:"startAt"`
	EndAt       string  `json:"endAt"`
	Channel     int     `json:"channel"`
	ServiceInfo Service `json:"serviceInfo"`
}

type RecordInfo struct {
	FileName  string `json:"fileName"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Resp struct {
	Status string      `json:"status"`
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func SuccResp() *Resp {
	return &Resp{
		Status: "ok",
		Error:  "",
		Data:   nil,
	}
}

func SuccDataResp(data interface{}) *Resp {
	return &Resp{
		Status: "ok",
		Error:  "",
		Data:   data,
	}
}
func ErrResp(msg string) *Resp {
	return &Resp{
		Status: "failed",
		Error:  msg,
		Data:   nil,
	}
}

// 获取和下发卡例子
func main() {

	// 初始化
	success := hik.NET_DVR_Init()
	// 设置连接超时时间与重连功能
	success = !success || hik.NET_DVR_SetConnectTime(2000, 1)
	success = !success || hik.NET_DVR_SetReconnect(10000, true)
	if !success {
		// 初始化失败
		printError("Initial failed")
		panic("SDK Initial failed")
	}
	defer hik.NET_DVR_Cleanup()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/sendVoice", handleSendVoice)
	r.POST("/searchRecord", handleSearchRecord)
	r.Run(":50101")
}

// 发送语音
func handleSendVoice(c *gin.Context) {
	fh, _, err := c.Request.FormFile("voice")
	if err != nil {
		c.JSON(http.StatusOK, ErrResp(err.Error()))
		return
	}
	ip := c.PostForm("ip")
	if ip == "" {
		c.JSON(http.StatusOK, ErrResp("ip is empty"))
		return
	}
	port := c.PostForm("port")
	if port == "" {
		c.JSON(http.StatusOK, ErrResp("port is empty"))
		return
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		c.JSON(http.StatusOK, ErrResp("port is not number"))
		return
	}
	user := c.PostForm("user")
	if user == "" {
		c.JSON(http.StatusOK, ErrResp("user is empty"))
		return
	}
	pass := c.PostForm("pass")
	if pass == "" {
		c.JSON(http.StatusOK, ErrResp("pass is empty"))
		return
	}
	result := make(chan bool)

	loginVideo(&Service{
		Ip:   ip,
		Port: portInt,
		User: user,
		Pass: pass,
	}, func() hik.FLoginResultCallBack {
		return func(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
			if 1 != dwResult {
				printError("登陆失败")
				c.JSON(http.StatusOK, ErrResp("登陆失败"))
				return
			}
			go sendAudio(lUserID, pUser, fh, func() {
				fmt.Println("退出登陆了")
				hik.NET_DVR_Logout(lUserID)
			})
			c.JSON(http.StatusOK, SuccResp())
			close(result)
		}
	})
	<-result
	return
}

// 搜索记录
func handleSearchRecord(c *gin.Context) {
	var searchRecordReq SearchRecordReq
	err := c.ShouldBindJSON(&searchRecordReq)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("channel %v \n", searchRecordReq.StartAt)
	result := make(chan bool)

	loginVideo(&searchRecordReq.ServiceInfo, func() hik.FLoginResultCallBack {
		return func(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
			defer hik.NET_DVR_Logout(lUserID)
			defer close(result)
			if 1 != dwResult {
				printError("登陆失败")
				c.JSON(http.StatusOK, ErrResp("登陆失败"))
				return
			}
			list, err := listRecord(lUserID, pUser, &searchRecordReq)
			// c.JSON(http.StatusOK, gin.H{"list": list, "error": err})
			if err != nil {
				c.JSON(http.StatusOK, ErrResp(fmt.Sprintf("错误: %s", err.Error())))
				return
			}
			c.JSON(http.StatusOK, SuccDataResp(list))
			fmt.Println("退出登陆了")
			return
		}
	})
	<-result
	return

}

// vlc rtps://admin:lw123456@222.134.214.134:8011/Streaming/Channels/101
var deviceInfo30 = make(chan hik.LPNET_DVR_DEVICEINFO_V30, 1)
var hangOn = make(chan bool)

func loginVideo(service *Service, handle func() hik.FLoginResultCallBack) {
	// 登陆参数
	loginInfo := hik.NET_DVR_USER_LOGIN_INFO{}
	deviceInfo := hik.NET_DVR_DEVICEINFO_V40{}
	copy(loginInfo.SDeviceAddress[:], service.Ip)
	loginInfo.WPort = uint16(service.Port)
	loginInfo.ByUseTransport = 0
	loginInfo.BUseAsynLogin = 1
	copy(loginInfo.SUserName[:], service.User)
	copy(loginInfo.SPassword[:], service.Pass)
	// 异步登陆回调函数
	loginInfo.CbLoginResult = handle()
	// 登陆
	hik.NET_DVR_Login_V40(&loginInfo, &deviceInfo)
}

// 打印错误信息
func printError(desc string) {
	errCode := hik.NET_DVR_GetLastError()
	fmt.Println(desc, ", error code: ", errCode)
	fmt.Println("Error message: " + hik.NET_DVR_GetErrorMsg(errCode))
}

// 登陆异步回调函数
func loginCallback(lUserID int, dwResult uint32, lpDeviceInfo hik.LPNET_DVR_DEVICEINFO_V30, pUser unsafe.Pointer) {
	if 1 != dwResult {
		fmt.Println("异步登陆失败")
		printError("Login failed")
		hangOn <- false
		return
	} else {
		fmt.Printf("异步登陆成功%x \n", pUser)
		fmt.Println("异步登陆成功，userId: ", strconv.Itoa(lUserID))
		fmt.Println("设备序列号: ", string(lpDeviceInfo.SSerialNumber[:]))
		fmt.Printf("设备序列号: %v \n", lpDeviceInfo)
	}

}

// 查询记录
func listRecord(lUserID int, pUser unsafe.Pointer, req *SearchRecordReq) ([]RecordInfo, error) {
	list := make([]RecordInfo, 0, 4000)

	if lUserID == -1 {
		return list, errors.New("登陆失败")
	}

	st, err := time.Parse("2006-01-02 15:04:05", req.StartAt)
	if err != nil {
		return list, err
	}
	et, err := time.Parse("2006-01-02 15:04:05", req.EndAt)
	if err != nil {
		return list, err
	}
	startime := hik.NET_DVR_TIME{
		DwYear:   int32(st.Year()),
		DwMonth:  int32(st.Month()),
		DwDay:    int32(st.Day()),
		DwHour:   int32(st.Hour()),
		DwMinute: int32(st.Minute()),
		DwSecond: int32(st.Second()),
	}

	endtime := hik.NET_DVR_TIME{
		DwYear:   int32(et.Year()),
		DwMonth:  int32(et.Month()),
		DwDay:    int32(et.Day()),
		DwHour:   int32(et.Hour()),
		DwMinute: int32(et.Minute()),
		DwSecond: int32(et.Second()),
	}

	fileCond := hik.NET_DVR_FILECOND_V40{
		LChannel:      32 + req.Channel,
		DwFileType:    0,
		DwIsLocked:    0,
		DwUseCardNo:   0,
		SCardNumber:   [hik.CARDNUM_LEN_OUT]byte{},
		StruStartTime: startime,
		StruStopTime:  endtime,
	}

	// 查询记录

	handle := hik.NET_DVR_FindFile_V40(lUserID, &fileCond)
	defer hik.NET_DVR_FindClose_V30(handle)

	fmt.Printf("handle: %d \n", handle)
	if handle == -1 {
		printError("Start Find file failed")
		return list, errors.New("查询记录功能,账号密码错误")
	}

	// var data = make([]hik.NET_DVR_FINDDATA_V40, 4000)
	// 逐个查询记录
	var v int
	for {
		d := hik.NET_DVR_FINDDATA_V40{}
		v = hik.NET_DVR_FindNextFile_V40(handle, &d)
		fmt.Println(v)

		if hik.NET_DVR_FILE_FIND_RES(v) == hik.NET_DVR_ISFINDING_res {
			continue
		} else if hik.NET_DVR_FILE_FIND_RES(v) == hik.NET_DVR_NOMOREFILE_res {
			break
		} else if hik.NET_DVR_FILE_FIND_RES(v) == hik.NET_DVR_FILE_NOFIND_res {
			return list, errors.New("没有记录")
		} else if hik.NET_DVR_FILE_FIND_RES(v) == hik.NET_DVR_FILE_SUCCESS_res {
			starts := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", d.StruStartTime.DwYear, d.StruStartTime.DwMonth, d.StruStartTime.DwDay, d.StruStartTime.DwHour, d.StruStartTime.DwMinute, d.StruStartTime.DwSecond)
			ends := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", d.StruStopTime.DwYear, d.StruStopTime.DwMonth, d.StruStopTime.DwDay, d.StruStopTime.DwHour, d.StruStopTime.DwMinute, d.StruStopTime.DwSecond)
			list = append(list, RecordInfo{
				FileName:  string(d.SFileName[:]),
				StartTime: starts,
				EndTime:   ends,
			})
			fmt.Printf("%s->%d %s-%s\n", string(d.SFileName[:]), len(d.SFileName), starts, ends)
			continue
		} else {
			return list, errors.New("没有记录")
		}
	}
	return list, nil
}

// 发送声音
func sendAudio(lUserID int, pUser unsafe.Pointer, file multipart.File, cb func()) bool {
	info := hik.NET_DVR_COMPRESSION_AUDIO{}

	hasInfo := hik.NET_DVR_GetCurrentAudioCompress(lUserID, &info)

	if hasInfo {
		fmt.Printf("ByAudioEncType %d\n", info.ByAudioEncType)
		fmt.Printf("ByAudioSamplingRate %d\n", info.ByAudioSamplingRate)
		fmt.Printf("ByAudioBitRate %d\n", info.ByAudioBitRate)
		fmt.Printf("BySupport %d\n", info.BySupport)
		fmt.Printf("auto %v\n", info)
		var f hik.FVoiceDataCallBack
		start := hik.NET_DVR_StartVoiceCom_MR_V30(lUserID, 0, f, pUser)
		if start == -1 {
			printError("StartVoiceCom_MR_V30 failed")
			return false
		}
		var inputBuffer = make([]byte, 1920)
		// file, err := os.Open(filePath)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return false
		// }
		stari := time.Now().Unix()
		var sendfinish = make(chan bool)
		defer file.Close()
		go func() {
			for {
				_, err := file.Read(inputBuffer)

				if err != io.EOF {
					// fmt.Printf("will send %d", len(inputBuffer))
					ok := hik.NET_DVR_VoiceComSendData(start, inputBuffer, 1920)
					if ok {
						fmt.Println("send data success: ", len(inputBuffer))
					}
					time.Sleep(20 * time.Millisecond)
				} else {
					sendfinish <- true
					break
				}
			}
		}()
		fmt.Println("send voice finished")
		<-sendfinish
		fmt.Printf("send audio success duration: %d \n", time.Now().Unix()-stari)
		ok := hik.NET_DVR_StopVoiceCom(start)
		if !ok {
			fmt.Println("stop voice failed")
		}
		cb()
		return true
	}
	return false
}

// func getCurrentAudioInfo() {

// }
