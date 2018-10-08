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
					Protocol: "",
					Method:   "cmd",
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
		case "haruka:echo":
			if len(command.ArgsStr) == 0 {
				continue
			}
			msg := ubm_api.Message{
				Type: "rich_text",
				RichText: &ubm_api.RichText{
					{
						Type: "text",
						Text: command.ArgsStr,
					},
				},
			}
			sendMessage(command.Message.Chat.CID, msg)
		case "haruka:sticker":
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
		case "haruka:hitokoto":
			msg := ubm_api.Message{
				Type: "rich_text",
				RichText: &ubm_api.RichText{
					{
						Type: "text",
						Text: formatHitokotoRespMsg(command.ArgsTxt),
					},
				},
			}
			sendMessage(command.Message.Chat.CID, msg)
		case "ping":
			fallthrough
		case "haruka:ping":
			msg := ubm_api.Message{
				Type: "rich_text",
				RichText: &ubm_api.RichText{
					{
						Type: "text",
						Text: formatPingRespMsg(),
					},
				},
			}
			sendMessage(command.Message.Chat.CID, msg)
		}
	}
}
