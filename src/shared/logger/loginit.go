package logger

import (
	"github.com/Abramovic/logrus_influxdb"
	log "github.com/Sirupsen/logrus"
	"os"
	"time"
)

func Init(logfile *os.File) {
	// use logrotate.NewFile when log rated by logrotate
	//logfile,err:=logrotate.NewFile(logpath)

	/*
		logfile, err := os.Create(logpath)
		if err != nil {
			log.Error(err)
			return
		}
		//这样创建文件,报bad file descriptor
		//原因是每次从Init()方法return的时候都会执行这个defer，关闭了log文件！！！
		//没有显式的return，defer也会执行
		defer logfile.Close()
	*/

	log.SetOutput(logfile)

	config := &logrus_influxdb.Config{
		Host:        "localhost",
		Port:        8086,
		Database:    "mydb",
		UseHTTPS:    false,
		Measurement: "helloweb",
		Precision:   "ns",
		//Tags:          []string{"server"},
		BatchInterval: (5 * time.Second),
		//TODO BatchCount设置成非零会报错？！
		BatchCount: 0, // set to "0" to disable batching
	}

	/*
	  Use nil if you want to use the default configurations

	  hook, err := logrus_influxdb.NewInfluxDB(nil)
	*/

	hook, err := logrus_influxdb.NewInfluxDB(config)
	if err != nil {
		log.Error(err)
		return
	}
	log.StandardLogger().Hooks.Add(hook)
}
