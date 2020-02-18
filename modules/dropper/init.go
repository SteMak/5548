package dropper

import "github.com/SteMak/vanilla/modules"

func init() {
	m := new(module)
	modules.Register(m.ID(), m)
}
