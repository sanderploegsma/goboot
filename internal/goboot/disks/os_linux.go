package disks

import (
	"encoding/json"
	"io"

	"github.com/sanderploegsma/goboot/internal/goboot/command"
)

func newOS(runner command.Runner) OS {
	return &linux{runner}
}

type linux struct {
	runner command.Runner
}

func (os *linux) GetDisks() ([]DiskInfo, error) {
	disks, err := os.getDiskInfo()
	if err != nil {
		return nil, err
	}

	var results []DiskInfo
	for i := range disks.BlockDevices {
		disk := disks.BlockDevices[i]
		if disk.Type == "disk" {
			results = append(results, DiskInfo{
				Name: disk.Name,
				Size: disk.Size,
			})
		}
	}

	return results, nil
}

func (os *linux) Write(disk DiskInfo, data io.Reader) error {
	return nil
}

// getDiskInfo lists all disks on the host OS using 'lsdsk' and parses the output
func (os *linux) getDiskInfo() (*linuxDiskInfo, error) {
	output, err := os.runner.Run("lsdsk", []string{"--json", "--bytes"})
	if err != nil {
		return nil, err
	}

	var info linuxDiskInfo
	err = json.Unmarshal(output, &info)
	return &info, err
}

type linuxDiskInfo struct {
	BlockDevices []blockDevice
}

type blockDevice struct {
	Name string
	Size string
	Type string
}
