# Developer Guide for vSphere Charts for Rancher

This guide is going to highlight the steps that are required to develop and maintian the vSphere charts for [CPI](https://github.com/kubernetes/cloud-provider-vsphere) and [CSI](https://github.com/kubernetes-sigs/vsphere-csi-driver) for Rancher. Upstream is only releasing manifests at this point and they do that in version specific branches. We use these upstream manifests to create the charts that we use. This makes the process a touch more complex as each manifest needs to be compared to the helm chart template to make the appropriate changes. Since we mirror images into our own container registry, we also have to check the manifests every time a release happens to gather the correct image versions to keep our registry synced. Once the charts and images have been updated, we then need to ensure that the charts are updated in [Rancher Charts](https://github.com/rancher/charts), [RKE2-Charts](https://github.com/rancher/rke2-charts), and [RKE2](https://github.com/rancher/rke2).

#### You may be asking why both Rancher charts and RKE2?

RKE2 embeds the charts that it ships with each version of RKE2 that is released. This means that RKE2 always uses its own version of the vSphere charts regardless if deployed standalone or through Rancher. The charts being deployed are determined by the selected cloud provider when provisioning. Rancher charts shouldn't be consumed for RKE2 and currently exist for use with non-RKE2 clusters that are provisioned in Rancher, primarily RKE1.

#### Upstream Releases

Upstream releases are not immediate when a new version of Kubernetes is released. This means that the charts typically lag behind the Kubernetes version by at least a month if not longer. This often means that users can't upgrade clusters immediately or deploy the newest version that is released in RKE2 or Rancher with full support.

#### Cloud Provider Skew Policy

Cloud providers, cloud controller manager, in Kubernetes support a [skew policy](https://kubernetes.io/releases/version-skew-policy/#kube-controller-manager-kube-scheduler-and-cloud-controller-manager). This allows us to use one minor version older of a CPI on Kubernetes cluster. The primary purpose is to allow upgrades, yet also gives us the ability to deploy an older CPI before the official one has been released for the new version of Kubernetes.

## Chart Design

Both the CSI and CPI charts in this repository are designed to be what have been called **super charts**. A super chart is a chart that uses a helm template function to determine which image versions should be deployed based on the Kubernetes version of the cluster. The image versions and configuration is part of an array of values that use semver to determine if it can be deployed on a Kubernetes version.

By creating a super chart, it means there is only ever one chart and we can ensure that users are always getting the most up to date images based on their Kubernetes version. This technique is used in other charts inside of Rancher.

## Development

Features and upstream releases are the primary reason to be working in this repository. Upstream releases through either bug fixes or new Kubernetes version support is the majority of the work that is required. Below is a high level of the steps that need to occur.

* Check the upstream manifests for changes and new image versions
* Update the [image-mirror](https://github.com/rancher/image-mirror) repository for new images
* Update the charts with the new images and manifest changes
* Update the chart versions (if necessary)
* Update the chart tests
* Locally compile and deploy to a Kubernetes cluster to ensure that the chart functions
* Check with RKE2 on the next release cycle and plan out PRs for getting these changes included by the target release date.
* Update Rancher Charts with the new versions

Before getting started, it is suggested to create an issue in the Rancher/Rancher repository to track the work that has to occur. It spans several different repositories which is often hard to coordinate. Here is the recommend issue format.

```
**Is your feature request related to a problem? Please describe.**



**Describe the solution you'd like**



# Checklist for items that need completed

- [ ] Update Image-mirror repo
- [ ] Update manifests in rancher/vsphere-charts
- [ ] Verify Slack hook notified about chart "release" in relevant channels
- [ ] Update chart in rancher/charts
- [ ] Update chart in rancher/rke2-charts (RKE2 team)
- [ ] Updated RKE2 to consume new chart (RKE2 team)
```

For every PR, link back to this issue so everything can be tracked. Whoever does the RKE2 updates will need to create a separate issue in the RKE2 repository in addition to the Rancher one.

Here are some examples:

* [Rancher Issue](https://github.com/rancher/rancher/issues/38188)
* [RKE2 issue](https://github.com/rancher/rke2/issues/2867)

### Step 1: Checking upstream manifests

NOTE: the charts _do not match_ the upstream charts exactly - the `rancher/vsphere-charts` package actually predates those by a lot due to the fact that upstream did NOT ship a chart instead they just shipped raw manifests which isn't very easy to package/update. Engineers used those initial manifests and made these from them.

There are two goals when checking upstream manifests. The first goal is to check to see if any image versions have been updated and the second is to check if any changes in the manifests have occured.

If you find that image versions have been updated, then you will need to make a pull request to the [image-mirror](https://github.com/rancher/image-mirror) repository first before proceeding. [Here](https://github.com/rancher/image-mirror/pull/261) is an example PR and follow the requirements outlined in the PR template for those changes. After the images have been checked and image-mirror updated with any new images or image versions, then proceed to check check the manifests.

Checking the manifests is best done by doing a tag compare in each upstream chart. Compare an upstream tag to the tag that the rancher chart is based off of by viewing the `Full Changelog` for that tag. This compares the manifest directories. If there is a change that needs to be added to charts in this repository, bring it in. [Here](https://github.com/kubernetes/cloud-provider-vsphere/compare/v1.25.0...v1.24.2) is an example. Just look for important changes and bring them into this chart as necessary. New environment variables, configs, etc.

Make sure that the unit tests for each chart have been updated, or new tests or test cases are added based on the set of manifest changes.

Finally increment the chart version, if necessary. We need to update the chart version anytime that a new tag of CSI/CPI is released and we update to that version of upstream.

CPI and CSI versioning is different because CSI tracks the version from upstream so it needs a `-rancher1` for when we need to do a Rancher specific patch, whereas CPI doesn’t track the same version from upstream (they have multiple CPIs based on Kubernetes version. Our CPI chart is based on all of them together). We use our own versioning so we can bump the patch for Rancher specific versions. [Here](https://github.com/rancher/vsphere-charts/pull/36/commits/5bb4f6a042dc31ace66522fe670a72930d5d3e63) is an example.


### Step 2: Compile and check locally

At this point you can verify everything passes linting and tests by running:

```
$ make ci

### linting charts ###
helm lint ./charts/rancher-vsphere-cpi/
==> Linting ./charts/rancher-vsphere-cpi/

1 chart(s) linted, 0 chart(s) failed
helm lint ./charts/rancher-vsphere-csi/
engine.go:167: [INFO] Missing required value: .Values.vCenter.clusterId must be provided
==> Linting ./charts/rancher-vsphere-csi/

1 chart(s) linted, 0 chart(s) failed
### packaging charts for testing purposes ###
helm package ./charts/rancher-vsphere-cpi/
Successfully packaged chart and saved it to: /home/jphillips/work/vsphere-charts/rancher-vsphere-cpi-1.4.0.tgz
helm package ./charts/rancher-vsphere-csi/
Successfully packaged chart and saved it to: /home/jphillips/work/vsphere-charts/rancher-vsphere-csi-2.5.1-rancher1.tgz
### removing chart packages ###
rm *.tgz
### setup ###
go mod tidy
### running unit tests ###
go test -v -tags helm ./tests/unit
--- PASS: TestCSITemplateRenderedControllerDeployment (0.53s)
    --- PASS: TestCSITemplateRenderedControllerDeployment/Kubernetes_1.23 (0.09s)
    --- PASS: TestCSITemplateRenderedControllerDeployment/Kubernetes_1.22 (0.09s)
    --- PASS: TestCSITemplateRenderedControllerDeployment/Kubernetes_1.21 (0.09s)
    --- PASS: TestCSITemplateRenderedControllerDeployment/Kubernetes_1.20 (0.09s)
    --- PASS: TestCSITemplateRenderedControllerDeployment/Kubernetes_1.23_with_CSI_Resizer_Enabled (0.09s)
    --- PASS: TestCSITemplateRenderedControllerDeployment/Kubernetes_1.20_with_CSI_Resizer_Enabled (0.08s)
PASS
ok  	github.com/rancher/vsphere-charts/tests/unit	(cached)
```

If everything passes here then it's time to test deployment to a cluster. If something errors, then please correct it and ensure all checks pass.

### Step 3: Deploying the chart

At this stage, it's now safe to test the charts on a real cluster. Typically, a RKE1 cluster is deployed on vSphere using Rancher. Then the kubeconfig is downloaded from the cluster and used to deploy the local chart to the cluster using helm. Here is an example.

```
$ helm install -f values.yaml --kubeconfig <path to kube config> csi ./charts/rancher-vsphere-csi
```

```
$ helm install -f values.yaml --namespace kube-system --kubeconfig <path to kube config> cpi ./charts/rancher-vsphere-cpi
```

The CPI chart **must** be installed in the `kube-system` namespace.

Make sure to supply the vSphere configuration required. Your `values.yaml` is the values that get passed to helm on a chart deploy. Here's an example.

```
vCenter:
  host: <host>
  port: <port>
  insecureFlag: true
  clusterId: "test-cluster-id"
  datacenters: <datacenter>
  username: <username>
  password: <password>
```

If you want to test that CPI is functioning properly, add the `node.cloudprovider.kubernetes.io/uninitialized=true:NoSchedule` taint to the cluster nodes before installing the chart. When you install CPI, the kubelet should remove the taints.

If you deploy CPI via the Rancher UI, **do not** `Define vSphere Tags` without specifying the `Region` and `Zone` in which to tag pods. This will cause the chart to malfunction and the kubelet will not remove taints or set the provider ID on the nodes.

Ensure that all pods for both the CPI and CSI charts come up without errors. Make sure that the provider gets set on each node to be vSphere.

### Step 4: PR to Rancher vSphere Charts Repository

You are now ready to submit a PR to the vSphere charts repository. Fill out the PR template and request reviews. Once the PR has been approved and passed the Drone run, you can merge. [Here](https://github.com/rancher/vsphere-charts/pull/35) is an example PR.

### Step 5: Create RKE2-Charts PR

A member from the RKE2 team can now pull the chart into the RKE2-Charts repository. [Here](https://github.com/rancher/rke2-charts/pull/250) is an example PR.

### Step 6: Create RKE2 PR

Once the RKE2-Charts PR has been merged, a member from the RKE2 team can update the chart in the RKE2 repository. This step requires a backport PR to each branch of the still supported versions of RKE2. [Here](https://github.com/rancher/rke2/pull/2868) is an example PR and an [example](https://github.com/rancher/rke2/pull/2870) backport.

### Step 7: Create Rancher Charts PR

The last step is to pull the vSphere charts into the Rancher charts repository. The process is well documented, so follow the instructions listed in that repository for the workflow. [Here](https://github.com/rancher/charts/pull/1898) is an example PR.

If there is an existing patch for CSI/CPI in charts upstream, you may have to remove the generated-changes folder, reapply your changes, manually apply the upstream patch, then regenerate the patch. [Here](https://github.com/rancher/charts/pull/2204) is an example PR. Team 3 is working on a 3-way rebase script that may be available soon.

Here are walkthrough steps,

* Delete the generated-changes dir
* Run `make prepare`
* Apply your changes: bump chart versions and `release.yaml` to track the new versions. Manually reapply any patches from charts upstream
* Run `make patch && make clean`
* Run `git commit -am <message>`
* Run `make charts`
* Run `git commit -m "Make charts"`
* Check the generated CSI/CPI chart files against vsphere-charts upstream
* Run `make validate` (I always do this to confirm the PR build will pass before pushing!)
