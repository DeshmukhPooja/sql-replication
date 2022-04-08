package main

import (
	"/POC/utils/confighelper"

	"github.com/spf13/viper"
)

func main() {
	confighelper.InitViper()

	readDBConfiguration()
}

func readDBConfiguration() {
	primaryDB := viper.GetString("primaryDB")
	secondaryDB := viper.GetString("secondaryDB")
}

// get main DB and replica db from configuration
//initialise them
// create structure of table to insert data
// get table records frequeantly
