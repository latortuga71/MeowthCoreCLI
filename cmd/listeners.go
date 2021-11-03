package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/desertbit/grumble"
	"github.com/latortuga71/MeowthCoreCLI/internal"
)


type Listeners struct {
	Name     *string `json:"name"`
	BindPort *int    `json:"bindPort,omitempty"`
}


func init(){
    listenerCommand := &grumble.Command{
        Name:"listener",
        Help:"Interact with listeners",
    }
    App.AddCommand(listenerCommand)
    listenerCommand.AddCommand(&grumble.Command{
        Name:"add",
        Help:"Add listeners",
        Run: addListeners,
        Args: func (a *grumble.Args){
            a.String("name","listener name")
            a.Int("port","listener port")
        },
    })

    listenerCommand.AddCommand(&grumble.Command{
        Name:"list",
        Help:"List listeners",
        Run: listListeners,
    })

    listenerCommand.AddCommand(&grumble.Command{
        Name:"get",
        Help:"Get listener",
        Run: getListeners,
        Args: func (a *grumble.Args){
            a.String("name","listener name")
        },
    })

    listenerCommand.AddCommand(&grumble.Command{
        Name:"delete",
        Help:"Delete listener",
        Run: deleteListeners,
        Args: func (a *grumble.Args){
            a.String("name","listener name")
        },
    })
}

func deleteListeners(c *grumble.Context) error {
	name := c.Args.String("name")
	var listener Listeners
	target := fmt.Sprintf("%s%s%s",FullServerURI,"Listeners/",name)
	err := internal.Delete(target, &listener)
	if err != nil {
		return err
	}
	_, err = json.Marshal(listener)
	if err != nil {
		return err
	}
	fmt.Println("OK")
    return nil
}

func getListeners(c *grumble.Context) error {
	name := c.Args.String("name")
	var listener Listeners
	target := fmt.Sprintf("%s%s%s",FullServerURI,"Listeners/",name)
	err := internal.Get(target, &listener)
	if err != nil {
		return err
	}
	data, err := json.Marshal(listener)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
    return nil
}


func listListeners(c *grumble.Context) error {
	var listener []Listeners
	target := fmt.Sprintf("%s%s",FullServerURI,"Listeners")
	err := internal.Get(target, &listener)
	if err != nil {
		return err
	}
	for _,l := range listener{
		data, err := json.Marshal(l)
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	}
    return nil
}

func addListeners(c *grumble.Context) error {
	name := c.Args.String("name")
    port := c.Args.Int("port")
    listener := &Listeners{
        Name: &name,
        BindPort: &port,
    }
	target := fmt.Sprintf("%s%s",FullServerURI,"Listeners")
    internal.Post(target,listener)
    return nil
}
