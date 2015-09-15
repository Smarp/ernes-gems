package ernes

import (
	"bytes"
	"testing"

	. "github.com/onsi/gomega"
)

func TestDelivery(t *testing.T) {
	RegisterTestingT(t)
	const jsonInput = `{
      "notificationType":"Delivery",
      "mail":{
         "timestamp":"2014-05-28T22:40:59.638Z",
         "messageId":"0000014644fe5ef6-9a483358-9170-4cb4-a269-f5dcdf415321-000000",
         "source":"sender@example.com",
         "sourceArn": "arn:aws:ses:us-west-2:888888888888:identity/example.com",
         "sendingAccountId":"123456789012",
         "destination":[
            "success@simulator.amazonses.com",
            "recipient@example.com"
         ]
      },
      "delivery":{
         "timestamp":"2014-05-28T22:41:01.184Z",
         "recipients":["success@simulator.amazonses.com"],
         "processingTimeMillis":546,
         "reportingMTA":"a8-70.smtp-out.amazonses.com",
         "smtpResponse":"250 ok:  Message 64111812 accepted"
      }
   }`
	body := bytes.NewBufferString(jsonInput)
	d := parseDelivery(body)

	Expect(d).NotTo(BeNil())

	Expect(d.Timestamp).To(Equal("2014-05-28T22:41:01.184Z"))

	Expect(d.Recipients).To(HaveLen(1))
	Expect(d.Recipients[0]).To(Equal("success@simulator.amazonses.com"))

	Expect(d.ProcessingTimeMillis).To(Equal(546))
	Expect(d.ReportingMTA).To(Equal("a8-70.smtp-out.amazonses.com"))
	Expect(d.SmtpResponse).To(Equal("250 ok:  Message 64111812 accepted"))

	Expect(d.Mail).NotTo(BeNil())

}
