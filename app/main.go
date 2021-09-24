package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/yedf/dtm/dtmcli"
	"github.com/yedf/dtm/dtmsvr"
	"github.com/yedf/dtm/examples"
)

var usage = `dtm is a lightweight distributed transaction manager.

usage:
    dtm [command]

Available commands:
` + commands()

func main() {
	if len(os.Args) == 1 {
		fmt.Println(usage)
		return
	}
	if os.Args[1] != "dtmsvr" { // 实际线上运行，只启动dtmsvr，不准备table相关的数据
		dtmsvr.PopulateDB(true)
		examples.PopulateDB(true)
	}
	dtmsvr.StartSvr()              // 启动dtmsvr的api服务
	go dtmsvr.CronExpiredTrans(-1) // 启动dtmsvr的定时过期查询

	switch os.Args[1] {
	case "quick_start", "qs":
		// quick_start 比较独立，单独作为一个例子运行，方便新人上手
		examples.QsStartSvr()
		examples.QsFireRequest()
	case "dev", "dtmsvr":
	default:
		// 下面是各类的例子
		examples.GrpcStartup()
		examples.BaseAppStartup()

		sample := examples.Samples[os.Args[1]]
		dtmcli.LogIfFatalf(sample == nil, "no sample name for %s", os.Args[1])
		sample.Action()
	}
	select {}
}

func commands() string {
	b := bytes.Buffer{}
	w := tabwriter.NewWriter(&b, 4, 4, 1, ' ', 0)

	type cmd struct {
		name string
		desc string
	}

	cmds := []cmd{
		{"dtmsvr", "run dtm as a server"},
		{"dev", "create all needed table and run dtm as a server"},
		{"quick_start", "run quick start example (dtm will create needed table)"},
		{"qs", "same as quick_start"},
	}

	for name := range examples.Samples {
		cmds = append(cmds, cmd{name, "run sample includes " + strings.ReplaceAll(name, "_", " ")})
	}

	for _, c := range cmds {
		fmt.Fprintf(w, "\t%s\t%s\n", c.name, c.desc)
	}
	w.Flush()

	return b.String()
}
