package checkin

import "encoding/json"

type Answer map[string]interface{}

type ReqCheckIn struct {
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Answer    string `json:"answerJsonStr"`
}

type ApiResp struct {
	Error int             `json:"error"`
	Msg   string          `json:"msg"`
	Data  json.RawMessage `json:"data"`
}

type TokenValid struct {
	AccessToken string `json:"accessToken"`
	Valid       int    `json:"isValid"`
}
