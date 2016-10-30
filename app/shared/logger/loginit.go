package logger

import (
	log "github.com/Sirupsen/logrus"
	"github.com/jiada8866/logrus_influxdb"
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
		Tags:        []string{"server", "api", "type"},
		//当启动程序后，BatchInterval内没有打日志，在BatchInterval间隔后触发写influxDB会触发空指针的panic
		//因为此时hook.batchP==nil
		BatchInterval: (5 * time.Second),
		BatchCount:    200, // set to "0" to disable batching
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
