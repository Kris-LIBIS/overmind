package main

import (
	"fmt"

	"github.com/Kris-LIBIS/overmind/utils"

	"github.com/urfave/cli"
)

type cmdQuitHandler struct{ dialer }

func (h *cmdQuitHandler) Run(_ *cli.Context) error {
	conn, err := h.Dial()
	utils.FatalOnErr(err)

	fmt.Fprintf(conn, "quit")

	return nil
}
