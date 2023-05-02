package kubernetes

import (
	"flag"
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"sync"
)

type ImageFetcher func() []string

type KubeInformation interface {
	Namespace() string
	Context() string
}

type RestClient interface {
	KubeInformation
	RestConfig() *rest.Config
}

type Kube struct {
	context    string
	namespace  string
	kubeconfig *string
	rest       *rest.Config
	clientset  *k8s.Clientset
}

func New(context, namespace string) *Kube {
	kubeconfig := loadConfigFromPath()
	restConfig, err := buildConfigFromFlags(context, *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := k8s.NewForConfig(restConfig)
	if err != nil {
		panic(err.Error())
	}
	return &Kube{
		kubeconfig: kubeconfig,
		rest:       restConfig,
		namespace:  namespace,
		context:    context,
		clientset:  clientset,
	}
}

func (k *Kube) Clientset() *k8s.Clientset {
	return k.clientset
}

func (k *Kube) RestConfig() *rest.Config {
	return k.rest
}

func (k *Kube) Namespace() string {
	return k.namespace
}

func (k *Kube) Context() string {
	return k.context
}

func (k *Kube) GetAllImages() []string {
	var requestList = []ImageFetcher{
		k.GetJobImages,
		k.GetDaemonsetImages,
		k.GetCronjobImages,
		k.GetPodImages,
		k.GetDeploymentImages,
		k.GetStatefulSet,
	}
	var wg sync.WaitGroup
	wg.Add(len(requestList))
	p := sync.Pool{New: func() any {
		return []string{}
	}}

	for _, request := range requestList {
		go func(req ImageFetcher) {
			defer wg.Done()
			images := req()
			all := p.Get().([]string)
			all = append(all, images...)
			p.Put(all)
		}(request)
	}

	wg.Wait()

	return p.Get().([]string)
}

func loadConfigFromPath() *string {
	if home := homedir.HomeDir(); home != "" {
		return flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		return flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
}

func buildConfigFromFlags(context, config string) (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: config},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}
