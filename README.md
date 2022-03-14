# vSphere Charts

The repository provides charts for vSphere CSI and CPI based on the upstream repositories and the manifests provided by those
repositories.

## Prerequisites
- Helm 3.x

## vSphere CSI Chart

This chart is produced using the following repository: https://github.com/kubernetes-sigs/vsphere-csi-driver

The manifests are located [here](https://github.com/kubernetes-sigs/vsphere-csi-driver/tree/master/manifests). The workflow is to compare 
the existing helm templates to manifests in that repository when a new version has been released. Make any changes that are required to make the 
templates have parity with the manifests. Then submit your PR.

Any images consumed by this chart need to be mirrored in the following [repository](https://github.com/rancher/image-mirror) first.

## vSphere CPI Charts

This chart is produced using the following repository: https://github.com/kubernetes/cloud-provider-vsphere/

The manifests are located [here](https://github.com/kubernetes/cloud-provider-vsphere/tree/master/releases). The workflow is to compare
the existing helm templates to manifests in that repository when a new version has been released. Make any changes that are required to make the
templates have parity with the manifests. Then submit your PR.

Any images consumed by this chart need to be mirrored in the following [repository](https://github.com/rancher/image-mirror) first.

## Using charts in rancher/charts and rancher/rke2-charts

Charts from this repository should be consumed by commit hash based on the version and features that you want to have
included.
