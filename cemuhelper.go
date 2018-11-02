// this package helps adding cemu games to steam big-picture mode with different settings
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

// read file and gamepath from cli
var fileName = flag.String("file", "EXAMPLE", "calling -flag EXAMPLE Specifies the file EXAMPLE.bin in the settings directory")
var gamePath = flag.String("game", "filepath", "calling -game FILEPATH Specifies the game loaded by CEMU")

func main() {
	// parse cli args
	flag.Parse()
	// if there are cli args missing, just print the default and exit
	if *fileName == "EXAMPLE" || *gamePath == "filepath" {
		flag.PrintDefaults()
		os.Exit(0)
	}
	// read settings file to be copied
	src, err := os.Open("settings/" + *fileName + ".bin")
	checkErr(err)
	// remove current settings file
	os.Remove("settings.bin")
	// set destination for settings.bin to be copied
	dst, err := os.OpenFile("settings.bin", os.O_RDWR|os.O_CREATE, 0666)
	checkErr(err)
	// start copying
	_, copyErr := io.Copy(dst, src)
	checkErr(copyErr)
	// if we are here, no errors occured and we can start CEMU.exe with cli args
	cmd := exec.Command("CEMU.exe", "-f", "-g", *gamePath)
	cmd.Start()
	os.Exit(0)
}

// helperfunction to check for erros, print them and then exit the program
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
