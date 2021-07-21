package checkin

import "time"

func (c *Client) GetTodayLeave() ([]byte, error) {
	req := c.req.Clone()
	req.Get(todayLeaveUrl)
	return checkError(req.EndBytes())
}

func (c *Client) RequestLeave(reason string) ([]byte, error) {
	today := time.Now().Format("2006-01-02")
	payload := []interface{}{
		//Reply: ["当天出入（23时前返校）", "2021-07-22", "2021-07-14", "0", "否", null]
		"当天出入（23时前返校）",
		today,
		today,
		reason,
		"否",
		nil,
	}
	req := c.req.Clone()
	req.Post(requestLeaveUrl).
		Type("json").
		Send(map[string]interface{}{
			"Reply": payload,
		})
	return checkError(req.EndBytes())
}
