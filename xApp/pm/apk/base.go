package apk

import (
	"os/exec"

	"github.com/unix755/xtools/xExec"
)

func Install(pkg string) (err error) {
	return xExec.Run(exec.Command("apk", "add", pkg))
}

func Uninstall(pkg string) (err error) {
	return xExec.Run(exec.Command("apk", "del", pkg, "--purge"))
}

func Refresh() (err error) {
	return xExec.Run(exec.Command("apk", "update"))
}

func Update(pkg string) (err error) {
	return xExec.Run(exec.Command("apk", "upgrade", pkg))
}
