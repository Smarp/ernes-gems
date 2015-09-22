package ernes

import (
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
)

func Main() {
	http.HandleFunc("/bounce", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if confirmSubscription(r) {
			return
		}
		bounce := parseBounce(r.Body)
		if bounce == nil || bounce.Mail == nil {
			logrus.
				WithField("Headers", r.Header).
				Error("Bounced object is nil")
			return
		}
		logrus.WithFields(logrus.Fields{
			"Type":         bounce.Type(),
			"Recipients":   bounce.Recipients,
			"FeedbackId":   bounce.FeedbackId,
			"MessageId":    bounce.Mail.MessageId,
			"ReportingMTA": bounce.ReportingMTA,
			"Timestamp":    bounce.Timestamp,
		}).Error("Bounced email")
	})

	http.HandleFunc("/complaint", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if confirmSubscription(r) {
			return
		}
		complaint := parseComplaint(r.Body)
		if complaint == nil || complaint.Mail == nil {
			logrus.
				WithField("Headers", r.Header).
				Error("Complaint object is nil")
			return
		}
		logrus.WithFields(logrus.Fields{
			"UserAgent":   complaint.UserAgent,
			"Recipients":  complaint.Recipients,
			"Type":        complaint.Type,
			"FeedbackId":  complaint.FeedbackId,
			"ArrivalDate": complaint.ArrivalDate,
			"MessageId":   complaint.Mail.MessageId,
			"Timestamp":   complaint.Timestamp,
		}).Error("Complaint email")
	})

	http.HandleFunc("/delivery", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if confirmSubscription(r) {
			return
		}
		delivery := parseDelivery(r.Body)
		if delivery == nil || delivery.Mail == nil {
			logrus.
				WithField("Headers", r.Header).
				Error("Delivery object is nil")
			return
		}
		logrus.WithFields(logrus.Fields{
			"Timestamp":    delivery.Timestamp,
			"Recipients":   delivery.Recipients,
			"ReportingMTA": delivery.ReportingMTA,
			"SmtpResponse": delivery.SmtpResponse,
			"MessageId":    delivery.Mail.MessageId,
		}).Info("Delivery email")
	})

	if len(os.Args) < 2 {
		logrus.
			WithField("len(os.Args)", len(os.Args)).
			Fatal("No port set, expect len(os.Args) 2")
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		logrus.
			WithField("err", err).
			Fatal("No port set")
	}
	logrus.Info("Starting listenting")
	// but still waiting for dbWaitGroup
	ln, err := net.ListenTCP("tcp", &net.TCPAddr{
		Port: port,
	})
	if err != nil {
		logrus.
			WithFields(logrus.Fields{
			"Err":  err,
			"Port": port,
		}).
			Fatal("Cannot listen to TCP Port")
	}
	logrus.Info("Starting serving")
	// Start the server!

	srv := &http.Server{
		// be careful of FD leaks
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logrus.Fatal(srv.Serve(ln))
	// code can never reach here!!!
}
