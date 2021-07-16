package main

import "github.com/Ptt-official-app/go-pttbbs/configutil"

const configPrefix = "go-pttbbs"

func InitConfig() error {
	config()
	return nil
}

func setStringConfig(idx string, orig string) string {
	return configutil.SetStringConfig(configPrefix, idx, orig)
}
