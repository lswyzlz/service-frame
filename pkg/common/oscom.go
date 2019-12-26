package common

import (
	"os"
)

//DirExists  判断文件夹是否存在
func DirExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//MakeDir 如果文件夹不存在，创建文件夹
func MakeDir(dir string) error {
	if b, _ := DirExists(dir); b {
		return nil
	}

	// 创建文件夹
	err := os.Mkdir(dir, os.ModePerm)
	return err
}
