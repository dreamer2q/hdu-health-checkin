# HDU Health CheckIn & Request For A Leave

> This Repo only contains a go package && a demo, which means you have to code for yourself
>
>If you cannot code, there is nothing for you here

A go package providing auto checkin ability for `hduers`.

## *Note*

This package has timeliness, and cannot guarantee its functionality.

## Setting Up

1. GET `AccessToken`

open [this link](https://api.hduhelp.com/login/direct/yiban?clientID=healthcheckin&redirect=https%3A%2F%2Fhealthcheckin.hduhelp.com%2F%23%2Fauth)
to get your access token

*Or, you could get it by any other means.*

2. Config YOUR **PERSONAL** && **Answer** Information

check [province.json5](province.json5) to get the location code

```go
client := checkin.New(checkin.Config{
StaffID:     "your-staff-id",
StaffName:   "your-real-name",
//YOUR **LOCATION**
Province:    "000000",
City:        "000000",
Country:     "000000",
AccessToken: "token xxxx-xxxxxx-xxxxxx-xxxxxxx-xx",
})
```

3. Fill YOUR **ANSWER** form

```go
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
"000000",
//City
"000000",
//Country
"000000",
},
}
```

4. See [main.go](main.go) for more details

# LICENSE

MIT
