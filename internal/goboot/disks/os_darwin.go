package disks

import (
	"fmt"
	"io"

	"github.com/groob/plist"
	"github.com/sanderploegsma/goboot/internal/goboot/command"
)

func newOS(runner command.Runner) OS {
	return &darwin{runner}
}

type darwin struct {
	runner command.Runner
}

func (os *darwin) GetDisks() ([]DiskInfo, error) {
	info, err := os.getDiskInfo()
	if err != nil {
		return nil, err
	}

	var results []DiskInfo
	for i := range info.AllDisksAndPartitions {
		disk := info.AllDisksAndPartitions[i]
		results = append(results, DiskInfo{
			Name: disk.DeviceIdentifier,
			Size: fmt.Sprintf("%d", disk.Size),
		})
	}

	return results, nil
}

func (os *darwin) Write(disk DiskInfo, data io.Reader) error {
	return nil
}

type appleDiskInfo struct {
	AllDisksAndPartitions []diskAndPartition
}

type diskAndPartition struct {
	DeviceIdentifier string
	Size             uint64
}

// getDiskInfo queries all disks on the host OS using `diskutil` and returns the result
func (os *darwin) getDiskInfo() (*appleDiskInfo, error) {
	output, err := os.runner.Run("diskutil", []string{"list", "-plist"})
	if err != nil {
		return nil, err
	}

	var info appleDiskInfo
	err = plist.Unmarshal(output, &info)
	return &info, err
}
