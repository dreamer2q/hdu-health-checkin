package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/checkin"
	"main/notify"
	"os"
	"strings"
)

var (
	client *checkin.Client
	sb     = &strings.Builder{}
)

func main() {
	notify.Init(conf.AccessKey, conf.AccessSecret)
	client = checkin.New(checkin.Config{
		StaffID:     conf.StaffID,
		StaffName:   conf.StaffName,
		Province:    conf.Province,
		City:        conf.City,
		Country:     conf.Country,
		AccessToken: conf.Token,
	})

	_ = start()
	_ = notify.New("健康打卡成功", sb.String()).
		Notify(conf.Receiver)
	log.Printf("success")
}

func writeLine(format string, v ...interface{}) {
	log.Printf(format+"\n", v...)
	_, _ = fmt.Fprintf(sb, format+"\n", v...)
}

func writeError(format string, v ...interface{}) error {
	log.Printf(format, v...)
	_, _ = fmt.Fprintf(sb, format, v...)

	var mail = notify.New("健康打卡失败", sb.String())
	if err := mail.Notify(conf.Receiver); err != nil {
		log.Panic(err)
	}
	os.Exit(-1)
	return nil
}

func ensure(err error) {
	if err != nil {
		_ = writeError("错误：%v", err)
	}
}

func start() error {
	writeLine("执行结果")
	if token, err := client.ValidToken(); err != nil {
		return writeError("健康打卡失败： %v", err)
	} else {
		writeLine("1. 检查 Token: %s", token)
	}
	info, err := client.GetDaily()
	if err != nil {
		return writeError("查询信息失败：%v\n", err)
	}
	writeLine("2. 查询打卡信息：%s\n", info)
	var infoResp = struct {
		Msg string `json:"message"`
	}{}
	ensure(json.Unmarshal(info, &infoResp))
	if strings.Contains(infoResp.Msg, "已填报") {
		return writeError("3. 已填报，跳过打卡")
	}
	if result, err := client.CheckIn(answer); err != nil {
		return writeError("打卡失败: %s\n", err)
	} else {
		writeLine("3. 打卡结果: %s\n", result)
	}
	info, err = client.GetTodayLeave()
	if err != nil {
		return writeError("获取通行信息失败：%v", err)
	}
	writeLine("4. 获取通信信息：%s", info)
	var leave = struct {
		detail []json.RawMessage
	}{}
	ensure(json.Unmarshal(info, &leave))
	if len(leave.detail) > 0 {
		return writeError("已有出入通行许可，跳过")
	}
	info, err = client.RequestLeave("0")
	if err != nil {
		return writeError("申请出入校失败：%v", err)
	}
	writeLine("5. 获取出入校通行许可：%s", info)
	return nil
}
