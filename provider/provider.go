package provider

import (
    "github.com/CenturylinkLabs/kube-cluster-deploy/deploy"
    "strings"
)

type CloudProvider interface {
    ProvisionAgent() (deploy.CloudServer, error)
}

func New(providerType string) CloudProvider {
    providerType = strings.ToLower(providerType)
    switch providerType {
        case "centurylink":
            return NewCenturylink()
        case "amazon":
            return NewAmazon()
    }
    return nil
}
