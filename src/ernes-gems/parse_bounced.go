package ernes

import (
	"fmt"
	"io"

	"github.com/Sirupsen/logrus"
)

type bounceType string

const (
	bounceUndetermined bounceType = "Undetermined"
	bouncePermanent    bounceType = "Permanent"
	bounceTransient    bounceType = "Transient"
)

type bounceSubType string

const (
	bounceSubUndetermined       bounceSubType = "Undetermined"
	bounceSubGeneral            bounceSubType = "General"
	bounceSubNoEmail            bounceSubType = "NoEmail"
	bounceSubSuppressed         bounceSubType = "Suppressed"
	bounceSubMailboxFull        bounceSubType = "MailboxFull"
	bounceSubMessageTooLarge    bounceSubType = "MessageTooLarge"
	bounceSubContentRejected    bounceSubType = "ContentRejected"
	bounceSubAttachmentRejected bounceSubType = "AttachmentRejected"
)

type bouncedRecipient struct {
	Email          string `json:"emailAddress"`
	Action         string `json:"action,omitempty"`
	Status         string `json:"status,omitempty"`
	DiagnosticCode string `json:"diagnosticCode,omitempty"`
}
type bounce struct {
	BounceType    bounceType         `json:"bounceType"`
	BounceSubType bounceSubType      `json:"bounceSubType"`
	Recipients    []bouncedRecipient `json:"bouncedRecipients"`
	Timestamp     string             `json:"timestamp"`
	FeedbackId    string             `json:"feedbackId"`
	ReportingMTA  string             `json:"reportingMTA,omitempty"`
	Mail          *mail
}

func (this *bounce) Type() string {
	return fmt.Sprintf("%s.%s", this.BounceType, this.BounceSubType)
}

var parseBounce = func(body io.Reader) (b *bounce) {
	t := parseTop(body)
	if t == nil {
		return nil
	}
	if t.NotificationType != notificationBounce {
		logrus.
			WithField("Type", t.NotificationType).
			Error("Expect notication type to be `Bounce`")
	}
	if t.Bounce == nil {
		logrus.
			WithField("Mail", t).
			Error("`Bounce` object is nil, logging the whole top object")
	}
	b = t.Bounce
	b.Mail = t.Mail
	return b
}
