package common

import "os"

func ReadFile(filename string) (string, error) {
	contents, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}
