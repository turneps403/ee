package yaml

import (
	"log"
	"os"
	"path/filepath"
)

func GetFiles(dir, file string) []string {
	if len(dir) == 0 {
		path, err := os.Getwd()
		if err != nil {
			log.Println(err)
		}
		dir = path
	}

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	homeFile := filepath.Join(home, file)

	ret := []string{}
	for len(dir) > 0 {
		fullPath := filepath.Join(dir, file)
		if fullPath != homeFile {
			if isFileReadable(fullPath) {
				ret = append(ret, fullPath)
			}
		}
		newDir := filepath.Dir(dir)
		if newDir == dir {
			break
		}
		dir = newDir
	}

	if isFileReadable(homeFile) {
		ret = append(ret, homeFile)
	}

	return ret
}

func isFileReadable(path string) bool {
	finfo, err := os.Stat(path)
	if err != nil || finfo.IsDir() || finfo.Size() == 0 {
		return false
	}

	if file, err := os.Open(path); err != nil {
		return false
	} else {
		defer file.Close()
	}
	return true
}
