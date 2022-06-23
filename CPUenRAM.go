package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"
)

var total, cpuUserPerc, cpuSysPerc, cpuIdlePerc float64
var memTotalMB, memUsedMB, memFreeMB int
var location string

func CPU() {
	cpu, err := cpu.Get() //hier worden de cpu gegevens opgehaald
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total = float64(cpu.Total)
	cpuUserPerc = float64(cpu.User) / total * 100 //hier reken ik de cpu gebruik om naar een percentage
	cpuSysPerc = float64(cpu.System) / total * 100
	cpuIdlePerc = float64(cpu.Idle) / total * 100
}

func Memory() {
	memory, err := memory.Get() //Hier wordt de ramgebruik ophehaald
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	memTotalMB = int(memory.Total) / 1048576
	memUsedMB = int(memory.Used) / 1048576
	memFreeMB = int(memory.Free) / 1048576
}

func write_to_file() {
	cpuUserAfgerond := fmt.Sprintf("%.2f", cpuUserPerc) //hier geef ik aan dat er maar 2 decimalen achter de komma komen
	cpuSysAfgerond := fmt.Sprintf("%.2f", cpuSysPerc)
	cpuIdleAfgerond := fmt.Sprintf("%.2f", cpuIdlePerc)

	memTotalMBstring := strconv.Itoa(memTotalMB) //hier wordt de int naar string omgezet
	memUsedMBstring := strconv.Itoa(memUsedMB)
	memFreeMBstring := strconv.Itoa(memFreeMB)

	file, err := os.OpenFile(location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //hier wordt een logbestand aangemaakt als die nog niet bestaat, als deze al bestaat wordt er in de bestaande log geschreven
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	Tijd := time.Now()
	file.WriteString(Tijd.Format("01-02-2006 15:04:05") + "\n") //hier wordt de datum en tijd geprint
	file.WriteString("cpu user: " + cpuUserAfgerond + " %\n")   //hier wordt de cpu geprint
	file.WriteString("cpu system: " + cpuSysAfgerond + " %\n")
	file.WriteString("cpu idle: " + cpuIdleAfgerond + " %\n\n")

	file.WriteString("memory total: " + memTotalMBstring + " Mb\n") //hier wordt de ram geprint
	file.WriteString("memory used: " + memUsedMBstring + " Mb\n")
	file.WriteString("memory free: " + memFreeMBstring + " Mb\n\n")
}

func main() {
	scanner := bufio.NewScanner(os.Stdin) //hier geef ik aan dat de user iets in de terminal moet typen zodat er een logbestand gemaakt wordt
	fmt.Println("Vul hier de log locatie in: (Voorbeeld = ./Log.txt)")
	scanner.Scan()
	location = scanner.Text()
	//hier geef ik aan dat de user iets in de terminal moet typen zodat de frequentie bepaald kan worden
	fmt.Println("Vul hier de log frequentie in: (Voorbeeld = 5s)")
	scanner.Scan()
	tijd, err := time.ParseDuration(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}

	for {
		time.Sleep(tijd)
		CPU()
		Memory()
		write_to_file()
	}
}
