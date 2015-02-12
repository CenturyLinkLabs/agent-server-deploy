package provider

import (
    "github.com/CenturylinkLabs/kube-cluster-deploy/deploy"
    "strings"
    "fmt")

type CloudProvider interface {
    ProvisionAgent() (deploy.CloudServer, error)
}

func New(providerType string) CloudProvider {
    pt := strings.TrimSpace(strings.ToLower(providerType))
    switch string(pt) {
        case "centurylink":
            return NewCenturylink()
        case "amazon":
            fmt.Printf("\n\nAMAZON\n\n")
            return NewAmazon()
    }
    fmt.Printf("NIL Provider:%s", pt)
    return nil
}
