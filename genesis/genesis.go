package main

import (
	"time"
	"github.com/pterm/pterm"
	"github.com/leaanthony/clir"
	"strconv"
	"io/ioutil"
	"encoding/json"
)

func to_file(){
	
}

func tim(duration int,buffer int){
	area, _ := pterm.DefaultArea.WithCenter().Start() 
	pterm.DefaultCenter.Println("Start")
	pterm.DefaultCenter.Println("You have a "+strconv.Itoa(buffer)+" second buffer period before the timer starts\n\n\n\n\n\n")
	
	for i:=0;i<=buffer;i++{
		time.Sleep(time.Second * 1)
	}

	for i:=0;i<=duration;i++{

		str, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(strconv.Itoa(i))).Srender()

		time.Sleep(time.Second * 1)

		area.Update(str)
	}
	area.Stop()
}

func main() {
	cli := clir.NewCli("Flags", "A simple example", "v0.0.1")

	// timer function. sets timer according to passed value and buffer period
	timer:=cli.NewSubCommand("timer","opens timer and sub commands")
	var set int
	timer.IntFlag("set","sets the time for the timer",&set)
	var buf int
	timer.IntFlag("buffer","set buffer period",&buf)

	timer.Action(func() error{
		if set & buf > 0 {
			tim(set,buf)
		}

		return nil
	})

	// subjects. search for subjects for the day
	subjects:=cli.NewSubCommand("sub","search for subjects")

	var sub string
	subjects.StringFlag("today","shows today's subjects",&sub)

	cli.Run()
}