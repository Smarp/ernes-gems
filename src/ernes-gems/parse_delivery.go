package ernes

import (
	"io"

	"github.com/Sirupsen/logrus"
)

type delivery struct {
	Timestamp            string   `json:"timestamp"`
	ProcessingTimeMillis int      `json:"processingTimeMillis"`
	Recipients           []string `json:"recipients"`
	SmtpResponse         string   `json:"smtpResponse"`
	ReportingMTA         string   `json:"reportingMTA"`
	Mail                 *mail
}

var parseDelivery = func(body io.Reader) (d *delivery) {
	t := parseTop(body)
	if t == nil {
		return nil
	}
	if t.NotificationType != notificationDelivery {
		logrus.
			WithField("Type", t.NotificationType).
			Error("Expect notication type to be `Delivery`")
		return nil
	}
	if t.Delivery == nil {
		logrus.
			WithField("Mail", t).
			Error("`Delivery` object is nil, logging the whole top object")
		return nil
	}
	d = t.Delivery
	d.Mail = t.Mail
	return d
}
