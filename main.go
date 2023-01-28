package main

import (
	"github.com/sirupsen/logrus"
)

var dbFile = "local/db.yaml"
var logger = logrus.New()

func main() {
	router()
}
