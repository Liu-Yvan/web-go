// log library
package log

import (
	"os"
	"fmt"
	"strings"

	"path/filepath"
	"encoding/json"
	"github.com/pkg/errors"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/config"
)

const configPath = "config\\log.json"

type LogConfig struct {
	Filename string `json:"filename"`
	MaxLines int `json:"maxlines"`
	MaxSize  int `json:"maxsize"`
	Daily    bool  `json:"daily"`
	MaxDays  int64 `json:"maxdays"`
	Rotate   bool `json:"rotate"`
	Perm     string `json:"perm"`
}

type FileConfig struct {
	Filename string `json:"filename"`
}

// 获取默认配置
func getDefaultConfig() config.Configer {
	conf, err := config.NewConfig("json", configPath)
	if err != nil {
		panic("read log config failed")
	}
	return conf
}

// 获取文件名
func getFileName(filename string) string {
	conf := getDefaultConfig()
	path := conf.DefaultString("dir", "log") + "\\" + filename
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0600)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	return path
}

// 获取BeeLogger对象
func GetLogger(filename string) *logs.BeeLogger {
	// 读取配置
	conf := getDefaultConfig()

	// 日志配置
	config := new(LogConfig)
	config.Filename = getFileName(filename)
	config.MaxLines = conf.DefaultInt("maxLines", 100000)
	config.MaxSize = conf.DefaultInt("maxsize", 1024000) //100MB
	config.Daily = conf.DefaultBool("daily", true)
	config.MaxDays = conf.DefaultInt64("maxdays", 30)
	config.Rotate = conf.DefaultBool("rotate", true)
	config.Perm = conf.DefaultString("perm", "0600")

	data, err := json.Marshal(config)
	if err != nil {
		panic(err)
	}

	// 设置配置
	var l = logs.NewLogger()
	if conf.DefaultBool("console", true) {
		l.SetLogger(logs.AdapterConsole)
		return l
	}

	if err := l.SetLogger(logs.AdapterFile, string(data)); err != nil {
		panic(err)
	}

	if conf.DefaultBool("async", true) {
		l.Async(1e3)
	}

	return l
}

// 格式化
func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

// 业务日志
func Info(filename string, f interface{}, v ...interface{}) {
	l := GetLogger(filename)
	l.Info(formatLog(f, v...))
}

// 错误日志，添加了堆栈信息
func Error(filename string, f interface{}, v ...interface{}) {
	l := GetLogger(filename)
	l.Error(fmt.Sprintf("%+v", errors.New(formatLog(f, v...))))
}


