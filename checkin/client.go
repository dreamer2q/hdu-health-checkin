package checkin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

type Config struct {
	StaffID     string
	StaffName   string
	Province    string
	City        string
	Country     string
	AccessToken string
	Debug       bool
}

type Client struct {
	Config
	req *gorequest.SuperAgent
}

func New(c Config) *Client {
	req := gorequest.New()
	req.Set("User-Agent", agent)
	req.Set("X-Requested-With", "com.tencent.mm")
	req.Set("Referer", "https://healthcheckin.hduhelp.com/")
	req.Set("Origin", "https://healthcheckin.hduhelp.com")
	req.Set("Authorization", c.AccessToken)
	return &Client{
		Config: c,
		req:    req,
	}
}

func (c *Client) ValidToken() ([]byte, error) {
	req := c.req.Clone()
	req.Get(validUrl)
	var body, err = checkError(req.EndBytes())
	if err != nil {
		return nil, err
	}
	var token = TokenValid{}
	err = json.Unmarshal(body, &token)
	if err != nil || token.Valid == 0 {
		return nil, fmt.Errorf("err: %v, token: %v", err, token)
	}
	return body, err
}

func (c *Client) GetInfo() ([]byte, error) {
	req := c.req.Clone()
	req.Get(infoUrl)
	return checkError(req.EndBytes())
}

func (c *Client) GetDaily() ([]byte, error) {
	req := c.req.Clone()
	req.Get(dailyUrl)
	return checkError(req.EndBytes())
}

func (c *Client) CheckIn(ans Answer) ([]byte, error) {
	req := c.req.Clone()
	now := time.Now().Unix()
	sign := c.getSign(now)
	req.Post(fmt.Sprintf(checkInUrl, sign))
	req.Type("json")
	ansJson, err := json.Marshal(ans)
	if err != nil {
		return nil, fmt.Errorf("marshal answer: %v", err)
	}
	req.Send(ReqCheckIn{
		Name:      c.StaffName,
		Timestamp: now,
		Province:  c.Province,
		City:      c.City,
		Country:   c.Country,
		Answer:    string(ansJson),
	})
	return checkError(req.EndBytes())
}

func (c *Client) getSign(timestamp int64) string {
	sign := fmt.Sprintf("%s%s%d%s%s%s", c.StaffName, b64(c.StaffID), timestamp, b64(c.Province), c.City, c.Country)
	return s1(sign)
}

func checkError(_ gorequest.Response, body []byte, errs []error) ([]byte, error) {
	if errs != nil {
		return nil, fmt.Errorf("error: %v", errs)
	}
	api := ApiResp{}
	err := json.Unmarshal(body, &api)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %v", err)
	}
	if api.Error != 0 {
		return nil, fmt.Errorf("api error: %d %s", api.Error, api.Msg)
	}
	data := bytes.NewBuffer(nil)
	if err := json.Indent(data, api.Data, "", "  "); err != nil {
		return nil, fmt.Errorf("json indent: %v", err)
	}
	return data.Bytes(), nil
}
