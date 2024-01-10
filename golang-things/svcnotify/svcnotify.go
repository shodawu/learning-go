package svcnotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
	"github.com/tealeg/xlsx/v3"
	"gopkg.in/gomail.v2"
)

// FTPConfig ...
type FTPConfig struct {
	Host    string
	Port    int
	Account string
	Pwd     string
	DocPath string
}

// EmailConfig ...
type EmailConfig struct {
	Host    string
	Port    int
	Account string
	Pwd     string
}

// Config FTP Email 設定
type Config struct {
	FTP   FTPConfig
	Email EmailConfig
}

// MailMessage ...
type MailMessage struct {
	ReceiverAddr string
	MessageBody  string
}

// SvcNotify ...
type SvcNotify struct {
	Config   Config
	FileBuf  []byte
	Messages []MailMessage
}

// LoadConfig 讀取設定
func (svc *SvcNotify) LoadConfig(filePath string) {
	var file, _ = os.OpenFile(filePath, os.O_RDONLY, 0755)
	b, _ := ioutil.ReadAll(file)

	json.Unmarshal(b, &svc.Config)

}

// FTPController ...
func (svc *SvcNotify) FTPController() {
	c, err := ftp.Dial(fmt.Sprintf("%v:%v", svc.Config.FTP.Host, svc.Config.FTP.Port), ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login(svc.Config.FTP.Account, svc.Config.FTP.Pwd)
	if err != nil {
		log.Fatal(err)
	}

	r, err := c.Retr(svc.Config.FTP.DocPath)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	svc.FileBuf, err = ioutil.ReadAll(r)

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}

// ExcelController ...
func (svc *SvcNotify) ExcelController() {
	wb, err := xlsx.OpenBinary(svc.FileBuf)
	if err != nil {
		panic(err)
	}
	sh, ok := wb.Sheet["總表"]
	if !ok {
		panic(errors.New("Sheet not found"))
	}

	err = sh.ForEachRow(svc.rowVisitor)

	if err != nil {
		log.Fatal(err)
	}
}

func (svc *SvcNotify) rowVisitor(r *xlsx.Row) error {

	taskID := strings.TrimSpace(r.GetCell(0).Value)
	taskContent := strings.TrimSpace(r.GetCell(1).Value)
	chargers := strings.Split(strings.TrimSpace(r.GetCell(2).Value), ",")
	isClosed := false
	if strings.TrimSpace(r.GetCell(3).Value) == "V" {
		isClosed = true
	}
	if taskID != "" && taskID[0] == '#' && !isClosed {
		sendText := fmt.Sprintf(`
		<table>
		<tr><td>事項ID</td><td>事項內容</td></tr>
		<tr><td>%v</td><td>%v</td></tr>
		</table>
		`, taskID, taskContent)

		for _, charger := range chargers {
			charger := strings.TrimSpace(charger)
			svc.Messages = append(svc.Messages, MailMessage{
				ReceiverAddr: charger,
				MessageBody:  sendText,
			})
		}
	}

	return nil
}

// MailController ...
func (svc *SvcNotify) MailController() {
	for _, msg := range svc.Messages {
		if err := svc.sendMail("通知", msg.MessageBody, msg.ReceiverAddr); err != nil {
			log.Fatal(err)
		}
	}
}

func (svc *SvcNotify) sendMail(subject, body, receiver string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", svc.Config.Email.Account)
	rec := strings.Split(receiver, "@")
	m.SetAddressHeader("To", receiver, rec[0])

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(svc.Config.Email.Host, svc.Config.Email.Port, svc.Config.Email.Account, svc.Config.Email.Pwd)

	err := d.DialAndSend(m)

	return err
}

// TriggerExec ...
func TriggerExec() {
	svc := SvcNotify{}
	svc.LoadConfig("config.json")
	svc.FTPController()
	svc.ExcelController()
	svc.MailController()
}
