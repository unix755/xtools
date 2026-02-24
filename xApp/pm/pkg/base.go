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

func Upgrade(pkg string) (err error) {
	return xExec.Run(exec.Command("pkg", "upgrade", "-y", pkg))
}

func update() (err error) {
	return xExec.Run(exec.Command("pkg", "update"))
}
