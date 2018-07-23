package main

import (
	"cli/skycryptor/cmd"
	//"cli/skycryptor/logging"
	//"skycryptor-sdk-go/skycryptor"
)

func main() {
	//logging.SetupLogger()
	//configs.ParseConfig("config.json")
	//skycryptor.CM_init()
	cmd.Execute()
}
