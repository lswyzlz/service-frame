package common

import (
	"os"
	"path/filepath"
)

//GetAppPath 获取程序运行目录
func GetAppPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}
