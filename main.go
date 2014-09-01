package main

import (
    "os"
    "huabot-sched/store"
    sch "huabot-sched/sched"
    "huabot-sched/cmd"
    "github.com/codegangsta/cli"
)


func main() {
    app := cli.NewApp()
    app.Name = "huabot-sched"
    app.Usage = ""
    app.Version = "0.0.1"
    app.Flags = []cli.Flag {
        cli.StringFlag{
            Name: "H",
            Value: "unix://huabot-sched.sock",
            Usage: "the server address eg: tcp://127.0.0.1:5000",
            EnvVar: "HUABOT_SCHED_PORT",
        },
        cli.StringFlag{
            Name: "redis",
            Value: "tcp://127.0.0.1:6379",
            Usage: "the redis server address",
            EnvVar: "REDIS_PORT",
        },
        cli.BoolFlag{
            Name: "d",
            Usage: "Enable daemon mode",
        },
    }
    app.Commands = []cli.Command{
        {
            Name: "status",
            Usage: "Show status",
            Action: func(c *cli.Context) {
                cmd.ShowStatus(c.GlobalString("H"))
            },
        },
    }
    app.Action = func(c *cli.Context) {
        if c.Bool("d") {
            sched := sch.NewSched(c.String("H"), store.NewRedisStore(c.String("redis")))
            sched.Serve()
        } else {
            cli.ShowAppHelp(c)
        }
    }

    app.Run(os.Args)
}