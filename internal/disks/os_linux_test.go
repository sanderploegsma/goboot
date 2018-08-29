package disks

import (
	"errors"
	"testing"
)

func TestGetDisks(t *testing.T) {
	t.Run("lsblk returns error", func(t *testing.T) {
		runner := &testRunner{
			shouldError: true,
			err:         errors.New("lsblk failed"),
		}

		os := &linux{runner}

		_, err := os.GetDisks()
		if err == nil {
			t.Fatal("expected error but got none")
		}
	})

	t.Run("lsblk returns valid output", func(t *testing.T) {
		runner := &testRunner{
			shouldError: false,
			output:      []byte(lsblkOutput),
		}

		os := &linux{runner}
		disks, err := os.GetDisks()
		if err != nil {
			t.Fatalf("expected no error but got '%s'", err.Error())
		}

		if got, want := len(disks), 1; got != want {
			t.Fatalf("expected number of disks to be %d but got %d", want, got)
		}

		if got, want := disks[0].Name, "sda"; got != want {
			t.Fatalf("expected first disk name to be %s but got %s", want, got)
		}

		if got, want := disks[0].Size, "63999836160"; got != want {
			t.Fatalf("expected first disk size to be %s but got %s", want, got)
		}
	})
}

type testRunner struct {
	shouldError bool
	err         error
	output      []byte
}

func (t *testRunner) Run(command string, args []string) ([]byte, error) {
	if t.shouldError {
		return nil, t.err
	}
	return t.output, nil
}

var lsblkOutput = `{
	"blockdevices": [
	   {"name": "sda", "maj:min": "8:0", "rm": "0", "size": "63999836160", "ro": "0", "type": "disk", "mountpoint": null,
		  "children": [
			 {"name": "sda1", "maj:min": "8:1", "rm": "0", "size": "63998787584", "ro": "0", "type": "part", "mountpoint": "/etc/hosts"}
		  ]
	   },
	   {"name": "sr0", "maj:min": "11:0", "rm": "1", "size": "984238080", "ro": "0", "type": "rom", "mountpoint": null},
	   {"name": "sr1", "maj:min": "11:1", "rm": "1", "size": "49152", "ro": "0", "type": "rom", "mountpoint": null}
	]
 }`
