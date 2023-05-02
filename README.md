# K8Simages

## Summary

This is a little script that will report all images used in your
kubernetes cluster.

Actually, it supports these `Kind`:
- [`Deployment`](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [`Job`](https://kubernetes.io/docs/concepts/workloads/controllers/job/)
- [`Pod`](https://kubernetes.io/fr/docs/concepts/workloads/pods/)
- [`Cronjob`](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/)
- [`Statefulset`](https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/)
- [`Daemonset`](https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/)
- [`Workflow`](https://argoproj.github.io/argo-workflows/workflow-concepts/)
- [`WorkflowTemplate`](https://argoproj.github.io/argo-workflows/workflow-templates/)
- [`CronWorkflow`](https://argoproj.github.io/argo-workflows/cron-workflows/)
- [`ClusterWorkflowTemplate`](https://argoproj.github.io/argo-workflows/cluster-workflow-templates/)

## Install

You can get the binary on the release section

Or you can get it with:

```shell
go install github.com/pawndev/k8simages-list/cmd/k8simages-list
```

## Usage

The binary can be configured with these environment variables:

| name          | summary                             | default              |
|---------------|-------------------------------------|----------------------|
| K8S_NAMESPACE | the namespace where to fetch images | '' => all namespaces |
| K8S_CONTEXT   | the context to use                  | your active context  |

And then you can start it with:

```shell
K8S_NAMESPACE=default K8S_CONTEXT=preprod k8simages-list
```


