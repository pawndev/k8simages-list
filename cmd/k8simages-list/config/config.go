package config

import (
	apiv1 "k8s.io/api/core/v1"
	"os"
)

const (
	KubernetesContext   = "K8S_CONTEXT"
	KubernetesNamespace = "K8S_NAMESPACE"
)

type Config struct {
	Context   string
	Namespace string
}

func New() *Config {
	return &Config{
		Context:   getKubernetesContextName(""),
		Namespace: getNamespace(apiv1.NamespaceAll),
	}
}

func getNamespace(defaultValue string) string {
	envNamespace := os.Getenv(KubernetesNamespace)
	if envNamespace != "" {
		return envNamespace
	}
	return defaultValue
}

func getKubernetesContextName(defaultValue string) string {
	envContext := os.Getenv(KubernetesContext)
	if envContext != "" {
		return envContext
	}
	return defaultValue
}
