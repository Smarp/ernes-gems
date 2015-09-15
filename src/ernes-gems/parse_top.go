package ernes

import (
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
)

type notificationType string

const (
	notificationBounce    notificationType = "Bounce"
	notificationComplaint notificationType = "Complaint"
	notificationDelivery  notificationType = "Delivery"
)

type mail struct {
	Timestamp        string   `json:"timestamp"`
	MessageId        string   `json:"messageId"`
	Source           string   `json:"source"`
	SourceArn        string   `json:"sourceArn"`
	SendingAccountId string   `json:"sendingAccountId"`
	Destination      []string `json:"destination"`
}

type sns struct {
	Message string `json:"message"`
}

type top struct {
	NotificationType notificationType `json:"notificationType"`
	Mail             *mail            `json:"mail"`
	Bounce           *bounce          `json:"bounce"`
	Complaint        *complaint       `json:"complaint"`
	Delivery         *delivery        `json:"delivery"`
}

var parseTop = func(body io.Reader) (t *top) {
	byteBody := parsePrepare(body)
	if byteBody == nil {
		return nil
	}
	s := new(sns)
	err := json.Unmarshal(byteBody, s)
	if err != nil {
		logrus.
			WithFields(logrus.Fields{
			"Err":  err,
			"Body": string(byteBody),
		}).
			Error("Cannot parse SNS Body")
		return nil
	}
	t = new(top)
	err = json.Unmarshal([]byte(s.Message), t)
	if err != nil {
		logrus.
			WithFields(logrus.Fields{
			"Err":  err,
			"Body": s.Message,
		}).
			Error("Cannot parse top Body")
		return nil
	}
	return t
}

var parsePrepare = func(body io.Reader) []byte {
	byteBody, err := ioutil.ReadAll(body)
	if err != nil {
		logrus.
			WithField("Err", err).
			Error("Cannot read body")
		return nil
	}
	return byteBody
}
