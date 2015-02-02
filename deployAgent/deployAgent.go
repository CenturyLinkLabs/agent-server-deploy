package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/CenturyLinkLabs/kube-cluster-deploy/deploy"
	"github.com/CenturyLinkLabs/kube-cluster-deploy/utils"
	"os"
	"strings"
)

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

	if !strings.ContainsAny(ln, "AGENT_KUBER_API") {
		panic("Missing Key, AGENT_KUBER_API. Cannot proceed.")
	}

	kv := strings.Split(ln, "=")
	utils.SetKey("AGENT_KUBER_API", kv[1])

	pk, puk, _ := utils.CreateSSHKey()

	c := deploy.CenturyLink{
		PublicSSHKey:   puk,
		APIUsername:    os.Getenv("USERNAME"),
		APIPassword:    os.Getenv("PASSWORD"),
		GroupID:        os.Getenv("GROUP_ID"),
		CPU:            1,
		MemoryGB:       1,
		TCPOpenPorts:   []int{3001},
		ServerName:     "AGENT",
		ServerTemplate: "UBUNTU-14-64-TEMPLATE",
	}

	utils.LogInfo("\nWaiting for server creation")
	s, e := c.DeployVM()

	if e != nil {
		panic(e)
	}

	utils.SetKey("AGENT_PRIVATE_KEY", base64.StdEncoding.EncodeToString([]byte(pk)))
	utils.SetKey("AGENT_PUBLIC_IP", s.PublicIP)

	utils.LogInfo("\nAgent server deployment complete!!")
}
