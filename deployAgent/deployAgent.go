package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/CenturylinkLabs/kube-cluster-deploy/utils"
	"os"
	"strings"
    "github.com/CenturylinkLabs/agent-server-deploy/provider")

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	utils.LogInfo("\nDeploying Agent Server")

	var ln string
	var e error
	rd := bufio.NewReader(os.Stdin)

	for {
		ln, e = rd.ReadString('\n')
		println(ln)
		if e != nil {
			panic(e)
		} else if strings.ContainsAny(ln, "AGENT_KUBER_API") {
			break
		}
	}

    kv := strings.Split(ln, "=")
	if !strings.ContainsAny(ln, "AGENT_KUBER_API") || len(kv) == 0{
		panic("Missing Key, AGENT_KUBER_API. Cannot proceed.")
	}
	utils.SetKey("AGENT_KUBER_API", kv[1])

    utils.LogInfo("\nStarting Agent installation")
    cp := provider.New("amazon")
    s, e := cp.ProvisionAgent()

    if e != nil {
        panic(e)
    }

    utils.SetKey("AGENT_PRIVATE_KEY", base64.StdEncoding.EncodeToString([]byte(s.PrivateSSHKey)))
    utils.SetKey("AGENT_PUBLIC_IP", s.PublicIP)

    utils.LogInfo("\nAgent server deployment complete!!")
}
