package disks

import (
	"io"

	"github.com/sanderploegsma/goboot/internal/goboot/command"
)

func newOS(runner command.Runner) OS {
	return &linux{runner}
}

type linux struct {
	runner command.Runner
}

func (os *linux) GetDisks() ([]*DiskInfo, error) {
	return nil, nil
}

func (os *linux) Write(disk *DiskInfo, data io.Reader) error {
	return nil
}
