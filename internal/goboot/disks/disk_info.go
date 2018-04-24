package disks

import (
	"io"

	"github.com/sanderploegsma/goboot/internal/goboot/command"
)

// OS represents an operating system capable of retrieving and working with disks
type OS interface {
	GetDisks() ([]DiskInfo, error)
	Write(disk DiskInfo, data io.Reader) error
}

// DiskInfo represents a disk on the host OS, such as:
// - `/dev/sdb1` on Linux
// - `/dev/disk1s1` on MacOS
// - `D:` on Windows
type DiskInfo struct {
	Name string
	Size string
}

// NewOS creates a new OS implementation based on the architecture
func NewOS(runner command.Runner) OS {
	return newOS(runner)
}
