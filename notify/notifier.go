package notify

import (
	"fmt"
)

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dm"
)

var (
	notifier *dm.Client
)

func Init(key, secret string) {
	client, err := dm.NewClientWithAccessKey(
		"cn-hangzhou", key, secret)
	if err != nil {
		panic(err)
	}
	notifier = client
}

func New(title, body string) *Notifier {
	return &Notifier{
		Title: title,
		Body:  body,
	}
}

func (n *Notifier) Notify(to string) error {
	req := newMailRequest(to, n.Title, n.Body)
	resp, err := notifier.SingleSendMail(req)
	if err != nil {
		return fmt.Errorf("send: %v", err)
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("response: %v", resp.String())
	}
	return nil
}

func newMailRequest(to, title, body string) *dm.SingleSendMailRequest {
	req := dm.CreateSingleSendMailRequest()
	req.Scheme = "https"
	req.AccountName = "noreply@dreamer2q.wang"
	req.AddressType = requests.NewInteger(1)
	req.ReplyToAddress = requests.NewBoolean(false)
	req.ToAddress = to
	req.Subject = title
	req.TextBody = body
	return req
}
