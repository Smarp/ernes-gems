package ernes

import (
	"io"

	"github.com/Sirupsen/logrus"
)

type complaintFeedbackType string

const (
	complaintFeedbackAbuse       complaintFeedbackType = "abuse"
	complaintFeedbackAuthFailure complaintFeedbackType = "auth-failure"
	complaintFeedbackFraud       complaintFeedbackType = "fraud"
	complaintFeedbackNotSpam     complaintFeedbackType = "not-spam"
	complaintFeedbackOther       complaintFeedbackType = "other"
	complaintFeedbackVirus       complaintFeedbackType = "virus"
)

type complainedRecipient struct {
	Email string `json:"emailAddress"`
}
type complaint struct {
	Recipients  []complainedRecipient `json:"complainedRecipients"`
	Timestamp   string                `json:"timestamp"`
	FeedbackId  string                `json:"feedbackId"`
	UserAgent   string                `json:"userAgent,omitempty"`
	Type        complaintFeedbackType `json:"complaintFeedbackType,omitempty"`
	ArrivalDate string                `json:"arrivalDate,omitempty"`
	Mail        *mail
}

var parseComplaint = func(body io.Reader) (c *complaint) {
	t := parseTop(body)
	if t == nil {
		return nil
	}
	if t.NotificationType != notificationComplaint {
		logrus.
			WithField("Type", t.NotificationType).
			Error("Expect notication type to be `Complaint`")
	}
	if t.Complaint == nil {
		logrus.
			WithField("Mail", t).
			Error("`Complaint` object is nil, logging the whole top object")
	}
	c = t.Complaint
	c.Mail = t.Mail
	return c
}
