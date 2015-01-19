package main

import (
	"bufio"
	"fmt"
	"github.com/CenturyLinkLabs/k8s-provision-vms/deploy"
	"github.com/CenturyLinkLabs/k8s-provision-vms/utils"
	"os"
	"strings"
)

func main() {

	fmt.Printf("\nDeploying Agent Server")

	puk, pk, _ := utils.CreateSSHKey()

	fmt.Printf("\nSSH Key Created")

	c := deploy.CenturyLink{
		PrivateSSHKey: pk,
		PublicSSHKey:  puk,
		APIUsername:   os.Getenv("USERNAME"),
		APIPassword:   os.Getenv("PASSWORD"),
		GroupID:       os.Getenv("GROUP_ID"),
		CPU:           1,
		MemoryGB:      1,
		TCPOpenPorts:  []int{3036, 8080},
		ServerName:    "AGENT",
	}

	s, e := c.DeployVM()

	if e != nil {
		fmt.Printf("\n%s", e.Error())
		panic(e)
	}

	utils.SetKey("AGENT_PRIVATE_KEY", pk)
	utils.SetKey("AGENT_PUB_IP", s.PublicIP)

	r := bufio.NewReader(os.Stdin)
	text, _ := r.ReadString('\n')
	fmt.Printf("STDIN:%s", text)
	kv := strings.Split(text, "=")
	utils.SetKey("AGENT_KUBE_API", kv[1])
}
