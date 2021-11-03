package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/desertbit/grumble"
	"github.com/latortuga71/MeowthCoreCLI/internal"
)

type TaskResult struct {
    Id string `json:"id"`
    Result string `json:"result"`

}

type Task struct {
    TaskId string `json:"id"`
	Command string   `json:"command"`
	Args    []string `json:"args"`
	File    string   `json:"file"`
}

func init(){
    taskCommand := &grumble.Command{
        Name:"task",
        Help:"Interact with tasks",
    }
    App.AddCommand(taskCommand)
    taskCommand.AddCommand(&grumble.Command{
        Name:"get",
        Help:"get task by agentid and task id ",
        Run: getTask,
        Args: func (a *grumble.Args){
            a.String("taskId","task id")
            a.String("agentId","agent id")
        },
    })
    taskCommand.AddCommand(&grumble.Command{
        Name:"list",
        Help:"List tasks by agent id",
        Run: listTasks,
        Args: func (a *grumble.Args){
            a.String("agentId","agent id")
        },
    })
}


func getTask(c *grumble.Context) error {
	id := c.Args.String("agentId")
    tid := c.Args.String("taskId")
	var task TaskResult
	target := fmt.Sprintf("%s%s%s%s%s",FullServerURI,"Agents/",id,"/tasks/",tid)
	err := internal.Get(target, &task)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(task,""," ")
	if err != nil {
		return err
    }
	fmt.Println(string(data))
    return nil
}

func listTasks(c *grumble.Context) error{
	id := c.Args.String("agentId")
	var tasks []TaskResult
	target := fmt.Sprintf("%s%s%s%s",FullServerURI,"Agents/",id,"/tasks")
    fmt.Println(target)
	err := internal.Get(target, &tasks)
	if err != nil {
		return err
	}
	for _,t := range tasks{
		data, err := json.Marshal(t)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	}
    return nil
}
