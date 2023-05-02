package kubernetes

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (k *Kube) GetDaemonsetImages() []string {
	var imageList []string
	daemonSetList, err := k.clientset.AppsV1().DaemonSets(k.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, daemonSet := range daemonSetList.Items {
		for _, container := range daemonSet.Spec.Template.Spec.Containers {
			imageList = append(imageList, strings.Split(container.Image, ":")[0])
		}
	}

	return imageList
}
