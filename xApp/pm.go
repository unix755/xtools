package xApp

import (
	"fmt"

	"github.com/unix755/xtools/xApp/packageManager"
	"github.com/unix755/xtools/xApp/packageManager/apk"
	"github.com/unix755/xtools/xApp/packageManager/apt"
	"github.com/unix755/xtools/xApp/packageManager/pkg"
)

type Pm struct {
	packageManager.Pm
}

// NewPm 新建包管理器
func NewPm(pkgManager string, pkgName string) (*Pm, error) {
	var np packageManager.Pm
	switch pkgManager {
	case "apt":
		np = apt.NewPm(pkgName)
	case "pkg":
		np = pkg.NewPm(pkgName)
	case "apk":
		np = apk.NewPm(pkgName)
	default:
		return nil, fmt.Errorf("no supported package manager")
	}
	return &Pm{Pm: np}, nil
}

// Install 安装包
func (p *Pm) Install() (err error) {
	return p.Pm.Install()
}

// Uninstall 卸载包
func (p *Pm) Uninstall() (err error) {
	return p.Pm.Uninstall()
}

// Upgrade 更新包
func (p *Pm) Upgrade() (err error) {
	return p.Pm.Upgrade()
}
