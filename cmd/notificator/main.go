package main

import (
	"context"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"github.com/magiconair/properties"
	"github.com/robfig/cron"
	"notificator/notificator/internal/config"
	"notificator/notificator/internal/database"
	"notificator/notificator/internal/job"
	"notificator/notificator/internal/server"
)

var (
	Url             string
	PortSend        string
	Cron            string
	CronAddMessage  string
	Log             l4g.Logger
	ServerLog       l4g.Logger
	Version         = "1.0.0"
	LocalServerUrl  string
	LocalServerPort string
)

func main() {
	props := config.ReadPropsConfig("./configs/config.properties", properties.UTF8)
	initLoggers(props)
	defer Log.Close()
	Log.Info(fmt.Sprintf("<-------------------- Application version %s is starting -------------------->", Version))
	initParams(props)
	conn := database.InitDb(props, &Log)
	defer conn.Close(context.Background())
	go database.InitDao(conn, &Log)
	go server.InitServer(LocalServerUrl, LocalServerPort, &ServerLog)
	job.InitJob(&Log, Url, PortSend)
	c := initCron()
	c.Start()
	Log.Info(fmt.Sprintf("<-------------------- Application started -------------------->"))
	for i := 1; i > 0; {
		i++
	}
	//TODO THIS ONE BREAKS JOBS. MAIN DOESN'T RETURN BUT INTERRUPTS?
	//while there are working goroutines main doesn't return
	//runtime.Goexit()
	_ = Log.Error("Thread was interrupted")
}

// initialize loggers
func initLoggers(props *properties.Properties) {
	name := config.ParseField(config.LOG_NAME, props)
	Log = make(l4g.Logger)
	Log.AddFilter(name, l4g.DEBUG, l4g.NewFileLogWriter(name, false))

	serverLogName := config.ParseField(config.SERVER_LOG_NAME, props)
	ServerLog = make(l4g.Logger)
	ServerLog.AddFilter(serverLogName, l4g.DEBUG, l4g.NewFileLogWriter(serverLogName, false))
}

// initialize params from config
func initParams(props *properties.Properties) {
	Url = config.ParseField(config.URL, props)
	Log.Info(fmt.Sprintf("Got URL = \"%s\"", Url))

	PortSend = config.ParseField(config.PORT_SEND, props)
	Log.Info(fmt.Sprintf("Got PortSend = \"%s\"", PortSend))

	Cron = config.ParseField(config.CRON, props)
	Log.Info(fmt.Sprintf("Got Cron = \"%s\"", Cron))

	CronAddMessage = config.ParseField(config.CRON_ADD_MESSAGE, props)
	Log.Info(fmt.Sprintf("Got CronAddMessage = \"%s\"", CronAddMessage))

	LocalServerUrl = config.ParseField(config.SERVER_HOST, props)
	Log.Info(fmt.Sprintf("Got LocalServerUrl = \"%s\"", LocalServerUrl))

	LocalServerPort = config.ParseField(config.SERVER_PORT, props)
	Log.Info(fmt.Sprintf("Got LocalServerPort = \"%s\"", LocalServerPort))
}

// initialize cron jobs
func initCron() *cron.Cron {
	c := cron.New()
	_, err := c.AddFunc(Cron, job.DoJob)
	Log.Info("Added DoJob job")
	if err != nil {
		_ = Log.Error("Error while adding notification job")
	}
	_, err = c.AddFunc(CronAddMessage, job.AddMessage)
	Log.Info("Added AddMessage job")
	if err != nil {
		_ = Log.Error("Error while adding new messages job")
	}
	Log.Info("Created cron jobs")
	return c
}
