# vSphere Charts

## Prerequisites
- Helm 3.x

## Use released charts in rancher/charts
Use below tarballs in rancher/charts to pull released charts.

```
CPI: https://github.com/rancher/vsphere-charts/releases/download/rancher-vsphere-cpi-x.y.z/rancher-vsphere-cpi-x.y.z.tgz

CSI: https://github.com/rancher/vsphere-charts/releases/download/rancher-vsphere-csi-x.y.z-rancherx/rancher-vsphere-csi-x.y.z-rancherx.tgz
```
## Releasing New Charts
When ready to release a new version or add a new chart, make changes to CPI and CSI which are in `charts/` directory. Make sure the chart directories are named after the actual charts (for example: `rancher-vsphere-cpi`)

Once a PR is merged or pushed, GitHub Actions will look for changes to charts in the `charts/` directory since the last release. It will package the updated chart and then release it with a new tag.

Note that changes should only be synced to this repository when you are ready to create a new release. GitHub Actions will fail if changes are made to the charts, without updating the chart version. Chart Releaser will not attempt to override a previously released version. In case of a new change, update the charts along with the chart version & then raise a PR. This will create a new release and tag it with the given version.

