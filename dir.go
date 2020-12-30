package fileutil

import (
	"io/ioutil"
)

// AllSubfolders gets all subfolders under given dir.
// Hidden folders are ignored.
func AllSubfolders(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var result []string

	for _, f := range files {
		if f.IsDir() && !IsHiddenFile(f) {
			result = append(result, f.Name())
		}
	}

	return result, err
}
