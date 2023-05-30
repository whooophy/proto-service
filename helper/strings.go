package helper

import "path/filepath"

func TrimString(args []string) []string {
	var fNames []string
	for _, arg := range args {
		file := filepath.Base(arg)
		ext := filepath.Ext(file)
		// check extension
		if ext == "" {
			fNames = append(fNames, file)
			continue
		}
		fNames = append(fNames, file[:len(file)-len(ext)])
	}
	return fNames
}
