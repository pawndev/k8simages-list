package main

import (
	"fmt"
	"github.com/pawndev/k8simages/cmd/k8simages-list/config"
	"github.com/pawndev/k8simages/pkg/kubernetes"
	"github.com/pawndev/k8simages/pkg/sliceutils"
)

func main() {
	cfg := config.New()
	kube := kubernetes.New(cfg.Context, cfg.Namespace)
	//argo := argoproj.New(kube)

	var allImages []string
	allImages = append(allImages, kube.GetAllImages()...)
	//allImages = append(allImages, argo.GetAllImages()...)
	dedupeImages := sliceutils.RemoveDuplicate(allImages)

	for _, image := range dedupeImages {
		fmt.Printf("%s\n", image)
	}
}
