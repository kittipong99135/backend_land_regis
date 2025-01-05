package main

import (
	"agent_office/src/database"
	"agent_office/src/routes"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/alecthomas/kingpin/v2"
)

const (
	ProgramName = "agent office"
	Version     = "0.0.0"
)

var (
	startArgs = struct {
		host *net.IP
		port *string
	}{}
)

func init() {
	if err := database.Connect("", ""); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	a := kingpin.New(filepath.Base(os.Args[0]), fmt.Sprintf("%s %s", ProgramName, Version))
	a.Version(Version)
	a.HelpFlag.Short('h')

	startCommand := a.Command("start", "Start server command.")
	startArgs.host = startCommand.Flag("host", "Set server host address.").Envar("SERVER_HOST").Default("0.0.0.0").IP()
	startArgs.port = startCommand.Flag("post", "Set server listen port").Envar("SERVER_PORT").Default("5000").String()

	//Ruc case service
	switch kingpin.MustParse(a.Parse(os.Args[1:])) {
	case startCommand.FullCommand():
		if err := routes.Start("0.0.0.0", "5000"); err != nil {
			log.Println(err)
			os.Exit(1)
		}

	default:
		fmt.Println("Command not found.")
	}
}
