package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	monthNames := map[string]string{
		"1":  "ENERO",
		"2":  "FEBRERO",
		"3":  "MARZO",
		"4":  "ABRIL",
		"5":  "MAYO",
		"6":  "JUNIO",
		"7":  "JULIO",
		"8":  "AGOSTO",
		"9":  "SEPTIEMBRE",
		"10": "OCTUBRE",
		"11": "NOVIEMBRE",
		"12": "DICIEMBRE",
	}
	start := time.Now()
	Banner()
	filegenerate, filereception := ReadEnv()
	if len(filegenerate) == 0 || len(filereception) == 0 {
		ExitProgram("\tAlguna variable de entorno no esta siendo declarada")
		return
	}
	FileMove(filegenerate, monthNames)
	FileMove(filereception, monthNames)

	elapsed := time.Since(start)
	ExitProgram("\tTiempo transcurrido: " + elapsed.String())
}

func Banner() {
	fmt.Println("\t" + `   ______     _____            __ `)
	fmt.Println("\t" + `  / ____/___ / ___/____  _____/ /_`)
	fmt.Println("\t" + ` / / __/ __ \\__ \/ __ \/ ___/ __/`)
	fmt.Println("\t" + `/ /_/ / /_/ /__/ / /_/ / /  / /_  `)
	fmt.Println("\t" + `\____/\____/____/\____/_/   \__/  `)
	fmt.Println(`==================================================================`)
	fmt.Println("\t" + `Empezando a organizar los archivos del servidor SIVE`)
}

func ExitProgram(message string) {
	fmt.Println(message)
	fmt.Println("==================================================================")
	fmt.Println("\t\tgracias por usar la aplicacion")
	fmt.Println("\t\tBy - Yosimar Zahid Aquino Sosa")
	fmt.Println("==================================================================")
	fmt.Println("\tPresiona Enter para finalizar el programa...")
	bufio.NewReader(os.Stdin).ReadByte()
}

func FileMove(directory string, monthNames map[string]string) {
	dirRead, _ := os.Open(directory)
	dirFiles, _ := dirRead.Readdir(0)
	for fileIndex := range dirFiles {
		currentFile := dirFiles[fileIndex]
		destiny := directory + strconv.Itoa(currentFile.ModTime().Year()) + string(os.PathSeparator) + monthNames[currentFile.ModTime().Month().String()] + string(os.PathSeparator) + strings.Split(currentFile.Name(), ".")[0] + string(os.PathSeparator)
		os.MkdirAll(destiny, 0755)
		os.Rename(directory+currentFile.Name(), destiny+currentFile.Name())
	}
}

func ReadEnv() (string, string) {
	return os.Getenv("FileGenerate"), os.Getenv("FileReception")
}
