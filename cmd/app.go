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
var PendingTaskQueue struct {
	sync.Mutex
	results []TaskResult
}
var AgentList struct {
    sync.Mutex
    agents []Agent
}
var CurrentAgent *Agent 

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
}
