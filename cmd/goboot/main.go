package main

import (
	"github.com/sanderploegsma/goboot/internal/goboot/command"
	"github.com/sanderploegsma/goboot/internal/goboot/disks"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	runner := command.NewRunner()
	os := disks.NewOS(runner)

	disks, err := os.GetDisks()
	if err != nil {
		panic(err)
	}

	logrus.Debug(disks)
}
