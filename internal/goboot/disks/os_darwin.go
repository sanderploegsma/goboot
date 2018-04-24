package disks

import (
	"io"
	"io/ioutil"
	"regexp"

	"github.com/sanderploegsma/goboot/internal/goboot/command"
	"github.com/sirupsen/logrus"
)

func newOS(runner command.Runner) OS {
	return &darwin{runner}
}

type darwin struct {
	runner command.Runner
}

func (os *darwin) GetDisks() ([]*DiskInfo, error) {
	diskNames, err := os.getDiskNames()
	if err != nil {
		return nil, err
	}
	for _, name := range diskNames {
		output, err := os.runner.Run("diskutil", []string{"info", name})
		if err != nil {
			return nil, err
		}

		logrus.Debug(output)
	}
	return nil, nil
}

func (os *darwin) Write(disk *DiskInfo, data io.Reader) error {
	return nil
}

// getDiskNames looks for all disks matching /dev/disk*
func (os *darwin) getDiskNames() ([]string, error) {
	regex := regexp.MustCompile("^/dev/disk([0-9]+)$")
	var devDisks []string
	files, err := ioutil.ReadDir("/dev/")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		fileName := "/dev/" + f.Name()
		if regex.MatchString(fileName) {
			logrus.WithField("disk", fileName).Debug("Found disk")
			devDisks = append(devDisks, fileName)
		}
	}
	return devDisks, nil
}
