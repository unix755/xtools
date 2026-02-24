package apk

import (
	"os/exec"

	"github.com/unix755/xtools/xExec"
)

func Install(pkg string) error {
	return xExec.Run(exec.Command("apk", "add", pkg))
}

func Uninstall(pkg string) error {
	return xExec.Run(exec.Command("apk", "del", pkg, "--purge"))
}

func Upgrade(pkg string) error {
	return xExec.Run(exec.Command("apk", "upgrade", pkg))
}

func update() error {
	return xExec.Run(exec.Command("apk", "update"))
}
