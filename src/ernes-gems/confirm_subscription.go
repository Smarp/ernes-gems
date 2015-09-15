package ernes

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
)

type subscriptionConfirmation struct {
	SubscribeURL string `json:"subscribeURL"`
}

var confirmSubscription = func(r *http.Request) bool {
	if r.Header.Get("x-amz-sns-message-type") != "SubscriptionConfirmation" {
		return false
	}
	defer r.Body.Close()
	d := json.NewDecoder(r.Body)
	c := new(subscriptionConfirmation)
	err := d.Decode(c)
	if err != nil {
		logrus.
			WithField("Err", err).
			Error("Cannot decode subscriptionConfirmation")
		return false
	}
	logrus.WithField("Url", c.SubscribeURL).Info("SubscribeURL")
	return true
}
