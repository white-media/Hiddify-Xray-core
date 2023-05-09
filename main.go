package main

import (
	"encoding/json"
	"fmt"
	"github.com/xtls/xray-core/core"
	"github.com/xtls/xray-core/infra/conf"
	"os"
)

func main() {
	//read config from "config.json
	file, err := os.ReadFile("config.json")
	if err != nil {
		fmt.Println("read file err: ", err)
		return
	}

	config := &conf.Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		fmt.Println("unmarshal config err: ", err)
		return
	}

	//create core config from our json
	coreConfig, err := config.Build()
	if err != nil {
		fmt.Println("build core config err: ", err)
		return
	}

	//create new instance to start
	instance, err := core.New(coreConfig)
	if err != nil {
		fmt.Println("c	reate new instance err: ", err)
		return
	}
	fmt.Println("instance: ", instance)

}
