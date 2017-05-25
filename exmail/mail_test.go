package exmail

import (
	"testing"
)

func TestSend(t *testing.T) {
	host := "smtp.exmail.qq.com"
	port := 465
	email := "monitor2@maxpanda.cn"
	password := "Monitor123"
	toEmail := "wanghu@maxpanda.cn"
	body := ""
	body += "SLB报警：<br />"
	body += "IP：127.0，0.1<br />"
	body += "类型: 下线<br />"

	header := make(map[string]string)
	header["From"] = "monitor " + "<" + email + ">"
	header["To"] = toEmail
	header["Subject"] = "游戏SLB监控"
	header["Content-Type"] = "text/html; charset=UTF-8"

	err := sendMail(host, port, email, password, body, header)
	if err != nil {
		t.Errorf("TestSend error:%v", err.Error())
	}
}

func TestSendMulti(t *testing.T) {
	host := "smtp.exmail.qq.com"
	port := 465
	email := "monitor2@maxpanda.cn"
	password := "Monitor123"
	toAddress := "wanghu@maxpanda.cn;357658079@qq.com"
	from := "MonitorTest"
	subject := "MonitorSubject"
	body := ""
	body += "SLB报警：<br />"
	body += "IP：127.0，0.1<br />"
	body += "类型: 下线<br />"
	sendErrors := Send(host, port, email, password, toAddress, from, subject, body)
	if len(sendErrors) > 0 {
		t.Errorf("TestSendMulti error:%v", sendErrors[0].Error())
	}
}
