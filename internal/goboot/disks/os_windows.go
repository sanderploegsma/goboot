package disks

import (
	"io"

	"github.com/sanderploegsma/goboot/internal/goboot/command"
)

func newOS(runner command.Runner) OS {
	return &windows{runner}
}

type windows struct {
	runner command.Runner
}

func (os *windows) GetDisks() ([]DiskInfo, error) {
	return nil, nil
}

func (os *windows) Write(disk DiskInfo, data io.Reader) error {
	return nil
}
