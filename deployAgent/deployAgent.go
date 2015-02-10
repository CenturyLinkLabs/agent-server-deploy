package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"github.com/CenturylinkLabs/kube-cluster-deploy/utils"
	"os"
	"strings"
    "github.com/CenturylinkLabs/agent-server-deploy/provider"
    "io")

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	utils.LogInfo("\nDeploying Agent Server...")

	var e error
	rd := bufio.NewReader(os.Stdin)

	for {
		ln, e := rd.ReadString('\n')
		if e == io.EOF {
            break
		} else if e != nil {
            panic(e)
        } else if strings.Contains(ln, "AGENT_KUBER_API") {
            kv := strings.SplitN(ln, "=", 2)
            utils.SetKey("AGENT_KUBER_API", kv[1])
        } else if strings.Contains(ln, "MASTER_PRIVATE_KEY") {
            kv := strings.SplitN(ln, "=", 2)
            utils.SetKey("MASTER_PRIVATE_KEY", kv[1])
            os.Setenv("MASTER_PRIVATE_KEY", kv[1])
        } else if strings.Contains(ln, "MASTER_PUBLIC_KEY") {
            kv := strings.SplitN(ln, "=", 2)
            utils.SetKey("MASTER_PUBLIC_KEY", kv[1])
            os.Setenv("MASTER_PUBLIC_KEY", kv[1])
        } else if strings.Contains(ln, "AMAZON_MASTER_KEY_NAME") {
            kv := strings.SplitN(ln, "=", 2)
            utils.SetKey("AMAZON_MASTER_KEY_NAME", kv[1])
            os.Setenv("AMAZON_MASTER_KEY_NAME", kv[1])
        }
    }

    cp := provider.New("amazon")
    s, e := cp.ProvisionAgent()

    if e != nil {
        panic(e)
    }

    utils.SetKey("AGENT_PRIVATE_KEY", base64.StdEncoding.EncodeToString([]byte(s.PrivateSSHKey)))
    utils.SetKey("AGENT_PUBLIC_IP", s.PublicIP)
    utils.SetKey("UBUNTU_LOGIN_USER", "ubuntu")

    utils.LogInfo("\nAgent server deployment complete!!")
}
