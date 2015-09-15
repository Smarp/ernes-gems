package ernes

import (
	"bytes"
	"testing"

	. "github.com/onsi/gomega"
)

func TestBounceWithoutDSN(t *testing.T) {
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

	Expect(b.Type()).To(Equal("Permanent.General"))

	Expect(b.Recipients).To(HaveLen(2))
	Expect(b.Recipients[0].Email).To(Equal("recipient1@example.com"))
	Expect(b.Recipients[1].Email).To(Equal("recipient2@example.com"))

	Expect(b.Timestamp).To(Equal("2012-05-25T14:59:38.237-07:00"))
	Expect(b.FeedbackId).To(Equal("00000137860315fd-869464a4-8680-4114-98d3-716fe35851f9-000000"))

	Expect(b.Mail).NotTo(BeNil())
}

func TestBounceWithDSN(t *testing.T) {
	RegisterTestingT(t)
	const jsonInput = `{
       "notificationType":"Bounce",
       "bounce":{
          "bounceType":"Permanent",
          "reportingMTA":"dns; email.example.com",
          "bouncedRecipients":[
             {
                "emailAddress":"username@example.com",
                "status":"5.1.1",
                "action":"failed",
                "diagnosticCode":"smtp; 550 5.1.1 <username@example.com>... User"
             }
          ],
          "bounceSubType":"General",
          "timestamp":"2012-06-19T01:07:52.000Z",
          "feedbackId":"00000138111222aa-33322211-cccc-cccc-cccc-ddddaaaa068a-000000"
       },
       "mail":{
          "timestamp":"2012-06-19T01:05:45.000Z",
          "source":"sender@example.com",
          "sourceArn": "arn:aws:ses:us-west-2:888888888888:identity/example.com",
          "sendingAccountId":"123456789012",
          "messageId":"00000138111222aa-33322211-cccc-cccc-cccc-ddddaaaa0680-000000",
          "destination":[
             "username@example.com"
          ]
       }
    }`
	body := bytes.NewBufferString(jsonInput)
	b := parseBounce(body)

	Expect(b).NotTo(BeNil())

	Expect(b.ReportingMTA).To(Equal("dns; email.example.com"))

	Expect(b.Recipients).To(HaveLen(1))
	Expect(b.Recipients[0].Status).To(Equal("5.1.1"))
	Expect(b.Recipients[0].Action).To(Equal("failed"))
	Expect(b.Recipients[0].DiagnosticCode).To(Equal("smtp; 550 5.1.1 <username@example.com>... User"))

	Expect(b.Mail).NotTo(BeNil())
}
