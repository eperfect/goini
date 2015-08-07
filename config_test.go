package main

import (
	"cfg"
	"fmt"
)

func main() {
	cfg.InitConfig("config.ini")

	fmt.Println(cfg.GetValue(nil, "ip"))
	//same with
	fmt.Println(GetValue("default", "ip"))

	fmt.Println(cfg.GetValue("server", "ip"))
	SetValue("server", "port", "9090")
	fmt.Println(cfg.GetValue("server", "port"))

	fmt.Println(cfg.GetValue("invalid_section", "ip"))
}
