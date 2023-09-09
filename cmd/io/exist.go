package io

import "os"

func IsDirExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsFileExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
