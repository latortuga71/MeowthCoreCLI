package cmd

import (
	"sync"

	"github.com/desertbit/grumble"
	"github.com/fatih/color"
)

// Global Server HostName And Port

var ServerIP string 
var ServerPort string
var FullServerURI string
var ServerSet bool
var AgentList struct {
    sync.Mutex
    agents []Agent
}
var CurrentAgentId string

var App = grumble.New(&grumble.Config{
	Name:                  "meowth",
	Description:           "Meowth C2 Command line interface.",
	HistoryFile:           "/tmp/meowth.hist",
	Prompt:                "meowth » ",
	PromptColor:           color.New(color.FgCyan, color.Bold),
	HelpHeadlineColor:     color.New(color.FgCyan),
	HelpHeadlineUnderline: true,
	HelpSubCommands:       false,
})

func CheckConnection() bool {
    return ServerSet
}

func init() {
    /*CmdArgs := os.Args[1:]
    if len(CmdArgs) < 2 {
        fmt.Println("cli <ip> <port>")
    }
    var serverPtr string
    var portPtr string
    serverPtr = os.Args[1]
    portPtr = os.Args[2]
    //flag.StringVar(&serverPtr,"h","","ip of c2 server")
    //flag.StringVar(&portPtr,"p","","port of c2 server")
    //flag.Parse()
    FullServerURI = fmt.Sprintf("http://%s:%d/",serverPtr,portPtr)
    ServerSet = true
    */
}
