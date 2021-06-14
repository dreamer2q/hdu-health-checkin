package main

import (
	"encoding/json"
	"fmt"
	"log"
	"main/checkin"
	"main/notify"
	"strings"
)

var (
	client *checkin.Client
)

func main() {
	client = checkin.New(checkin.Config{
		StaffID:     conf.StaffID,
		StaffName:   conf.StaffName,
		Province:    conf.Province,
		City:        conf.City,
		Country:     conf.Country,
		AccessToken: conf.Token,
	})

	checkJob()
}

func checkJob() {
	var sb = &strings.Builder{}
	_, _ = fmt.Fprintf(sb, "执行结果\n\n")
	if token, err := client.ValidToken(); err != nil {
		log.Printf("valid token: %s", err)
		_, _ = fmt.Fprintf(sb, "检查Token失败: %s\n", err)
		goto ret
	} else {
		_, _ = fmt.Fprintf(sb, "1. 检查Token: %s\n", token)
	}
	if info, err := client.GetDaily(); err != nil {
		log.Printf("get daily info: %s", err)
		_, _ = fmt.Fprintf(sb, "查询信息失败: %s\n", err)
		goto ret
	} else {
		_, _ = fmt.Fprintf(sb, "2. 查询打卡信息: %s\n", info)
		var infoResp = struct {
			Msg string `json:"message"`
		}{}
		if err := json.Unmarshal(info, &infoResp); err != nil {
			log.Printf("unmarshal: %v", err)
			goto ret
		}
		if strings.Contains(infoResp.Msg, "已填报") {
			_, _ = fmt.Fprintf(sb, "3. 已填报，跳过打卡")
			goto ret
		}
	}
	if res, err := client.CheckIn(answer); err != nil {
		log.Printf("check in: %s", err)
		_, _ = fmt.Fprintf(sb, "打卡失败: %s\n", err)
		goto ret
	} else {
		_, _ = fmt.Fprintf(sb, "3. 打卡结果: %s\n", res)
	}

ret:
	var mail = notify.New("健康打卡小助手", sb.String())
	if err := mail.Notify(conf.Receiver); err != nil {
		log.Panic(err)
	}
	log.Printf("success")
}
