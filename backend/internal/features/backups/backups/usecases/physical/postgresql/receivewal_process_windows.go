//go:build windows

package usecases_physical_postgresql

import (
	"os/exec"
)

// setReceivewalProcessAttributes is a no-op on Windows because Pdeathsig is Linux-only.
func setReceivewalProcessAttributes(cmd *exec.Cmd) {
	// No-op on Windows
}
