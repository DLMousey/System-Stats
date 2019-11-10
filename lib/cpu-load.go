package lib

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func ReadProcStat() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")

	if err != nil {
		return
	}

	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)

		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}

				total += val
				if i == 4 {
					idle = val
				}
			}
			return
		}
	}

	return
}

func GetCpuLoad() string {
	idle0, total0 := ReadProcStat()
	time.Sleep(3 * time.Second)
	idle1, total1 := ReadProcStat()

	idleTicks := float64(idle1 - idle0)
	totalTicks := float64(total1 - total0)
	cpuUsage := 100 * (totalTicks - idleTicks) / totalTicks

	return fmt.Sprintf("%f", cpuUsage)
}