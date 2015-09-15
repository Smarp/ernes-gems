package ernes

import (
	"bytes"
	"testing"

	. "github.com/onsi/gomega"
)

func TestTopMail(t *testing.T) {
	RegisterTestingT(t)
	const jsonInput = `{
      "notificationType":"Bounce",
      "bounce":{
         "bounceType":"Permanent",
         "bounceSubType": "General",
         "bouncedRecipients":[
            {
               "emailAddress":"recipient1@example.com"
            },
            {
               "emailAddress":"recipient2@example.com"
            }
         ],
         "timestamp":"2012-05-25T14:59:38.237-07:00",
         "feedbackId":"00000137860315fd-869464a4-8680-4114-98d3-716fe35851f9-000000"
      },
      "mail":{
         "timestamp":"2012-05-25T14:59:38.237-07:00",
         "messageId":"00000137860315fd-34208509-5b74-41f3-95c5-22c1edc3c924-000000",
         "source":"email_1337983178237@amazon.com",
         "sourceArn": "arn:aws:ses:us-west-2:888888888888:identity/example.com",
         "sendingAccountId":"123456789012",
         "destination":[
            "recipient1@example.com",
            "recipient2@example.com",
            "recipient3@example.com",
            "recipient4@example.com"
         ]
      }
   }`
	body := bytes.NewBufferString(jsonInput)
	b := parseBounce(body)

	Expect(b).NotTo(BeNil())

	Expect(b.Mail).NotTo(BeNil())

	Expect(b.Mail.Timestamp).To(Equal("2012-05-25T14:59:38.237-07:00"))
	Expect(b.Mail.MessageId).To(Equal("00000137860315fd-34208509-5b74-41f3-95c5-22c1edc3c924-000000"))
	Expect(b.Mail.Source).To(Equal("email_1337983178237@amazon.com"))
	Expect(b.Mail.SourceArn).To(Equal("arn:aws:ses:us-west-2:888888888888:identity/example.com"))
	Expect(b.Mail.SendingAccountId).To(Equal("123456789012"))

	Expect(b.Mail.Destination).To(HaveLen(4))
	Expect(b.Mail.Destination[0]).To(Equal("recipient1@example.com"))
	Expect(b.Mail.Destination[3]).To(Equal("recipient4@example.com"))
}
