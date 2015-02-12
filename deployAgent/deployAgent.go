package main

import (
	"encoding/base64"
	"fmt"
	"github.com/CenturylinkLabs/kube-cluster-deploy/utils"
	"os"
    "github.com/CenturylinkLabs/agent-server-deploy/provider"
    )

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	utils.LogInfo("\nDeploying Agent Server...")
    e := utils.LoadStdinToEnvAndKeys()
    if  e != nil {
        panic(e)
    }

    cp := provider.New(os.Getenv("CLOUD_PROVIDER"))
    s, e := cp.ProvisionAgent()
    if e != nil {
        panic(e)
    }

    utils.SetKey("AGENT_PRIVATE_KEY", base64.StdEncoding.EncodeToString([]byte(s.PrivateSSHKey)))
    utils.SetKey("AGENT_PUBLIC_IP", s.PublicIP)

    utils.LogInfo("\nAgent server deployment complete!!")
}
