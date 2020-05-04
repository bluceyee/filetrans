package handler

import (
	"encoding/json"
	"fmt"
	"os"

	"filetrans/trans/message"

	"github.com/fengdingfeilong/roshan"

	rhandler "github.com/fengdingfeilong/roshan/handler"
	"github.com/fengdingfeilong/roshan/roshantool"
)

type TransferComplete struct {
	rhandler.Base
	server *roshan.Server
}

func NewTransferComplete(s *roshan.Server) *TransferComplete {
	return &TransferComplete{server: s}
}

func (h *TransferComplete) GetBase() *rhandler.Base {
	return &h.Base
}

func (h *TransferComplete) Execute(data []byte) {
	var msg message.TransferComplete
	err := json.Unmarshal(data, &msg)
	if err != nil {
		roshantool.PrintErr("handler.transfercomplete.Execute", err.Error())
		return
	}
	fmt.Println("all files transfered complete")
	roshantool.Println("all files transfered complete")
	os.Exit(0)
}
