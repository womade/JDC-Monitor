package main

import (
	"JDC-Monitor/cmd"
	"fmt"
	"github.com/mizuki1412/go-core-kit/init/initkit"
)

var (
	version   string
	date      string
	goVersion string
)

func main() {
	fmt.Println(`


     _ ____   ____      __  __             _ _
    | |  _ \ / ___|    |  \/  | ___  _ __ (_) |_ ___  _ __
 _  | | | | | |   _____| |\/| |/ _ \| '_ \| | __/ _ \| '__|
| |_| | |_| | |__|_____| |  | | (_) | | | | | || (_) | |
 \___/|____/ \____|    |_|  |_|\___/|_| |_|_|\__\___/|_|

===========================================================
       Copyright © 1998 中國遠徵 All Rights Reserved
===========================================================

`)
	info := fmt.Sprintf("软件版本：%s \n 编译日期：%s \n 编译环境：%s \n\n\n", version, date, goVersion)
	fmt.Print(info)
	initkit.LoadConfig()
	cmd.Execute()
}
