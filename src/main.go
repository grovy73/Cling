package main

import (
	"fmt"
	"os"
	"strings"
)

type NoClingGuide struct{}

func (m *NoClingGuide) Error() string {
	return ".cpp file does not contain // CLING GUIDE"
}

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

func contains_cling(content []byte) (int, error) {
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if strings.Contains(line, "CLING GUIDE") {
			return i + 1, nil
		}
	}
	return -1, &NoClingGuide{}
}
func find_libs(content []byte, line_num int) ([]string, error) {
	lines := strings.Split(string(content), "\n")

	if libs := strings.Split(lines[line_num], " "); libs[1] == "libs" {
		return libs[2:], nil
	}
	return nil, &NoClingGuide{}
}

func main() {
	data, err := os.ReadFile("test.cpp")
	check_err(err)

	if line_num, err := contains_cling(data); err == nil {
		libs, err := find_libs(data, line_num)
		check_err(err)
		fmt.Println(libs)
	} else {
		panic(err)
	}

}
