package uuid

import (
	"bytes"
	"fmt"
	"net"

	uuid "github.com/satori/go.uuid"
)

// GenerateDeviceID generate a UUID from MAC address
func GenerateDeviceID() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, i := range interfaces {
		if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
			// Skip locally administered addresses
			if i.HardwareAddr[0]&2 == 2 {
				continue
			}
			space, err := uuid.FromString("00000000-0000-0000-0000-000000000000")
			if err != nil {
				fmt.Println(err)
				return ""
			}
			id := uuid.NewV5(space, string(i.HardwareAddr))
			return id.String()
		}
	}
	return ""
}
