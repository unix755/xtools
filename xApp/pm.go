package xApp

import (
	"fmt"

	"github.com/unix755/xtools/xApp/pm"
	"github.com/unix755/xtools/xApp/pm/apk"
	"github.com/unix755/xtools/xApp/pm/apt"
	"github.com/unix755/xtools/xApp/pm/pkg"
)

type Pm struct {
	pm.Pm
}

// NewPm 新建包管理器
func NewPm(pkgManager string, pkgName string) (*Pm, error) {
	var np pm.Pm
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

// Update 更新包
func (p *Pm) Update() (err error) {
	return p.Pm.Update()
}
