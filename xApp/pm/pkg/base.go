package pkg

import (
	"os/exec"

	"github.com/unix755/xtools/xExec"
)

func Install(pkg string) (err error) {
	return xExec.Run(exec.Command("pkg", "install", "-y", pkg))
}

func Uninstall(pkg string) (err error) {
	err = xExec.Run(exec.Command("pkg", "remove", "-y", pkg))
	if err != nil {
		return err
	}
	return xExec.Run(exec.Command("pkg", "autoremove", "-y"))
}

func Refresh() (err error) {
	return xExec.Run(exec.Command("pkg", "update"))
}

func Update(pkg string) (err error) {
	return xExec.Run(exec.Command("pkg", "upgrade", "-y", pkg))
}
