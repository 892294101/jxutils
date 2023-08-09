package jxutils

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
)

func GetProgramHome() (h string, err error) {
	file, _ := exec.LookPath(os.Args[0])
	ExecFilePath, _ := filepath.Abs(file)
	var dir string
	osName := runtime.GOOS
	switch osName {
	case "windows":
		execfileslice := strings.Split(ExecFilePath, `\`)
		HomeDirectory := execfileslice[:len(execfileslice)-2]
		for i, v := range HomeDirectory {
			if v != "" {
				if i > 0 {
					dir += `\` + v
				} else {
					dir += v
				}
			}
		}
	case "linux", "darwin":
		execfileslice := strings.Split(ExecFilePath, "/")
		HomeDirectory := execfileslice[:len(execfileslice)-2]
		for _, v := range HomeDirectory {
			if v != "" {
				dir += `/` + v
			}
		}
	default:
		return "", errors.Errorf("unsupported operating system type: %s", osName)
	}

	if dir == "" {
		return "", errors.Errorf("get program home directory failed: %s", dir)
	}
	return dir, nil
}

func JsonStr2Bson(str string) (interface{}, error) {
	var want interface{}
	err := bson.UnmarshalExtJSON([]byte(str), true, &want)
	if err != nil {
		return nil, err
	}
	return want, nil
}

// 自定义Panic异常处理,调用方式: 例如Test()函数, 指定defer ErrorCheckOfRecover(Test)
func GetFunctionName(i interface{}, seps ...rune) string {
	u := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Entry()
	f, _ := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).FileLine(u)
	return f
}

var GlobalProcessID string

func ErrorCheckOfRecover(n interface{}, log *logrus.Logger) {
	if err := recover(); err != nil {
		home, _ := GetProgramHome()
		if len(GlobalProcessID) > 0 {
			_ = os.Remove(filepath.Join(home, "pcs", GlobalProcessID))
		}
		log.Errorf("Panic Message: %s", err)
		log.Errorf("Exception File: %s", GetFunctionName(n))
		log.Errorf("Print Stack Message: %s", string(debug.Stack()))
	}
}
