package apt

import (
	"os/exec"

	"github.com/unix755/xtools/xExec"
)

func Install(pkg string) (err error) {
	return xExec.Run(exec.Command("apt", "install", "-y", pkg))
}

func Uninstall(pkg string) (err error) {
	return xExec.Run(exec.Command("apt", "purge", "-y", "--autoremove", pkg))
}

func Refresh() (err error) {
	return xExec.Run(exec.Command("apt", "update"))
}

func Update(pkg string) (err error) {
	return xExec.Run(exec.Command("apt", "upgrade", "-y", pkg))
}
