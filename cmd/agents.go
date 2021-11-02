package cmd

import (
	"fmt"
	"time"
	"encoding/json"

	"github.com/latortuga71/MeowthCoreCLI/internal"
	"github.com/desertbit/grumble"
)


type Agent struct {
	Metadata Metadata  `json:"metadata"`
	LastSeen time.Time`json:"lastSeen"`
}
type Metadata struct {
	ID           string `json:"id"`
	Hostname     string `json:"hostname"`
	Username     string `json:"username"`
	ProcessName  string `json:"processName"`
	ProcessID    int    `json:"processId"`
	Architecture string `json:"architecture"`
	Integrity    string `json:"integrity"`
}


func init(){
    agentCommand := &grumble.Command{
        Name:"agent",
        Help:"Interact with agents",
    }
    App.AddCommand(agentCommand)
    agentCommand.AddCommand(&grumble.Command{
        Name:"get",
        Help:"get agent",
        Run: getAgent,
        Args: func (a *grumble.Args){
            a.String("id","agent id")
        },
    })
    agentCommand.AddCommand(&grumble.Command{
        Name:"list",
        Help:"List agents",
        Run: listAgents,
    })
}


func getAgent(c *grumble.Context) error {
	id := c.Args.String("id")
	var agent Agent
	target := fmt.Sprintf("%s%s%s",FullServerURI,"Agents/",id)
	err := internal.Get(target, &agent)
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(agent,""," ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
    return nil
}

func listAgents(c *grumble.Context) error{
	var agents []Agent
	target := fmt.Sprintf("%s%s",FullServerURI,"Agents")
	err := internal.Get(target, &agents)
	if err != nil {
		return err
	}
	for _,a := range agents{
		data, err := json.Marshal(a)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	}
    return nil
}
