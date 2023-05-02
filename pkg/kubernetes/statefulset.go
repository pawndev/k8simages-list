package kubernetes

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (k *Kube) GetStatefulSet() []string {
	var imageList []string
	statefulsetList, err := k.clientset.AppsV1().StatefulSets(k.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, deploy := range statefulsetList.Items {
		for _, container := range deploy.Spec.Template.Spec.Containers {
			imageList = append(imageList, strings.Split(container.Image, ":")[0])
		}
	}

	return imageList
}
