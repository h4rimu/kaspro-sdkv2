package utils

import "github.com/h4rimu/kaspro-sdkv2/config"

func GetRunMode() string {
	serverMode := config.MustGetString("server.mode")
	return serverMode
}
