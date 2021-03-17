package commands

import "fmt"

type Help struct {
}

const message = `命令:
	help 帮助/使用说明
	run 启动app
`

func (command *Help) Run() {
	fmt.Print(message)
}