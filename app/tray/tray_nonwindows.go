//go:build !windows

package tray

import (
	"fmt"

	"github.com/a4to/ALM/app/tray/commontray"
)

func InitPlatformTray(icon, updateIcon []byte) (commontray.OllamaTray, error) {
	return nil, fmt.Errorf("NOT IMPLEMENTED YET")
}
