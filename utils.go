package jxutils

import (
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func GetProgramHome() (h string, err error) {
	file, _ := exec.LookPath(os.Args[0])
	ExecFilePath, _ := filepath.Abs(file)
	var dir string
	execFileSlice := strings.Split(ExecFilePath, `\`)
	HomeDirectory := execFileSlice[:len(execFileSlice)-2]
	osName := runtime.GOOS
	switch osName {
	case "windows":
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
