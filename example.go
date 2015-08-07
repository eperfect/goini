package goini

import (
	"fmt"
)

func main() {
	InitConfig("config.ini")

	fmt.Println(GetValue("", "ip"))
	//same with
	fmt.Println(GetValue("default", "ip"))

	fmt.Println(GetValue("server", "ip"))
	SetValue("server", "port", "9090")
	fmt.Println(GetValue("server", "port"))

	fmt.Println(GetValue("invalid_section", "ip"))
}
