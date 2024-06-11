package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type NoClingGuide struct{}

func (m *NoClingGuide) Error() string {
	return ".cpp file does not contain // CLING GUIDE"
}

type NoLibs struct{}

func (m *NoLibs) Error() string {
	return "Couldn't find any libs"
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
	return nil, &NoLibs{}
}

func add_libs_to_cmd(libs []string) []string {
	var args []string
	for _, lib := range libs {
		args = append(args, fmt.Sprintf("-l%s", lib))
	}
	return args
}

func main() {
	output_name := flag.String("o", "main", "used to change output file name")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Incorrect usage. Correct usage:\n\tcling example.cpp")
		return
	}

	cmd := "clang++"
	var args []string
	args = append(args, flag.Arg(0))
	args = append(args, fmt.Sprintf("-o%s", *output_name))

	data, err := os.ReadFile(flag.Arg(0))
	check_err(err)

	if line_num, err := contains_cling(data); err == nil {
		libs, err := find_libs(data, line_num)
		check_err(err)
		args = append(args, add_libs_to_cmd(libs)...)
	} else {
		panic(err)
	}
	out := exec.Command(cmd, args...)

	if output, err := out.CombinedOutput(); err != nil {
		fmt.Println(string(output))
		fmt.Println(err.Error())
		return
	}
}
