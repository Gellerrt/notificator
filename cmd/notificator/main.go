package main

import (
	"flag"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"notificator/internal/config"
	"notificator/internal/notificator"
)

//var (
//	Url             string
//	PortSend        string
//	Cron            string
//	CronAddMessage  string
//	Log             l4g.Logger
//	ServerLog       l4g.Logger
//	Version         = "1.0.0"
//	LocalServerUrl  string
//	LocalServerPort string
//)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/notificator.yaml", "path to config file")
}

func main() {
	flag.Parse()
	conf, _ := initConfig()
	notificator := notificator.New(conf)
	err := notificator.Start()
	if err != nil {
		//bad
	}
	/*	initLoggers(conf)
		defer Log.Close()
		Log.Info(fmt.Sprintf("<-------------------- Application version %s is starting -------------------->", Version))
		initParams(props)
		conn := database.InitDb(props, &Log)
		defer conn.Close()
		go database.InitDao(conn, &Log)
		go server.InitServer(LocalServerUrl, LocalServerPort, &ServerLog)
		job.InitJob(&Log, Url, PortSend)
		c := initCron()
		c.Start()
		Log.Info(fmt.Sprintf("<-------------------- Application started -------------------->"))
		for i := 1; i > 0; {
			i++
		}
		_ = Log.Error("Thread was interrupted")*/
}

func initConfig() (*config.Config, error) {
	conf := config.NewConfig()
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return conf, err
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return conf, err
	}
	return conf, nil
}

/*
// initialize loggers
func initLoggers(props *config.Config) {
	name := config.ParseField(config.LOG_NAME, props)
	Log = make(l4g.Logger)
	Log.AddFilter(name, l4g.DEBUG, l4g.NewFileLogWriter(name, false))

	serverLogName := config.ParseField(config.SERVER_LOG_NAME, props)
	ServerLog = make(l4g.Logger)
	ServerLog.AddFilter(serverLogName, l4g.DEBUG, l4g.NewFileLogWriter(serverLogName, false))
}*/
/*
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
}*/
/*
// initialize cron jobs
func initCron() *cron.Cron {
	c := cron.New()
	err := c.AddFunc(Cron, job.DoJob)
	Log.Info("Added DoJob job")
	if err != nil {
		_ = Log.Error("Error while adding notification job")
	}
	err = c.AddFunc(CronAddMessage, job.AddMessage)
	Log.Info("Added AddMessage job")
	if err != nil {
		_ = Log.Error("Error while adding new messages job")
	}
	Log.Info("Created cron jobs")
	return c
}*/
