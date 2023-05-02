package argoproj

import (
	"context"
	argoworkflow "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
	"github.com/pawndev/k8simages/pkg/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
	"sync"
)

type Argo struct {
	client *argoworkflow.Clientset
	kube   kubernetes.RestClient
}

type ImageFetcher func() []string

func New(kube kubernetes.RestClient) *Argo {
	client, err := argoworkflow.NewForConfig(kube.RestConfig())
	if err != nil {
		panic(err)
	}
	return &Argo{
		client: client,
		kube:   kube,
	}
}

func (a *Argo) GetAllImages() []string {
	var requestList = []ImageFetcher{
		a.GetWorkflowImages,
		a.GetWorkflowTemplateImages,
		a.GetClusterWorkflowTemplateImages,
		a.GetCronWorkflowImages,
	}
	var allImages []string
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, request := range requestList {
		wg.Add(1)
		go func(req ImageFetcher) {
			defer wg.Done()
			images := req()
			mu.Lock()
			allImages = append(allImages, images...)
			defer mu.Unlock()
		}(request)
	}

	wg.Wait()

	return allImages
}

func (a *Argo) GetWorkflowImages() []string {
	var imageList []string

	workflowList, err := a.client.ArgoprojV1alpha1().Workflows(a.kube.Namespace()).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, workflow := range workflowList.Items {
		for _, tpl := range workflow.Spec.Templates {
			if tpl.Container == nil || tpl.Container.Image == "" {
				continue
			}
			imageList = append(imageList, strings.Split(tpl.Container.Image, ":")[0])
		}
	}

	return imageList
}

func (a *Argo) GetCronWorkflowImages() []string {
	var imageList []string

	workflowList, err := a.client.ArgoprojV1alpha1().CronWorkflows(a.kube.Namespace()).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, workflow := range workflowList.Items {
		for _, tpl := range workflow.Spec.WorkflowSpec.Templates {
			if tpl.Container == nil || tpl.Container.Image == "" {
				continue
			}
			imageList = append(imageList, strings.Split(tpl.Container.Image, ":")[0])
		}
	}

	return imageList
}

func (a *Argo) GetWorkflowTemplateImages() []string {
	var imageList []string

	workflowTemplates, err := a.client.ArgoprojV1alpha1().WorkflowTemplates(a.kube.Namespace()).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, workflowTemplate := range workflowTemplates.Items {
		for _, tpl := range workflowTemplate.Spec.Templates {
			if tpl.Container == nil || tpl.Container.Image == "" {
				continue
			}
			imageList = append(imageList, strings.Split(tpl.Container.Image, ":")[0])
		}
	}

	return imageList
}

func (a *Argo) GetClusterWorkflowTemplateImages() []string {
	var imageList []string

	clusterWorkflowTemplateList, err := a.client.ArgoprojV1alpha1().ClusterWorkflowTemplates().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, clusterWorkflowTemplate := range clusterWorkflowTemplateList.Items {
		for _, tpl := range clusterWorkflowTemplate.Spec.Templates {
			imageList = append(imageList, strings.Split(tpl.Container.Image, ":")[0])
		}
	}

	return imageList
}
