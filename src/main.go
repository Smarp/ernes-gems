package main

import (
	ernes "ernes-gems"

	"github.com/Sirupsen/logrus"
)

func main() {
	logrus.Info("Starting main process")
	ernes.Main()
}
