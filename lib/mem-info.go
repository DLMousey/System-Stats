package lib

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func ReadMemInfo() (total, free uint64) {
	contents, err := ioutil.ReadFile("/proc/meminfo")

	if err != nil {
		return
	}

	lines := strings.Split(string(contents), "\n")
	fmt.Println(lines[0], lines[1])

	reg, err := regexp.Compile("[^0-9]")
	if err != nil {
		fmt.Println("Failed to copmile regex")
	}

	total, err = strconv.ParseUint(reg.ReplaceAllString(lines[0], ""), 10, 64)
	free, err = strconv.ParseUint(reg.ReplaceAllString(lines[1], ""), 10, 64)

	if err != nil {
		fmt.Println("Error parsing meminfo: ", err)
	}

	return
}