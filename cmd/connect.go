package cmd

import (
	"errors"
	"fmt"
    "log"
	"net/http"
	"time"

	"github.com/desertbit/grumble"
	"github.com/latortuga71/MeowthCoreCLI/internal"
)




func init(){
    connectCommand := &grumble.Command{
        Name:"connect",
        Help:"Connect to C2",
        Run: connect,
        Args: func (a *grumble.Args){
            a.String("ip","connect name")
            a.Int("port","connect port")
        },
    }
    App.AddCommand(connectCommand)
}

func pollForNewAgents(initialRun bool) {
    for {
        time.Sleep(time.Second * 5)
        localIds := make(map[string]bool)
	    var agents []Agent
	    target := fmt.Sprintf("%s%s",FullServerURI,"Agents")
	    err := internal.Get(target, &agents)
	    if err != nil {
	    	log.Fatal(err)
	    }
        // loop through local agent ids and save to map
        for _,local := range AgentList.agents {
            localIds[local.Metadata.ID] = true
        }
        // loop through remtoe ids and see if new ones exist
        for _, remote := range agents {
            if ok := localIds[remote.Metadata.ID]; !ok {
                AgentList.Lock()
                AgentList.agents = append(AgentList.agents,remote)
                AgentList.Unlock()
                if !initialRun {
                    fmt.Printf("\n::: New Agent %s Connected :::\n",remote.Metadata.ID)
                }
            }
        }
        if initialRun {
            break
        }
    }
}
func connect(c *grumble.Context) error {
    host := c.Args.String("ip")
    port := c.Args.Int("port")
    FullServerURI = fmt.Sprintf("http://%s:%d/",host,port)
    fmt.Printf("Connecting.... %s:%d\n",host,port)
    resp,err := http.Get(FullServerURI + "Agents")
    if resp.StatusCode == 200 && err == nil {
        ServerSet = true
        fmt.Println("Successfully connected to server")
        fmt.Println("Starting agents discovery goroutine...")
        // run once to get whoevers already connected
        pollForNewAgents(true) 
        // run again for new incoming agents
        go pollForNewAgents(false)
        return nil
    }
    return errors.New("Failed to connect to server") 
}