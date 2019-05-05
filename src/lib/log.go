package lib

import l "github.com/sirupsen/logrus"

func log() {
	l.Info("Simple log line")
	l.Warn("Another log line")
	l.Error("This is bad")
	l.Fatal("This is really bad, exiting...")
}