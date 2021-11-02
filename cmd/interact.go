package cmd

import (
	"errors"

	"github.com/desertbit/grumble"
	"github.com/manifoldco/promptui"
)



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
	CurrentAgentId = AgentList.agents[idx].Metadata.ID
    return nil
}
