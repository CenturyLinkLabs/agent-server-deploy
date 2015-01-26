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

	fmt.Printf("\nDeploying Agent Server")

	pk, puk, _ := utils.CreateSSHKey()

	fmt.Printf("\nSSH Key Created")

	c := deploy.CenturyLink{
		PublicSSHKey:   puk,
		APIUsername:    os.Getenv("USERNAME"),
		APIPassword:    os.Getenv("PASSWORD"),
		GroupID:        os.Getenv("GROUP_ID"),
		CPU:            1,
		MemoryGB:       1,
		TCPOpenPorts:   []int{3036, 8080},
		ServerName:     "AGENT",
		ServerTemplate: "UBUNTU-14-64-TEMPLATE",
	}

	s, e := c.DeployVM()

	if e != nil {
		fmt.Printf("\n%s", e.Error())
		panic(e)
	}

	utils.SetKey("AGENT_PRIVATE_KEY", base64.StdEncoding.EncodeToString([]byte(pk)))
	//utils.SetKey("AGENT_PUBLIC_KEY", base64.StdEncoding.EncodeToString([]byte(puk)))
	utils.SetKey("AGENT_PUBLIC_IP", s.PublicIP)

	line := ""
	reader := bufio.NewReader(os.Stdin)

	for {
		line, e = reader.ReadString('\n')
		if e != nil || strings.ContainsAny(line, "AGENT_KUBER_API") {
			break
		}
	}
	fmt.Printf("STDIN:%s", line)
	kv := strings.Split(line, "=")
	utils.SetKey("AGENT_KUBER_API", kv[1])
}
