package main

import "main/checkin"
import "github.com/BurntSushi/toml"

type config struct {
	Receiver  string
	StaffID   string
	StaffName string
	Province  string
	City      string
	Country   string
	Token     string
}

var (
	conf   config
	answer = checkin.Answer{
		"ques1":  "健康良好",
		"ques2":  "正常在家",
		"ques3":  nil,
		"ques4":  "否",
		"ques5":  "否",
		"ques6":  "",
		"ques7":  nil,
		"ques77": nil,
		"ques8":  nil,
		"ques88": nil,
		"ques9":  nil,
		"ques10": nil,
		"ques11": nil,
		"ques12": nil,
		"ques13": nil,
		"ques14": nil,
		"ques15": "否",
		"ques16": "否",
		"ques17": "无新冠肺炎确诊或疑似",
		"ques18": "37度以下",
		"ques19": nil,
		"ques20": "绿码",
		"ques21": "否",
		"ques22": "否",
		"ques23": "否",
		"carTo": []string{
			//Province
			"330000",
			////City
			"330100",
			////Country
			"330104",
		},
	}
)

func init() {
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		panic(err)
	}
	//override answer position
	if conf.Province != "" {
		answer["carTo"] = []string{
			conf.Province,
			conf.City,
			conf.Country,
		}
	}
}