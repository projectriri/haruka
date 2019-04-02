package main

import (
	"encoding/json"
	"github.com/BurntSushi/toml"
	"github.com/projectriri/bot-gateway/adapters/jsonrpc-server-any/client/golang"
	"github.com/projectriri/bot-gateway/router"
	"github.com/projectriri/bot-gateway/types"
	"github.com/projectriri/bot-gateway/types/cmd"
	"github.com/projectriri/bot-gateway/types/ubm-api"
	"io/ioutil"
	"fmt"
)

var C jsonrpc_sdk.Client

func main() {

	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		panic(err)
	}

	C = jsonrpc_sdk.Client{}
	C.Init(config.Addr, config.UUID)

	C.Accept = []router.RoutingRule{
		{
			From: ".*",
			To:   ".*",
			Formats: []types.Format{
				{
					API:      "cmd",
					Version:  "1.0",
					Method:   "cmd",
					Protocol: `{"command_prefix":["!!"],"response_mode":26}`,
				},
			},
		},
	}

	C.Dial()
	pkts, _ := C.GetUpdatesChan(0)
	for pkt := range pkts {
		var command cmd.Command
		json.Unmarshal(pkt.Body, &command)
		switch command.CmdStr {
		case "echo":
			if len(command.ArgsStr) == 0 {
				continue
			}
			sendText(command.Message.Chat.CID, command.ArgsStr)
		case "sticker":
			if len(command.ArgsTxt) == 0 {
				continue
			}
			f, err := getFile("sticker/" + command.ArgsTxt[0])
			if err != nil {
				fmt.Println(err)
				continue
			}
			b, err := ioutil.ReadFile(f)
			if err != nil {
				fmt.Println(err)
				continue
			}
			msg := ubm_api.Message{
				Type: "sticker",
				Sticker: &ubm_api.Sticker{
					Image: &ubm_api.Image{
						Data: b,
					},
				},
			}
			sendMessage(command.Message.Chat.CID, msg)
		case "hitokoto":
			sendText(command.Message.Chat.CID, formatHitokotoRespMsg(command.ArgsTxt))
		}
	}
}
