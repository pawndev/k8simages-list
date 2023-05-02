package kubernetes

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (k *Kube) GetJobImages() []string {
	var imageList []string
	podList, err := k.clientset.BatchV1().Jobs(k.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range podList.Items {
		for _, container := range pod.Spec.Template.Spec.Containers {
			imageList = append(imageList, strings.Split(container.Image, ":")[0])
		}
	}

	return imageList
}
