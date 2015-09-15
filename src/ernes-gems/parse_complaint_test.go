package ernes

import (
	"bytes"
	"testing"

	. "github.com/onsi/gomega"
)

func TestComplaintWithoutFeedback(t *testing.T) {
	RegisterTestingT(t)
	const jsonInput = `{
      "notificationType":"Complaint",
      "complaint":{
         "complainedRecipients":[
            {
               "emailAddress":"recipient1@example.com"
            }
         ],
         "timestamp":"2012-05-25T14:59:38.613-07:00",
         "feedbackId":"0000013786031775-fea503bc-7497-49e1-881b-a0379bb037d3-000000"
      },
      "mail":{
         "timestamp":"2012-05-25T14:59:38.613-07:00",
         "messageId":"0000013786031775-163e3910-53eb-4c8e-a04a-f29debf88a84-000000",
         "source":"email_1337983178613@amazon.com",
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
	c := parseComplaint(body)

	Expect(c).NotTo(BeNil())

	Expect(c.Recipients).To(HaveLen(1))
	Expect(c.Recipients[0].Email).To(Equal("recipient1@example.com"))

	Expect(c.Timestamp).To(Equal("2012-05-25T14:59:38.613-07:00"))
	Expect(c.FeedbackId).To(Equal("0000013786031775-fea503bc-7497-49e1-881b-a0379bb037d3-000000"))

}

func TestComplaintWithFeedback(t *testing.T) {
	RegisterTestingT(t)
	const jsonInput = `{
      "notificationType":"Complaint",
      "complaint":{
         "userAgent":"Comcast Feedback Loop (V0.01)",
         "complainedRecipients":[
            {
               "emailAddress":"recipient1@example.com"
            }
         ],
         "complaintFeedbackType":"abuse",
         "arrivalDate":"2009-12-03T04:24:21.000-05:00",
         "timestamp":"2012-05-25T14:59:38.623-07:00",
         "feedbackId":"000001378603177f-18c07c78-fa81-4a58-9dd1-fedc3cb8f49a-000000"
      },
      "mail":{
         "timestamp":"2012-05-25T14:59:38.623-07:00",
         "messageId":"000001378603177f-7a5433e7-8edb-42ae-af10-f0181f34d6ee-000000",
         "source":"email_1337983178623@amazon.com",
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
	c := parseComplaint(body)

	Expect(c).NotTo(BeNil())

	Expect(c.UserAgent).To(Equal("Comcast Feedback Loop (V0.01)"))
	Expect(c.Type).To(Equal(complaintFeedbackAbuse))
	Expect(c.ArrivalDate).To(Equal("2009-12-03T04:24:21.000-05:00"))

}
