package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

var total, cpuUserPerc, cpuSysPerc, cpuIdlePerc float64

func cpu_shit() {
	cpu, err := cpu.Get() //hier worden de cpu gegevens opgehaald
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total = float64(cpu.Total)
	cpuUserPerc = float64(cpu.User) / total * 100
	cpuSysPerc = float64(cpu.System) / total * 100
	cpuIdlePerc = float64(cpu.Idle) / total * 100

}

func write_to_file(file *File) {

	cpuUserAfgerond := fmt.Sprintf("%.2f", cpuUserPerc) //hier geef ik aan dat er maar 2 decimalen achter de komma komen
	cpuSysAfgerond := fmt.Sprintf("%.2f", cpuSysPerc)
	cpuIdleAfgerond := fmt.Sprintf("%.2f", cpuIdlePerc)

	//de rest dat nog moet :)

	Tijd := time.Now()
	file.WriteString(Tijd.Format("01-02-2006 15:04:05") + "\n") //hier wordt de datum en tijd geprint
	file.WriteString("cpu user: " + cpuUserAfgerond + " %\n")   //hier wordt de cpu geprint
	file.WriteString("cpu system: " + cpuSysAfgerond + " %\n")
	file.WriteString("cpu idle: " + cpuIdleAfgerond + " %\n\n")
}

func main() {
	//vraag de user om een locatie, mag jij doen e.g: ./log.txt :)
	location := "/Log.txt"
	//vraag de user om een tijd, zie https://www.geeksforgeeks.org/time-parse-function-in-golang-with-examples/ voor parsing :)
	tijd := time.Second * 5

	file, err := os.Create(location) //hier wordt een logbestand aangemaakt, //handle de error :)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	for {
		time.Sleep(tijd) //maak dit ook configureerbaar door te vragen

		cpu_shit()
		write_to_file(file)
	}
}
