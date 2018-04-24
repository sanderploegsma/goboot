package disks

import (
	"errors"
	"testing"
)

func TestGetDisks(t *testing.T) {
	t.Run("diskutil returns error", func(t *testing.T) {
		runner := &testRunner{
			shouldError: true,
			err:         errors.New("diskutil failed"),
		}

		os := &darwin{runner}

		_, err := os.GetDisks()
		if err == nil {
			t.Fatal("expected error but got none")
		}
	})

	t.Run("diskutil returns valid output", func(t *testing.T) {
		runner := &testRunner{
			shouldError: false,
			output:      []byte(diskUtilOutput),
		}

		os := &darwin{runner}
		disks, err := os.GetDisks()
		if err != nil {
			t.Fatalf("expected no error but got '%s'", err.Error())
		}

		if got, want := len(disks), 2; got != want {
			t.Fatalf("expected number of disks to be %d but got %d", want, got)
		}

		if got, want := disks[0].Name, "disk0"; got != want {
			t.Fatalf("expected first disk name to be %s but got %s", want, got)
		}

		if got, want := disks[0].Size, "121332826112"; got != want {
			t.Fatalf("expected first disk size to be %s but got %s", want, got)
		}

		if got, want := disks[1].Name, "disk1"; got != want {
			t.Fatalf("expected second disk name to be %s but got %s", want, got)
		}

		if got, want := disks[1].Size, "121123069952"; got != want {
			t.Fatalf("expected second disk size to be %s but got %s", want, got)
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

var diskUtilOutput = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>AllDisks</key>
	<array>
		<string>disk0</string>
		<string>disk0s1</string>
		<string>disk0s2</string>
		<string>disk1</string>
		<string>disk1s1</string>
		<string>disk1s2</string>
		<string>disk1s3</string>
		<string>disk1s4</string>
	</array>
	<key>AllDisksAndPartitions</key>
	<array>
		<dict>
			<key>Content</key>
			<string>GUID_partition_scheme</string>
			<key>DeviceIdentifier</key>
			<string>disk0</string>
			<key>Partitions</key>
			<array>
				<dict>
					<key>Content</key>
					<string>EFI</string>
					<key>DeviceIdentifier</key>
					<string>disk0s1</string>
					<key>DiskUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
					<key>Size</key>
					<integer>209715200</integer>
					<key>VolumeName</key>
					<string>EFI</string>
					<key>VolumeUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
				</dict>
				<dict>
					<key>Content</key>
					<string>Apple_APFS</string>
					<key>DeviceIdentifier</key>
					<string>disk0s2</string>
					<key>DiskUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
					<key>Size</key>
					<integer>121123069952</integer>
				</dict>
			</array>
			<key>Size</key>
			<integer>121332826112</integer>
		</dict>
		<dict>
			<key>APFSVolumes</key>
			<array>
				<dict>
					<key>DeviceIdentifier</key>
					<string>disk1s1</string>
					<key>DiskUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
					<key>MountPoint</key>
					<string>/</string>
					<key>Size</key>
					<integer>121123069952</integer>
					<key>VolumeName</key>
					<string>Macintosh HD</string>
					<key>VolumeUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
				</dict>
				<dict>
					<key>DeviceIdentifier</key>
					<string>disk1s2</string>
					<key>DiskUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
					<key>Size</key>
					<integer>121123069952</integer>
					<key>VolumeName</key>
					<string>Preboot</string>
					<key>VolumeUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
				</dict>
				<dict>
					<key>DeviceIdentifier</key>
					<string>disk1s3</string>
					<key>DiskUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
					<key>Size</key>
					<integer>121123069952</integer>
					<key>VolumeName</key>
					<string>Recovery</string>
					<key>VolumeUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
				</dict>
				<dict>
					<key>DeviceIdentifier</key>
					<string>disk1s4</string>
					<key>DiskUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
					<key>MountPoint</key>
					<string>/private/var/vm</string>
					<key>Size</key>
					<integer>121123069952</integer>
					<key>VolumeName</key>
					<string>VM</string>
					<key>VolumeUUID</key>
					<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
				</dict>
			</array>
			<key>Content</key>
			<string>XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX</string>
			<key>DeviceIdentifier</key>
			<string>disk1</string>
			<key>Partitions</key>
			<array/>
			<key>Size</key>
			<integer>121123069952</integer>
		</dict>
	</array>
	<key>VolumesFromDisks</key>
	<array>
		<string>Macintosh HD</string>
		<string>VM</string>
	</array>
	<key>WholeDisks</key>
	<array>
		<string>disk0</string>
		<string>disk1</string>
	</array>
</dict>
</plist>`
