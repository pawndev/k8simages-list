package kubernetes

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

func (k *Kube) GetPodImages() []string {
	var imageList []string
	podList, err := k.clientset.CoreV1().Pods(k.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range podList.Items {
		for _, container := range pod.Spec.Containers {
			imageList = append(imageList, strings.Split(container.Image, ":")[0])
		}
	}

	return imageList
}
