package ga

import (
	"github.com/hofstadter-io/yagu"

	"github.com/hofstadter-io/dma/cmd/dma/verinfo"
)

func SendGaEvent(action, label string, value int) {
	// Do something here to lookup / create
	cid := "13b3ad64-9154-11ea-9eba-47f617ab74f7"

	cfg := yagu.GaConfig{
		TID: "UA-103579574-5",
		CID: cid,
		UA:  "dma/" + verinfo.Version,
		CS:  "dma",
		CN:  verinfo.Version,
	}

	evt := yagu.GaEvent{
		Source:   cfg.UA,
		Category: "dma",
		Action:   action,
		Label:    label,
	}

	if value >= 0 {
		evt.Value = value
	}

	yagu.SendGaEvent(cfg, evt)
}