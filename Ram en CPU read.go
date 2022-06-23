package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

func main() {
	file, err := os.Create("Log.txt")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	for {
		time.Sleep(time.Second * 5) //pas de 5 naar een ander getal aan als je met een andere frequentie wilt meten.

		cpu, err := cpu.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		memory, err := memory.Get()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}

		dt := time.Now()

		total := float64(cpu.Total)
		cpuUserPerc := float64(cpu.User) / total * 100
		cpuUserAfgerond := fmt.Sprintf("%.2f", cpuUserPerc)
		cpuSysPerc := float64(cpu.System) / total * 100
		cpuSysAfgerond := fmt.Sprintf("%.2f", cpuSysPerc)
		cpuIdlePerc := float64(cpu.Idle) / total * 100
		cpuIdleAfgerond := fmt.Sprintf("%.2f", cpuIdlePerc)

		memTotal := strconv.FormatUint((memory.Total / 1048576), 10)
		memUsed := strconv.FormatUint((memory.Used / 1048576), 10)
		memCached := strconv.FormatUint((memory.Cached / 1048576), 10)
		memFree := strconv.FormatUint((memory.Free / 1048576), 10)

		file.WriteString(dt.Format("01-02-2006 15:04:05") + "\n")
		file.WriteString("cpu user: " + cpuUserAfgerond + " %\n")
		file.WriteString("cpu system: " + cpuSysAfgerond + " %\n")
		file.WriteString("cpu idle: " + cpuIdleAfgerond + " %\n\n")

		file.WriteString("memory total: " + memTotal + " Mb\n")
		file.WriteString("memory used: " + memUsed + " Mb\n")
		file.WriteString("memory cached: " + memCached + " Mb\n")
		file.WriteString("memory free: " + memFree + " Mb\n\n")
	}
}
