package flag

import (
	"flag"
	"fmt"
)

func Ssh() {

	//ssh -H 127.0.0.1 --port=8080 -v -h touch xxxx

	// 定义变量
	// 定义命令行格式
	// 元素 =>解析到 变量, 默认值帮助信息

	//定义变量
	var (
		host    string
		port    int
		verbose bool
		help    bool
	)

	// 定义命令行格式
	flag.StringVar(&host, "H", "127.0.0.1", "连接地址")
	flag.IntVar(&port, "port", 8080, "连接端口")
	flag.BoolVar(&verbose, "v", false, "显示详情")
	flag.BoolVar(&help, "h", false, "帮助信息")

	flag.Usage() = func() {
		fmt.println("ssh -H 127.0.0.1 --port=8080 -v -h touch xxxx")
		flag.PrintDefaults()
	}

	// 解析
	flag.Parse()

	if help {
		flag.Usage()
		return
	}

	fmt.Println(host, port, verbose, help, flag.Args())

}
