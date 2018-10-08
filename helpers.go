package main

import (
	"encoding/json"
	"github.com/projectriri/bot-gateway/adapters/jsonrpc-server-any/jsonrpc-any"
	"github.com/projectriri/bot-gateway/types"
	"github.com/projectriri/bot-gateway/types/ubm-api"
	"github.com/projectriri/bot-gateway/utils"
	"math/rand"
	"os"
	"path/filepath"
)

const DATA_PATH = "data"

func getFile(path string) (string, error) {
	path = DATA_PATH + "/" + path
	f, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if f.IsDir() {
		files := make([]string, 0)
		err := filepath.Walk(path, visit(&files))
		if err != nil {
			return "", err
		}
		f := files[rand.Intn(len(files))]
		return f, nil
	}
	return path, nil
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			*files = append(*files, path)
		}
		return nil
	}
}

func sendMessage(cid ubm_api.CID, msg ubm_api.Message) {
	msg.CID = &cid
	ubm := ubm_api.UBM{
		Type:    "message",
		Message: &msg,
	}
	b, _ := json.Marshal(ubm)
	C.MakeRequest(jsonrpc_any.ChannelProduceRequest{
		UUID: C.UUID,
		Packet: types.Packet{
			Head: types.Head{
				UUID: utils.GenerateUUID(),
				From: "haruka",
				To:   cid.Messenger,
				Format: types.Format{
					API:     "ubm-api",
					Version: "1.0",
					Method:  "send",
				},
			},
			Body: b,
		},
	})
}

func sendText(cid ubm_api.CID, text string)  {
	msg := ubm_api.Message{
		Type: "rich_text",
		RichText: &ubm_api.RichText{
			{
				Type: "text",
				Text: text,
			},
		},
	}
	sendMessage(cid, msg)
}
