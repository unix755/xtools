package apk

type Pm struct {
	Name string
}

func NewPm(name string) (p *Pm) {
	return &Pm{Name: name}
}

func (p *Pm) Install() (err error) {
	err = update()
	if err != nil {
		return err
	}
	return Install(p.Name)
}

func (p *Pm) Uninstall() (err error) {
	return Uninstall(p.Name)
}

func (p *Pm) Upgrade() (err error) {
	err = update()
	if err != nil {
		return err
	}
	return Upgrade(p.Name)
}
